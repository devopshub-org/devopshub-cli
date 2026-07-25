package output_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/devopshub-org/devopshub-cli/internal/output"
)

// newTestPrinter returns a Printer that writes to a bytes.Buffer for assertion.
func newTestPrinter(opts ...output.Option) (p *output.Printer, outBuf *bytes.Buffer, errOutBuf *bytes.Buffer) {
	var out, errOut bytes.Buffer
	base := []output.Option{
		output.WithWriter(&out),
		output.WithErrWriter(&errOut),
		output.WithNoColor(true), // disable ANSI in tests for clean assertions
	}
	p := output.New(append(base, opts...)...)
	return p, &out, &errOut
}

func TestSuccess_WritesToStdout(t *testing.T) {
	p, out, _ := newTestPrinter()
	p.Success("operation completed")
	got := out.String()
	if !strings.Contains(got, "operation completed") {
		t.Errorf("Success() output = %q; want to contain 'operation completed'", got)
	}
}

func TestSuccess_QuietMode_WritesNothing(t *testing.T) {
	p, out, _ := newTestPrinter(output.WithQuiet(true))
	p.Success("should be suppressed")
	if out.String() != "" {
		t.Errorf("Success() in quiet mode wrote %q; want empty", out.String())
	}
}

func TestError_WritesToStderr(t *testing.T) {
	p, _, errOut := newTestPrinter()
	p.Error("something went wrong")
	got := errOut.String()
	if !strings.Contains(got, "something went wrong") {
		t.Errorf("Error() errOut = %q; want to contain 'something went wrong'", got)
	}
}

func TestError_QuietMode_StillWritesToStderr(t *testing.T) {
	// Errors must always be shown, even in quiet mode.
	p, _, errOut := newTestPrinter(output.WithQuiet(true))
	p.Error("critical failure")
	if !strings.Contains(errOut.String(), "critical failure") {
		t.Error("Error() must write to stderr even in quiet mode")
	}
}

func TestInfo_WritesToStdout(t *testing.T) {
	p, out, _ := newTestPrinter()
	p.Info("cluster connected")
	if !strings.Contains(out.String(), "cluster connected") {
		t.Errorf("Info() output = %q; want to contain 'cluster connected'", out.String())
	}
}

func TestWarning_WritesToStdout(t *testing.T) {
	p, out, _ := newTestPrinter()
	p.Warning("resource limit approaching")
	if !strings.Contains(out.String(), "resource limit approaching") {
		t.Errorf("Warning() output = %q; want to contain message", out.String())
	}
}

func TestHeader_IncludesTitle(t *testing.T) {
	p, out, _ := newTestPrinter()
	p.Header("Docker Commands")
	got := out.String()
	if !strings.Contains(got, "Docker Commands") {
		t.Errorf("Header() output = %q; want to contain 'Docker Commands'", got)
	}
}

func TestPrintln_WritesLine(t *testing.T) {
	p, out, _ := newTestPrinter()
	p.Println("hello world")
	if !strings.Contains(out.String(), "hello world") {
		t.Errorf("Println() output = %q; want 'hello world'", out.String())
	}
}

func TestPrintf_WritesFormattedOutput(t *testing.T) {
	p, out, _ := newTestPrinter()
	p.Printf("count: %d", 42)
	if !strings.Contains(out.String(), "count: 42") {
		t.Errorf("Printf() output = %q; want 'count: 42'", out.String())
	}
}

func TestJSON_WritesValidJSON(t *testing.T) {
	p, out, _ := newTestPrinter(output.WithJSON(true))
	payload := map[string]string{"key": "value"}
	if err := p.JSON(payload); err != nil {
		t.Fatalf("JSON() returned error: %v", err)
	}
	got := out.String()
	if !strings.Contains(got, `"key"`) || !strings.Contains(got, `"value"`) {
		t.Errorf("JSON() output = %q; want valid JSON with key/value", got)
	}
}

func TestTable_RendersHeaderAndRows(t *testing.T) {
	p, out, _ := newTestPrinter()
	headers := []string{"COMMAND", "DESCRIPTION"}
	rows := [][]string{
		{"docker ps", "list running containers"},
		{"docker images", "list images"},
	}
	p.Table(headers, rows)
	got := out.String()
	for _, want := range []string{"COMMAND", "DESCRIPTION", "docker ps", "docker images", "list running containers"} {
		if !strings.Contains(got, want) {
			t.Errorf("Table() output does not contain %q\ngot: %s", want, got)
		}
	}
}

func TestNew_Default_DoesNotPanic(t *testing.T) {
	// Default printer must be creatable without options.
	p := output.Default()
	if p == nil {
		t.Fatal("Default() returned nil")
	}
}
