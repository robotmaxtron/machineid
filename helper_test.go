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

func Test_trim(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "nil",
			args: args{s: ""},
			want: "",
		},
		{
			name: "space",
			args: args{s: " space "},
			want: "space",
		},
		{
			name: "nl",
			args: args{s: "data\n"},
			want: "data",
		},
		{
			name: "combined",
			args: args{s: " some data \n"},
			want: "some data",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trim(tt.args.s); got != tt.want {
				t.Errorf("trim() = %v, want %v", got, tt.want)
			}
		})
	}
}
