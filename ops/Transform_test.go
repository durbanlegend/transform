package ops

import (
	"github.com/durbanlegend/transform/parser"
	"testing"
)

func TestTransform(*testing.T) {
	parser.Marshal()
	jsontype := parser.Unmarshal("/tmp/config.json")
	Transform(jsontype)
	//fmt.Printf("Unmarshalled %s: %v\n", reflect.TypeOf(jsontype), jsontype)
}
