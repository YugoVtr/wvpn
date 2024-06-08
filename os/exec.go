package os

import (
	"bytes"
	"fmt"
	"os/exec"
)

// Commander is a function that executes a command and returns the stdout and error.
// If stderr is not empty, it will be returned as an error.
func Command(c string, args ...string) (string, error) {
	cmd := exec.Command(c, args...)
	stdout, stderr := &bytes.Buffer{}, &bytes.Buffer{}
	cmd.Stdout, cmd.Stderr = stdout, stderr
	if err := cmd.Run(); err != nil {
		return stdout.String(), fmt.Errorf("%s\n%w", stderr.String(), err)
	}
	if stderr.Len() > 0 {
		return stdout.String(), fmt.Errorf(stderr.String())
	}
	return stdout.String(), nil
}

// Print prints the output of a command and an error if it exists.
func Print(out string, err error) {
	if err != nil {
		fmt.Printf("failed: %s with error %s", out, err)
	}
	fmt.Print(out)
}
