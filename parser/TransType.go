package parser

import (
	"math/big"
	"reflect"
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
	Type TypeId
}

type DataType struct {
	Id     TypeId
	GoType reflect.Type
}

//type tmpDataType struct {
//	typeName string
//}

type TypeId uint

const (
	Invalid TypeId = iota
	Integer
	Decimal
	String
)

type Enum interface {
	Name() string
	Ordinal() int
	ValueOf() *[]string
}

func (id TypeId) name() string {
	return TypeNames[id]
}
func (id TypeId) ordinal() int {
	return int(id)
}
func (id TypeId) String() string {
	return TypeNames[id]
}
func (id TypeId) values() *[]string {
	return &TypeNames
}

var GoTypes = []reflect.Type{
	reflect.TypeOf(nil),
	reflect.TypeOf(1),
	reflect.TypeOf(*big.NewRat(0, 1)),
	reflect.TypeOf(``),
}

var TypeNames = []string{
	"Invalid",
	"Integer",
	"Decimal",
	"String",
}

var typeNameMap = make(map[string]TypeId)

func GetTypeId(name string) TypeId {
	return typeNameMap[name]
}

func init() {
	for i, n := range TypeNames {
		typeNameMap[n] = TypeId(i)
	}
}

func (id TypeId) MarshalText() ([]byte, error) {
	return []byte(id.String()), nil
}

func (id *TypeId) UnmarshalText(text []byte) error {
	*id = GetTypeId(string(text))
	//fmt.Printf("string(text)=%v; id=%v\n", string(text), id)
	return nil
}
