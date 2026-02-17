package machineid

import (
	"bytes"
	"context"
	"io"
	"os"
	"os/exec"
	"time"
)

// run wraps `exec.Command` with easy access to stdout and stderr.
func run(stdout, stderr io.Writer, cmd string, args ...string) error {
	// Resolve executable to reduce PATH shadowing risk.
	path, err := exec.LookPath(cmd)
	if err != nil {
		return err
	}

	// Prevent indefinite hangs.
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	c := exec.CommandContext(ctx, path, args...)
	c.Stdin = nil // no stdin needed
	c.Stdout = stdout
	c.Stderr = stderr
	return c.Run()
}

func readFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

func trim(s string) string {
	return string(bytes.TrimSpace([]byte(s)))
}
