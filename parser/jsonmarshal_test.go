package parser

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMarshal(*testing.T) {
	Marshal()
	jsontype := Unmarshal("/tmp/config.json")
	fmt.Printf("Unmarshalled %s: %v\n", reflect.TypeOf(jsontype), jsontype)
}
