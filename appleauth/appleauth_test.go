package appleauth

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"math/big"
	"os"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestTokenClaimsAndSignature(t *testing.T) {
	privateKey, keyPEM := testKey(t)
	now := time.Date(2026, 11, 26, 12, 0, 0, 0, time.UTC)
	signer, err := New(Config{
		KeyID: "KEY123", Issuer: "issuer-123", TeamID: "team-123",
		Audience: "appstoreconnect-v1", TTL: 5 * time.Minute,
		RefreshBefore: time.Minute, Now: func() time.Time { return now },
	}, keyPEM)
	if err != nil {
		t.Fatal(err)
	}
	token, err := signer.Token()
	if err != nil {
		t.Fatal(err)
	}
	parts := splitToken(t, token)
	var header map[string]any
	var claims map[string]any
	decodeJSONPart(t, parts[0], &header)
	decodeJSONPart(t, parts[1], &claims)
	if header["alg"] != "ES256" || header["kid"] != "KEY123" || header["typ"] != "JWT" {
		t.Fatalf("unexpected JWT header: %#v", header)
	}
	if claims["iss"] != "issuer-123" || claims["aud"] != "appstoreconnect-v1" {
		t.Fatalf("unexpected JWT claims: %#v", claims)
	}
	if claims["iat"] != float64(now.Unix()) || claims["exp"] != float64(now.Add(5*time.Minute).Unix()) {
		t.Fatalf("unexpected token times: %#v", claims)
	}
	signature, err := base64.RawURLEncoding.DecodeString(parts[2])
	if err != nil || len(signature) != 64 {
		t.Fatalf("invalid ES256 signature encoding: %v", err)
	}
	hash := sha256.Sum256([]byte(parts[0] + "." + parts[1]))
	if !ecdsa.Verify(&privateKey.PublicKey, hash[:], new(big.Int).SetBytes(signature[:32]), new(big.Int).SetBytes(signature[32:])) {
		t.Fatal("JWT signature did not verify")
	}
}

func TestNewFromFileAndTeamIDIssuerFallback(t *testing.T) {
	_, keyPEM := testKey(t)
	path := t.TempDir() + "/AuthKey.p8"
	if err := os.WriteFile(path, keyPEM, 0o600); err != nil {
		t.Fatal(err)
	}
	signer, err := NewFromFile(Config{KeyID: "KEY123", TeamID: "TEAM123"}, path)
	if err != nil {
		t.Fatal(err)
	}
	token, err := signer.Token()
	if err != nil {
		t.Fatal(err)
	}
	parts := splitToken(t, token)
	var claims map[string]any
	decodeJSONPart(t, parts[1], &claims)
	if claims["iss"] != "TEAM123" {
		t.Fatalf("iss = %v, want team ID fallback", claims["iss"])
	}
	if _, exists := claims["aud"]; exists {
		t.Fatalf("audience should be omitted when unset: %#v", claims)
	}
}

func TestTokenCachingAndEarlyRefresh(t *testing.T) {
	_, keyPEM := testKey(t)
	now := time.Date(2026, 11, 26, 12, 0, 0, 0, time.UTC)
	signer, err := New(Config{
		KeyID: "KEY123", Issuer: "issuer", TTL: 10 * time.Minute,
		RefreshBefore: 2 * time.Minute, Now: func() time.Time { return now },
	}, keyPEM)
	if err != nil {
		t.Fatal(err)
	}
	first, err := signer.Token()
	if err != nil {
		t.Fatal(err)
	}
	second, err := signer.Token()
	if err != nil || second != first {
		t.Fatalf("token not reused: err=%v", err)
	}
	now = now.Add(8 * time.Minute)
	third, err := signer.Token()
	if err != nil || third == first {
		t.Fatalf("token not refreshed at early-refresh boundary: err=%v", err)
	}
}

func TestConcurrentTokenCallsShareCachedToken(t *testing.T) {
	_, keyPEM := testKey(t)
	signer, err := New(Config{KeyID: "KEY123", Issuer: "issuer"}, keyPEM)
	if err != nil {
		t.Fatal(err)
	}
	const workers = 24
	tokens := make([]string, workers)
	errs := make([]error, workers)
	var wg sync.WaitGroup
	for i := range workers {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			tokens[i], errs[i] = signer.Token()
		}(i)
	}
	wg.Wait()
	for _, err := range errs {
		if err != nil {
			t.Fatal(err)
		}
	}
	for _, token := range tokens[1:] {
		if token != tokens[0] {
			t.Fatal("concurrent Token calls generated different cached tokens")
		}
	}
}

func TestParsePrivateKeyVariantsAndErrors(t *testing.T) {
	key, pemKey := testKey(t)
	parsed, err := ParsePrivateKey(pemKey)
	if err != nil || !samePrivateKey(t, parsed, key) {
		t.Fatalf("parse PKCS#8 PEM: err=%v", err)
	}
	sec1, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		t.Fatal(err)
	}
	parsed, err = ParsePrivateKey(pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: sec1}))
	if err != nil || !samePrivateKey(t, parsed, key) {
		t.Fatalf("parse SEC1 PEM: err=%v", err)
	}
	if _, err := ParsePrivateKey([]byte("not a PEM key")); err == nil {
		t.Fatal("expected invalid key error")
	}
	p384, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	p384DER, err := x509.MarshalPKCS8PrivateKey(p384)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := ParsePrivateKey(p384DER); err == nil {
		t.Fatal("expected non-P-256 key error")
	}
}

func TestNewRequiresSigningMetadata(t *testing.T) {
	_, keyPEM := testKey(t)
	for name, config := range map[string]Config{
		"missing key ID":          {Issuer: "issuer"},
		"missing issuer and team": {KeyID: "key"},
	} {
		t.Run(name, func(t *testing.T) {
			if _, err := New(config, keyPEM); err == nil {
				t.Fatal("expected validation error")
			}
		})
	}
}

func testKey(t *testing.T) (*ecdsa.PrivateKey, []byte) {
	t.Helper()
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	der, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		t.Fatal(err)
	}
	return key, pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: der})
}

func samePrivateKey(t *testing.T, left, right *ecdsa.PrivateKey) bool {
	t.Helper()
	leftDER, err := x509.MarshalPKCS8PrivateKey(left)
	if err != nil {
		t.Fatal(err)
	}
	rightDER, err := x509.MarshalPKCS8PrivateKey(right)
	if err != nil {
		t.Fatal(err)
	}
	return bytes.Equal(leftDER, rightDER)
}

func splitToken(t *testing.T, token string) []string {
	t.Helper()
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		t.Fatalf("JWT has %d parts", len(parts))
	}
	return parts
}

func decodeJSONPart(t *testing.T, part string, target any) {
	t.Helper()
	decoded, err := base64.RawURLEncoding.DecodeString(part)
	if err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(decoded, target); err != nil {
		t.Fatal(err)
	}
}
