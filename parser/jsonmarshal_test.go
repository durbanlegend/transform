package parser

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMarshalJSON(*testing.T) {
	jsontype := UnmarshalJSON("/home/donf/person.json")
	MarshalJSON(jsontype, true)
	fmt.Printf("Unmarshalled %s: %v\n", reflect.TypeOf(jsontype), jsontype)
}
