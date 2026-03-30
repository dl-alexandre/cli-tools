package kongx

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/alecthomas/kong"
)

// Test CLI structure
type testCLI struct {
	CommonFlags

	List testListCmd `cmd:"" help:"List items"`
}

type testListCmd struct {
	All bool `help:"Show all items"`
}

func (c *testListCmd) Run(ctx *Context) error {
	return nil
}

func TestContext(t *testing.T) {
	ctx := Context{
		Config: map[string]string{"key": "value"},
		Debug:  true,
	}

	if ctx.Debug != true {
		t.Error("Debug should be true")
	}

	config, ok := ctx.Config.(map[string]string)
	if !ok {
		t.Fatal("Config should be map[string]string")
	}
	if config["key"] != "value" {
		t.Error("Config key should be 'value'")
	}
}

func TestCommonFlags(t *testing.T) {
	var cli testCLI

	// Parse with default values
	_, err := kong.New(&cli)
	if err != nil {
		t.Fatalf("Failed to create parser: %v", err)
	}

	// Verify default values
	if cli.Format != "table" {
		t.Errorf("Default Format = %q, want %q", cli.Format, "table")
	}
	if cli.Color != "auto" {
		t.Errorf("Default Color = %q, want %q", cli.Color, "auto")
	}
	if cli.NoHeaders != false {
		t.Errorf("Default NoHeaders = %v, want %v", cli.NoHeaders, false)
	}
	if cli.Debug != false {
		t.Errorf("Default Debug = %v, want %v", cli.Debug, false)
	}
}

func TestCommonUpdateFlags(t *testing.T) {
	flags := CommonUpdateFlags{}

	// Test default values
	if flags.Force != false {
		t.Errorf("Default Force = %v, want %v", flags.Force, false)
	}
	if flags.Format != "" {
		t.Errorf("Default Format = %q, want empty string", flags.Format)
	}
}

func TestErrorHandler(t *testing.T) {
	tests := []struct {
		name     string
		input    error
		expected error
		wantNil  bool
	}{
		{
			name:     "nil error",
			input:    nil,
			expected: nil,
			wantNil:  true,
		},
		{
			name:     "simple error",
			input:    errors.New("something went wrong"),
			expected: errors.New("✗ something went wrong"),
			wantNil:  false,
		},
		{
			name:     "wrapped error",
			input:    fmt.Errorf("outer: %w", errors.New("inner")),
			expected: nil,
			wantNil:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ErrorHandler(tt.input)
			if tt.wantNil && result != nil {
				t.Errorf("ErrorHandler() = %v, want nil", result)
			}
			if !tt.wantNil && result == nil {
				t.Errorf("ErrorHandler() = nil, want error")
			}
			if result != nil && !strings.HasPrefix(result.Error(), "✗") {
				t.Errorf("ErrorHandler() error should start with ✗, got %v", result)
			}
		})
	}
}

func TestHelpFormatter(t *testing.T) {
	result := HelpFormatter("myapp", "A test application")

	// Should contain app name
	if !strings.Contains(result, "myapp") {
		t.Errorf("HelpFormatter() should contain app name, got %q", result)
	}

	// Should contain description
	if !strings.Contains(result, "A test application") {
		t.Errorf("HelpFormatter() should contain description, got %q", result)
	}

	// Should contain usage section
	if !strings.Contains(result, "Usage:") {
		t.Errorf("HelpFormatter() should contain Usage section, got %q", result)
	}

	// Should contain flags section
	if !strings.Contains(result, "Flags:") {
		t.Errorf("HelpFormatter() should contain Flags section, got %q", result)
	}

	// Should contain commands section
	if !strings.Contains(result, "Commands:") {
		t.Errorf("HelpFormatter() should contain Commands section, got %q", result)
	}

	// Should contain Kong template placeholders
	if !strings.Contains(result, "{{flags .Flags}}") {
		t.Errorf("HelpFormatter() should contain flags template, got %q", result)
	}
}

func TestRunFunc(t *testing.T) {
	// Test that RunFunc can be assigned
	var fn RunFunc = func(ctx *Context) error {
		return nil
	}

	ctx := &Context{Debug: false}
	err := fn(ctx)
	if err != nil {
		t.Errorf("RunFunc() error = %v", err)
	}

	// Test error return
	fn2 := func(ctx *Context) error {
		return errors.New("test error")
	}

	err2 := fn2(ctx)
	if err2 == nil || err2.Error() != "test error" {
		t.Errorf("RunFunc() error = %v, want 'test error'", err2)
	}
}
