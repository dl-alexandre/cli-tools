package output

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name        string
		format      string
		color       string
		noHeaders   bool
		wantColor   bool
		wantFormat  string
		wantHeaders bool
	}{
		{
			name:        "table auto color",
			format:      "table",
			color:       "auto",
			noHeaders:   false,
			wantColor:   IsTerminal(), // Depends on terminal
			wantFormat:  "table",
			wantHeaders: false,
		},
		{
			name:        "json never color",
			format:      "json",
			color:       "never",
			noHeaders:   true,
			wantColor:   false,
			wantFormat:  "json",
			wantHeaders: true,
		},
		{
			name:        "table always color",
			format:      "table",
			color:       "always",
			noHeaders:   false,
			wantColor:   true,
			wantFormat:  "table",
			wantHeaders: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := New(tt.format, tt.color, tt.noHeaders)
			if f.Format != tt.wantFormat {
				t.Errorf("Format = %q, want %q", f.Format, tt.wantFormat)
			}
			if tt.color == "always" && !f.Color {
				t.Error("Color should be true when color='always'")
			}
			if tt.color == "never" && f.Color {
				t.Error("Color should be false when color='never'")
			}
			if f.NoHeaders != tt.wantHeaders {
				t.Errorf("NoHeaders = %v, want %v", f.NoHeaders, tt.wantHeaders)
			}
		})
	}
}

func TestPrintJSON(t *testing.T) {
	tests := []struct {
		name     string
		data     interface{}
		expected map[string]interface{}
	}{
		{
			name: "simple map",
			data: map[string]string{"key": "value"},
			expected: map[string]interface{}{
				"key": "value",
			},
		},
		{
			name: "nested struct",
			data: struct {
				Name    string `json:"name"`
				Version string `json:"version"`
			}{
				Name:    "test",
				Version: "v1.0.0",
			},
			expected: map[string]interface{}{
				"name":    "test",
				"version": "v1.0.0",
			},
		},
		{
			name:     "slice",
			data:     []string{"a", "b", "c"},
			expected: nil, // Slices are different, just check for no error
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := New("json", "never", false)

			// Capture stdout
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			err := f.PrintJSON(tt.data)

			w.Close()
			os.Stdout = old

			if err != nil {
				t.Errorf("PrintJSON() error = %v", err)
				return
			}

			var buf bytes.Buffer
			io.Copy(&buf, r)
			output := buf.String()

			// Verify it's valid JSON
			var result interface{}
			if err := json.Unmarshal([]byte(output), &result); err != nil {
				t.Errorf("Output is not valid JSON: %v", err)
			}

			// If expected is not nil, verify it matches
			if tt.expected != nil {
				resultMap, ok := result.(map[string]interface{})
				if !ok {
					t.Error("Expected map result")
					return
				}
				for key, val := range tt.expected {
					if resultMap[key] != val {
						t.Errorf("Expected %q = %v, got %v", key, val, resultMap[key])
					}
				}
			}
		})
	}
}

func TestValidateFormat(t *testing.T) {
	tests := []struct {
		format  string
		wantErr bool
	}{
		{"json", false},
		{"table", false},
		{"yaml", true},
		{"csv", true},
		{"", true},
		{"JSON", true}, // Case sensitive
	}

	for _, tt := range tests {
		t.Run(tt.format, func(t *testing.T) {
			err := ValidateFormat(tt.format)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateFormat(%q) error = %v, wantErr %v", tt.format, err, tt.wantErr)
			}
		})
	}
}

func TestTableBuilder(t *testing.T) {
	f := New("table", "never", false)
	tbl := f.NewTable("Name", "Value", "Status")

	// Test AddRow returns same builder for chaining
	tbl2 := tbl.AddRow("item1", "data1", "active")
	if tbl != tbl2 {
		t.Error("AddRow should return same TableBuilder for chaining")
	}

	// Test multiple rows
	tbl.AddRow("item2", "data2", "inactive")
	tbl.AddRow("item3", "data3", "pending")

	// We can't easily test the output without mocking the table library
	// but we can verify the structure was created
}

