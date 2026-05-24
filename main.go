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
	"time"
)

const Format = `[%s] %s
%s
SHA256 %x`

func main() {
	defer func() {
		if err := recover(); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "xc:", err)
			os.Exit(1)
		}
	}()

	if len(os.Args) >= 2 {
		now := time.Now().UTC()
		cmd := exec.Command(os.Args[1], os.Args[2:]...)
		cmd.Stdin = os.Stdin // optional input
		buf, err := cmd.CombinedOutput()

		if err != nil {
			fmt.Print(string(buf)) // mirror error
			if ex, ok := errors.AsType[*exec.ExitError](err); ok {
				os.Exit(ex.ExitCode())
			}
		}

		fmt.Printf(Format, now.Format(time.RFC3339), strings.Join(os.Args[1:], " "), buf, sha256.Sum256(buf))
	}
}
