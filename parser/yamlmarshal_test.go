package parser

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMarshalYAML(*testing.T) {
	// Initial bootstrap from JSON:
	//transtype := UnmarshalJSON("/home/donf/person.json")
	transtype := UnmarshalYAML("/home/donf/person.yaml")
	MarshalYAML(transtype)
	fmt.Printf("Unmarshalled %s: %v\n", reflect.TypeOf(transtype), transtype)
}
