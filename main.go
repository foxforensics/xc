// go-template.
//
// Usage:
//
//	go-template arg
//
// The arguments are:
//
//	arg
//	    Argument (required).
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 || os.Args[1] == "--help" {
		_, _ = fmt.Fprintln(os.Stderr, "usage: go-template arg")
		os.Exit(2)
	}
}
