package parser

import (
	"fmt"
	"testing"
)

func TestMarshal(*testing.T) {
	Marshal()
	jsontype := Unmarshal("/tmp/config.json")
	fmt.Printf("Results: %v\n", jsontype)
}
