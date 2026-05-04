// Package output provides output formatting for CLI applications.
// Handles table output, JSON formatting, color detection, and terminal width.
//
// Example usage:
//
//	formatter := output.New("table", "auto", false)
//	tbl := formatter.NewTable("Name", "Value")
//	tbl.AddRow("item1", "data1").AddRow("item2", "data2")
//	tbl.Print()
package output

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/mattn/go-isatty"
	"github.com/rodaine/table"
)

// Formatter handles output formatting
type Formatter struct {
	Format    string
	Color     bool
	NoHeaders bool
}

// New creates a new output formatter
func New(format, color string, noHeaders bool) *Formatter {
	useColor := false
	switch color {
	case "always":
		useColor = true
	case "never":
		useColor = false
	case "auto":
		useColor = isatty.IsTerminal(os.Stdout.Fd())
	}

	return &Formatter{
		Format:    format,
		Color:     useColor,
		NoHeaders: noHeaders,
	}
}

// PrintJSON outputs data as formatted JSON
func (f *Formatter) PrintJSON(data interface{}) error {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(data)
}

// ValidateFormat checks if the format is supported
func ValidateFormat(format string) error {
	switch format {
	case "json", "table":
		return nil
	default:
		return fmt.Errorf("unsupported format: %s (supported: json, table)", format)
	}
}

// TableBuilder helps construct tables with consistent formatting
type TableBuilder struct {
	formatter *Formatter
	tbl       table.Table
	headers   []string
}

// NewTable creates a new table builder
func (f *Formatter) NewTable(headers ...string) *TableBuilder {
	columns := make([]interface{}, len(headers))
	for i, header := range headers {
		columns[i] = header
	}

	tbl := table.New(columns...).WithWriter(os.Stdout)

	if f.Color && !f.NoHeaders {
		tbl.WithHeaderFormatter(func(format string, vals ...interface{}) string {
			return fmt.Sprintf("\033[1m%s\033[0m", fmt.Sprintf(format, vals...))
		})
	}

	return &TableBuilder{
		formatter: f,
		tbl:       tbl,
		headers:   headers,
	}
}

// AddRow adds a row to the table
func (tb *TableBuilder) AddRow(row ...interface{}) *TableBuilder {
	tb.tbl.AddRow(row...)
	return tb
}

// Print outputs the table (or raw data if NoHeaders without color)
func (tb *TableBuilder) Print() {
	if len(tb.headers) == 0 {
		fmt.Println("No data to display.")
		return
	}

	if !tb.formatter.NoHeaders || tb.formatter.Color {
		tb.tbl.Print()
	} else {
		// Raw tab-separated output
		// This is handled by the caller for specific data types
	}
}

// PrintWithFallback prints with fallback to raw output
func (tb *TableBuilder) PrintWithFallback(printRaw func()) {
	if !tb.formatter.NoHeaders || tb.formatter.Color {
		tb.tbl.Print()
	} else {
		printRaw()
	}
}

// IsTerminal returns true if stdout is a terminal
func IsTerminal() bool {
	return isatty.IsTerminal(os.Stdout.Fd())
}

// TruncateString truncates a string to max length with ellipsis
func TruncateString(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max-3] + "..."
}

// PrintEmptyMessage prints a standard "No X found" message
func PrintEmptyMessage(itemType string) {
	fmt.Printf("No %s found.\n", itemType)
}

// PrintSuccess prints a standard success message
func PrintSuccess(format string, args ...interface{}) {
	fmt.Printf("✓ "+format+"\n", args...)
}

// PrintError prints a standard error message
func PrintError(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "✗ "+format+"\n", args...)
}

// PrintWarning prints a standard warning message
func PrintWarning(format string, args ...interface{}) {
	fmt.Printf("⚠ "+format+"\n", args...)
}

// PrintBanner prints a standard banner (for update notifications, etc.)
func PrintBanner(title string) {
	width := 62
	titleLen := len(title)
	padding := (width - titleLen - 2) / 2

	border := strings.Repeat("═", width)
	fmt.Println("╔" + border + "╗")
	fmt.Printf("║%s %s %s║\n", strings.Repeat(" ", padding), title, strings.Repeat(" ", width-titleLen-padding-2))
	fmt.Println("╚" + border + "╝")
}
