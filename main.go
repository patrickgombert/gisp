package main

import (
	"fmt"
	"os"

	ds "github.com/patrickgombert/gisp/datastructures"
	"github.com/patrickgombert/gisp/function"
	p "github.com/patrickgombert/gisp/parser"
	r "github.com/patrickgombert/gisp/runtime"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: gisp <file>\n")
		return
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error %q\n", err)
		os.Exit(1)
		return
	}

	program, parseErr := p.FromString(string(data))
	if parseErr != nil {
		fmt.Fprintf(os.Stderr, "error %q\n", parseErr)
		os.Exit(1)
		return
	}

	if list, ok := program.(*ds.List); ok {
		fmt.Printf("%s", function.Eval(list, r.DefaultEnvironment()))
	}
}
