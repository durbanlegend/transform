package main

import (
	"flag" // command line option parser
	// "fmt"
	"github.com/durbanlegend/transform/ops"
)

var Usage = flag.Bool("u", false, "show usage and exit") // echo -u flag, of type *bool

const (
	Space   = " "
	Newline = "\n"
)

func main() {
	flag.Parse() // Scans the arg list and sets up flag

	if *Usage { // -u is parsed, flag becomes true
		flag.PrintDefaults()
		return
	}

	ops.Transform()
}
