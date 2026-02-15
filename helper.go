package machineid

import (
	"context"
	"io"
	"os"
	"strings"
	"time"

	exec "golang.org/x/sys/execabs"
)

// run wraps `exec.Command` with easy access to stdout and stderr.
func run(stdout, stderr io.Writer, cmd string, args ...string) error {
	// Resolve executable to reduce PATH shadowing risk.
	path, err := exec.LookPath(cmd)
	if err != nil {
		return err
	}

	// Prevent indefinite hangs.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
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
	return strings.TrimSpace(s)
}
