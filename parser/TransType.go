package parser

import "reflect"

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
	Type TypeId
}

type DataType struct {
	Id     TypeId
	GoType reflect.Type
}

type TypeId int

const (
	Unknown TypeId = iota
	Integer
	Decimal
	String
)

// Enums should be able to printout as strings
// so we declare the next best thing, a slice of strings
// for eg. the string value will be used in the println
var GoTypes = [...]reflect.Kind{
	reflect.Invalid,
	reflect.Int,
	reflect.Float64, // TODO replace with big.Rat
	reflect.String,
}

func (t TypeId) String() string {
	var s string
	switch t {
	case Integer:
		s = "Integer"
	case Decimal:
		s = "Decimal"
	case String:
		s = "String"
	case Unknown:
		s = "Invalid"
	default:
		s = "¡Unknown or composed Type!"
	}
	return s
}
