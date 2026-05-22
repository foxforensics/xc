// Experimental terminal chain-of-custody.
//
// Usage:
//
//	xc COMMAND ...
package main

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const Format = `%s
--
%s
--
%x`

func main() {
	defer func() {
		if err := recover(); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "xc:", err)
			os.Exit(1)
		}
	}()

	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stdin = os.Stdin // optional input

	b, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Print(string(b)) // mirror error

		if ex, ok := errors.AsType[*exec.ExitError](err); ok {
			os.Exit(ex.ExitCode())
		}
	}

	fmt.Printf(Format, strings.Join(os.Args[1:], " "), b, sha256.Sum256(b))
}
