package machineid

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
	"testing"
)

func Test_run(t *testing.T) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	wantStdout := "hello"
	wantStderr := ""
	if err := run(stdout, stderr, "echo", "hello"); err != nil {
		t.Error(err)
	}
	gotStdout := strings.TrimRight(stdout.String(), "\r\n")
	if gotStdout != wantStdout {
		t.Errorf("run() = %v, want %v", gotStdout, wantStdout)
	}
	if gotStderr := stderr.String(); gotStderr != wantStderr {
		t.Errorf("run() = %v, want %v", gotStderr, wantStderr)
	}
}

func Test_run_unknown(t *testing.T) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	err := run(stdout, stderr, "echolo", "hello")
	if err == nil {
		t.Error("expected error, got none")
	}
	if !errors.Is(err, exec.ErrNotFound) {
		// Keep a fallback message for debugging, but don't assert on OS-specific strings.
		t.Fatalf("expected exec.ErrNotFound, got: %v", err)
	}
}
