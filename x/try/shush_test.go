package try

import (
	"bytes"
	"errors"
	"io"
	"log"
	"log/slog"
	"os"
	"strings"
	"testing"
)

func TestShush_NoErrorProducesNoOutput(t *testing.T) {
	var buf bytes.Buffer
	Shush(&buf)(func() {})
	if buf.Len() != 0 {
		t.Fatalf("expected no output, got %q", buf.String())
	}
}

func TestShush_IOWriterTarget(t *testing.T) {
	var buf bytes.Buffer
	boom := errors.New("boom")

	Shush(&buf)(func() { panic(boom) })

	got := buf.String()
	if !strings.Contains(got, "Shush: suppressed error") || !strings.Contains(got, "boom") {
		t.Fatalf("output = %q, want it to mention the suppressed error", got)
	}
}

func TestShush_SlogLoggerTarget(t *testing.T) {
	var buf bytes.Buffer
	logger := slog.New(slog.NewTextHandler(&buf, nil))
	boom := errors.New("boom")

	Shush(logger)(func() { panic(boom) })

	got := buf.String()
	if !strings.Contains(got, "level=WARN") {
		t.Fatalf("output = %q, want a WARN level entry", got)
	}
	if !strings.Contains(got, "Shush: suppressed error") || !strings.Contains(got, "boom") {
		t.Fatalf("output = %q, want it to mention the suppressed error", got)
	}
}

func TestShush_LogLoggerTarget(t *testing.T) {
	var buf bytes.Buffer
	logger := log.New(&buf, "", 0)
	boom := errors.New("boom")

	Shush(logger)(func() { panic(boom) })

	got := buf.String()
	if !strings.Contains(got, "Shush: suppressed error") || !strings.Contains(got, "boom") {
		t.Fatalf("output = %q, want it to mention the suppressed error", got)
	}
}

func TestShush_NilTargetDiscardsSilently(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("Shush(nil) panicked: %v", r)
		}
	}()
	Shush(nil)(func() { panic(errors.New("boom")) })
}

func TestShush_UnknownTargetFallsBackToStderr(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("os.Pipe() error: %v", err)
	}
	orig := os.Stderr
	os.Stderr = w

	boom := errors.New("boom")
	Shush(42)(func() { panic(boom) })

	os.Stderr = orig
	if err := w.Close(); err != nil {
		t.Fatalf("close pipe writer: %v", err)
	}
	out, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("read pipe: %v", err)
	}

	got := string(out)
	if !strings.Contains(got, "Shush: suppressed error") || !strings.Contains(got, "boom") {
		t.Fatalf("stderr output = %q, want it to mention the suppressed error", got)
	}
}
