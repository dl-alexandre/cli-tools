// Package appleauth creates cached ES256 bearer tokens for Apple APIs.
package appleauth

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	defaultTTL          = 15 * time.Minute
	defaultRefreshEarly = time.Minute
)

// Config describes the Apple token claims and signing key metadata. Issuer is
// written to the JWT iss claim; TeamID is used as the issuer when Issuer is
// empty. TTL defaults to 15 minutes and RefreshBefore defaults to one minute.
type Config struct {
	KeyID         string
	Issuer        string
	TeamID        string
	Audience      string
	TTL           time.Duration
	RefreshBefore time.Duration
	Now           func() time.Time
}

// Signer signs ES256 tokens and caches each token until its early-refresh
// window begins. A Signer is safe for concurrent use.
type Signer struct {
	config Config
	key    *ecdsa.PrivateKey

	mu        sync.Mutex
	token     string
	refreshAt time.Time
}

// New parses a PEM-encoded ECDSA P-256 key and creates a signer.
func New(config Config, privateKeyPEM []byte) (*Signer, error) {
	key, err := ParsePrivateKey(privateKeyPEM)
	if err != nil {
		return nil, err
	}
	return newSigner(config, key)
}

// NewFromFile reads and parses a PEM-encoded ECDSA P-256 key file.
func NewFromFile(config Config, path string) (*Signer, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read Apple private key: %w", err)
	}
	return New(config, data)
}

// ParsePrivateKey accepts PEM PKCS#8 and SEC1 ECDSA private keys, plus their
// DER encodings. Only P-256 keys are valid for ES256.
func ParsePrivateKey(data []byte) (*ecdsa.PrivateKey, error) {
	block, _ := pem.Decode(data)
	if block != nil {
		data = block.Bytes
	}
	key, err := x509.ParsePKCS8PrivateKey(data)
	if err != nil {
		key, err = x509.ParseECPrivateKey(data)
	}
	if err != nil {
		return nil, fmt.Errorf("parse Apple private key: %w", err)
	}
	ecdsaKey, ok := key.(*ecdsa.PrivateKey)
	if !ok {
		return nil, errors.New("apple private key must be ECDSA")
	}
	if ecdsaKey.Curve.Params().Name != "P-256" {
		return nil, errors.New("apple ES256 private key must use the P-256 curve")
	}
	return ecdsaKey, nil
}

func newSigner(config Config, key *ecdsa.PrivateKey) (*Signer, error) {
	config.KeyID = strings.TrimSpace(config.KeyID)
	config.Issuer = strings.TrimSpace(config.Issuer)
	config.TeamID = strings.TrimSpace(config.TeamID)
	config.Audience = strings.TrimSpace(config.Audience)
	if config.KeyID == "" {
		return nil, errors.New("apple key ID is required")
	}
	if config.Issuer == "" {
		config.Issuer = config.TeamID
	}
	if config.Issuer == "" {
		return nil, errors.New("apple issuer or team ID is required")
	}
	if config.TTL <= 0 {
		config.TTL = defaultTTL
	}
	if config.RefreshBefore <= 0 || config.RefreshBefore >= config.TTL {
		config.RefreshBefore = min(defaultRefreshEarly, config.TTL/10)
	}
	if config.Now == nil {
		config.Now = time.Now
	}
	return &Signer{config: config, key: key}, nil
}

// Token returns a cached JWT or creates a new one when the token is within its
// early-refresh window.
func (s *Signer) Token() (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := s.config.Now().UTC()
	if s.token != "" && now.Before(s.refreshAt) {
		return s.token, nil
	}
	token, err := s.sign(now)
	if err != nil {
		return "", err
	}
	s.token = token
	s.refreshAt = now.Add(s.config.TTL - s.config.RefreshBefore)
	return token, nil
}

func (s *Signer) sign(now time.Time) (string, error) {
	header := map[string]string{"alg": "ES256", "typ": "JWT", "kid": s.config.KeyID}
	claims := map[string]any{
		"iss": s.config.Issuer,
		"iat": now.Unix(),
		"exp": now.Add(s.config.TTL).Unix(),
	}
	if s.config.Audience != "" {
		claims["aud"] = s.config.Audience
	}
	headerJSON, err := json.Marshal(header)
	if err != nil {
		return "", fmt.Errorf("encode Apple token header: %w", err)
	}
	claimsJSON, err := json.Marshal(claims)
	if err != nil {
		return "", fmt.Errorf("encode Apple token claims: %w", err)
	}
	encode := base64.RawURLEncoding.EncodeToString
	input := encode(headerJSON) + "." + encode(claimsJSON)
	hash := sha256.Sum256([]byte(input))
	r, sigS, err := ecdsa.Sign(rand.Reader, s.key, hash[:])
	if err != nil {
		return "", fmt.Errorf("sign Apple token: %w", err)
	}
	signature := make([]byte, 64)
	r.FillBytes(signature[:32])
	sigS.FillBytes(signature[32:])
	return input + "." + encode(signature), nil
}