func TestIsTerminal(t *testing.T) {
	// This test will behave differently in different environments
	// In tests, it should typically return false
	result := IsTerminal()

	// Just verify it doesn't panic
	_ = result
}

func TestTruncateString(t *testing.T) {
	tests := []struct {
		input    string
		max      int
		expected string
	}{
		{"hello", 10, "hello"},
		{"hello world", 8, "hello..."},
		{"a", 5, "a"},
		{"exactly", 7, "exactly"},
		{"1234567890", 5, "12..."},
		{"", 10, ""},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s_%d", tt.input, tt.max), func(t *testing.T) {
			result := TruncateString(tt.input, tt.max)
			if result != tt.expected {
				t.Errorf("TruncateString(%q, %d) = %q, want %q", tt.input, tt.max, result, tt.expected)
			}
		})
	}
}

func TestPrintEmptyMessage(t *testing.T) {
	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	PrintEmptyMessage("users")

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := strings.TrimSpace(buf.String())

	expected := "No users found."
	if output != expected {
		t.Errorf("PrintEmptyMessage() = %q, want %q", output, expected)
	}
}

func TestPrintSuccess(t *testing.T) {
	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	PrintSuccess("Operation completed")

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := strings.TrimSpace(buf.String())

	if !strings.HasPrefix(output, "✓") {
		t.Errorf("PrintSuccess() should start with ✓, got %q", output)
	}
	if !strings.Contains(output, "Operation completed") {
		t.Errorf("PrintSuccess() should contain message, got %q", output)
	}
}

func TestPrintError(t *testing.T) {
	// Capture stderr
	old := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	PrintError("Something went wrong")

	w.Close()
	os.Stderr = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := strings.TrimSpace(buf.String())

	if !strings.HasPrefix(output, "✗") {
		t.Errorf("PrintError() should start with ✗, got %q", output)
	}
	if !strings.Contains(output, "Something went wrong") {
		t.Errorf("PrintError() should contain message, got %q", output)
	}
}

func TestPrintWarning(t *testing.T) {
	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	PrintWarning("This is a warning")

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := strings.TrimSpace(buf.String())

	if !strings.HasPrefix(output, "⚠") {
		t.Errorf("PrintWarning() should start with ⚠, got %q", output)
	}
	if !strings.Contains(output, "This is a warning") {
		t.Errorf("PrintWarning() should contain message, got %q", output)
	}
}

func TestPrintBanner(t *testing.T) {
	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	PrintBanner("UPDATE AVAILABLE")

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	// Should contain the title
	if !strings.Contains(output, "UPDATE AVAILABLE") {
		t.Errorf("PrintBanner() should contain title, got %q", output)
	}

	// Should contain box drawing characters
	if !strings.Contains(output, "╔") || !strings.Contains(output, "╗") {
		t.Errorf("PrintBanner() should contain box characters, got %q", output)
	}

	// Should contain the border line
	if !strings.Contains(output, "═") {
		t.Errorf("PrintBanner() should contain border line, got %q", output)
	}
}

func TestPrintBannerWithLongTitle(t *testing.T) {
	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	PrintBanner("A very long title that might overflow the box width")

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	// Should still print without panicking
	if !strings.Contains(output, "A very long") {
		t.Errorf("PrintBanner() should handle long titles, got %q", output)
	}
}

func TestTableBuilderPrintWithEmptyData(t *testing.T) {
	f := New("table", "never", false)
	tbl := f.NewTable()

	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	tbl.Print()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := strings.TrimSpace(buf.String())

	if output != "No data to display." {
		t.Errorf("Print() with no headers = %q, want %q", output, "No data to display.")
	}
}

func TestTableBuilderPrintWithFallback(t *testing.T) {
	f := New("table", "never", true) // noHeaders = true
	tbl := f.NewTable("Col1", "Col2")
	tbl.AddRow("a", "b")

	fallbackCalled := false

	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	tbl.PrintWithFallback(func() {
		fallbackCalled = true
		fmt.Println("fallback output")
	})

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if !fallbackCalled {
		t.Error("Fallback function should have been called when NoHeaders=true")
	}
	if !strings.Contains(output, "fallback output") {
		t.Errorf("Output should contain fallback output, got %q", output)
	}
}
