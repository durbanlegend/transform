package main

import (
	"flag" // command line option parser
	//"fmt"
	"github.com/durbanlegend/transform/ops"
	"github.com/durbanlegend/transform/parser"
)

var Usage = flag.Bool("u", false, "Show usage and exit") // echo -u flag, of type *bool

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

	//ops.Transform()
	//parser.Parse()
	//fmt.Printf("Arguments: %v\n", flag.NArg())
	if flag.NArg() == 0 {
		flag.PrintDefaults()
		return
	}
	jsontype := parser.Unmarshal(flag.Arg(0))
	switch jsontype.Name {
	case "CacheFile":
		ops.CacheFile(jsontype)
	case "ReadFile":
		ops.ReadFile(jsontype)
	case "Transform":
		ops.Transform(jsontype)
	}
}
