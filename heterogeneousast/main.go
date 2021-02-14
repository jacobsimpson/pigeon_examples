//go:generate pigeon -o grammar.go grammar.peg
//go:generate goimports -w grammar.go
package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kr/pretty"
	flag "github.com/spf13/pflag"
)

var debug = flag.BoolP("debug", "d", false, "show additional debugging information for the parser")

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage: %s [OPTION]... <FILE>
Parse FILE and write something, anything, to standard out.

`, filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		flag.Usage()
		os.Exit(1)
	}

	ast, err := ParseFile(flag.Arg(0), Debug(*debug))
	expression := ast.(*Expression)

	fmt.Printf("expression    = %#v\n", expression)
	fmt.Printf("expression    = %# v", pretty.Formatter(expression))
	fmt.Printf("err           = %+v\n", err)
}
