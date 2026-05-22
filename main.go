// Witness terminal commands to build a chain-of-custody.
//
// Usage:
//
//	witness command [arg ...]
//
// The arguments are:
//
//	command
//		The command to execute (required).
//	arg
//	    The arguments to pass (optional).
package main

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

const buffer = 4096 // one page

func read(r io.ReadCloser, s string) {
	buf := make([]byte, buffer)
	sha := sha256.New()

	defer func(r io.Closer) {
		fmt.Printf("\n%v %x SHA256\n", s, sha.Sum(nil))

		if err := r.Close(); err != nil {
			log.Fatal(err)
		}
	}(r)

	for {
		switch n, err := r.Read(buf); {
		case errors.Is(err, io.EOF):
			return
		case err != nil:
			log.Fatal(err)
		default:
			sha.Sum(buf[:n])
			fmt.Print(string(buf[:n]))
		}
	}
}

func main() {
	if len(os.Args) == 1 || os.Args[1] == "--help" {
		_, _ = fmt.Fprintln(os.Stderr, "usage: witness COMMAND [ARG ...]")
		os.Exit(2)
	}

	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stdin = os.Stdin

	// cmd.CombinedOutput()

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		log.Fatal(err)
	}

	stderr, err := cmd.StderrPipe()

	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	go read(stdout, "stdout")
	go read(stderr, "stderr")

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}
