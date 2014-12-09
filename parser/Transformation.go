package parser

import (
//"reflect"
)

type TransType struct {
	ID      int
	Name    string
	Infile  string
	Outfile string
	// Parms   []string
	Inrec InrecType
}

type InrecType struct {
	Name     string
	Infields []InfieldsType
}

type InfieldsType struct {
	Name string
	Type string
}
