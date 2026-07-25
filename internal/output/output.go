// Package output provides consistent, opinionated terminal output functions
// for DevOpsHub CLI. All output respects global --no-color and --quiet flags.
//
// This is the single source of truth for all user-facing text rendering.
// Commands MUST NOT call fmt.Println directly for user-visible output.
package output

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

// ANSI escape codes for terminal colors and formatting.
// Only emitted when color is enabled (isColorEnabled returns true).
const (
	colorReset  = "\033[0m"
	colorBold   = "\033[1m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorCyan   = "\033[36m"
	colorGray   = "\033[90m"
	colorWhite  = "\033[97m"
)

// Printer holds rendering configuration for a single output context.
// Use New() or Default() to obtain a Printer. The zero value is not safe to use.
type Printer struct {
	out      io.Writer
	errOut   io.Writer
	noColor  bool
	jsonMode bool
	quiet    bool
}

// Option is a functional option for configuring a Printer.
type Option func(*Printer)

// WithNoColor disables ANSI color sequences.
func WithNoColor(v bool) Option { return func(p *Printer) { p.noColor = v } }

// WithJSON enables structured JSON output mode.
func WithJSON(v bool) Option { return func(p *Printer) { p.jsonMode = v } }

// WithQuiet suppresses informational messages; only errors are shown.
func WithQuiet(v bool) Option { return func(p *Printer) { p.quiet = v } }

// WithWriter overrides the standard output writer.
func WithWriter(w io.Writer) Option { return func(p *Printer) { p.out = w } }

// WithErrWriter overrides the error output writer.
func WithErrWriter(w io.Writer) Option { return func(p *Printer) { p.errOut = w } }

// New returns a Printer configured with the provided options.
func New(opts ...Option) *Printer {
	p := &Printer{
		out:    os.Stdout,
		errOut: os.Stderr,
	}
	for _, opt := range opts {
		opt(p)
	}
	return p
}

// Default is a package-level Printer using os.Stdout / os.Stderr.
// Commands can call output.Default() to get a zero-config printer,
// or override with CLI flag values via New().
func Default() *Printer {
	return New()
}

// color wraps text in an ANSI escape sequence if color is enabled.
func (p *Printer) color(code, text string) string {
	if p.noColor {
		return text
	}
	return code + text + colorReset
}

// bold wraps text in bold if color is enabled.
func (p *Printer) bold(text string) string {
	return p.color(colorBold, text)
}

// Success prints a green success message prefixed with ✓.
func (p *Printer) Success(msg string) {
	if p.quiet {
		return
	}
	fmt.Fprintln(p.out, p.color(colorGreen, "✓ ")+msg)
}

// Error prints a red error message prefixed with ✗ to stderr.
func (p *Printer) Error(msg string) {
	fmt.Fprintln(p.errOut, p.color(colorRed, "✗ ")+msg)
}

// Info prints a blue informational message prefixed with ℹ.
func (p *Printer) Info(msg string) {
	if p.quiet {
		return
	}
	fmt.Fprintln(p.out, p.color(colorBlue, "ℹ ")+msg)
}

// Warning prints a yellow warning message prefixed with ⚠.
func (p *Printer) Warning(msg string) {
	if p.quiet {
		return
	}
	fmt.Fprintln(p.out, p.color(colorYellow, "⚠ ")+msg)
}

// Header prints a bold, cyan section header.
func (p *Printer) Header(title string) {
	if p.quiet {
		return
	}
	line := strings.Repeat("─", len(title)+4)
	fmt.Fprintln(p.out, p.color(colorCyan, line))
	fmt.Fprintln(p.out, p.color(colorCyan, "  ")+p.bold(title))
	fmt.Fprintln(p.out, p.color(colorCyan, line))
}

// Println writes a plain line. Respects quiet mode.
func (p *Printer) Println(msg string) {
	if p.quiet {
		return
	}
	fmt.Fprintln(p.out, msg)
}

// Printf writes a formatted line. Respects quiet mode.
func (p *Printer) Printf(format string, args ...any) {
	if p.quiet {
		return
	}
	fmt.Fprintf(p.out, format, args...)
}

// JSONMode reports whether this printer is configured for JSON output.
// Commands should call this before deciding between human-readable and
// structured output, rather than attempting to infer mode from behavior.
func (p *Printer) JSONMode() bool {
	return p.jsonMode
}

// JSON marshals v and writes it as pretty-printed JSON to stdout.
// Returns an error if marshaling fails.
func (p *Printer) JSON(v any) error {
	enc := json.NewEncoder(p.out)
	enc.SetIndent("", "  ")
	return enc.Encode(v)
}

// Table renders a simple table with a header row and data rows.
// Column widths are auto-calculated based on content.
func (p *Printer) Table(headers []string, rows [][]string) {
	if p.quiet {
		return
	}

	// Calculate column widths.
	widths := make([]int, len(headers))
	for i, h := range headers {
		widths[i] = len(h)
	}
	for _, row := range rows {
		for i, cell := range row {
			if i < len(widths) && len(cell) > widths[i] {
				widths[i] = len(cell)
			}
		}
	}

	// Build separator line.
	sep := "+"
	for _, w := range widths {
		sep += strings.Repeat("─", w+2) + "+"
	}

	// Print header.
	fmt.Fprintln(p.out, p.color(colorGray, sep))
	headerRow := "|"
	for i, h := range headers {
		headerRow += " " + p.bold(fmt.Sprintf("%-*s", widths[i], h)) + " |"
	}
	fmt.Fprintln(p.out, headerRow)
	fmt.Fprintln(p.out, p.color(colorGray, sep))

	// Print data rows.
	for _, row := range rows {
		line := "|"
		for i, cell := range row {
			if i < len(widths) {
				line += fmt.Sprintf(" %-*s |", widths[i], cell)
			}
		}
		fmt.Fprintln(p.out, line)
	}
	fmt.Fprintln(p.out, p.color(colorGray, sep))
}

// Section prints a dimmed, labeled section divider.
func (p *Printer) Section(label string) {
	if p.quiet {
		return
	}
	fmt.Fprintln(p.out, p.color(colorGray, "\n  "+label))
	fmt.Fprintln(p.out, p.color(colorGray, "  "+strings.Repeat("·", len(label))))
}

// Command prints a formatted command example with a description.
func (p *Printer) Command(cmd, description string) {
	if p.quiet {
		return
	}
	fmt.Fprintf(p.out, "  %s  %s\n",
		p.color(colorCyan, fmt.Sprintf("%-45s", cmd)),
		p.color(colorGray, description),
	)
}
