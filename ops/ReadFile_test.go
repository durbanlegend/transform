package ops

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/durbanlegend/transform/parser"
)

func TestReadFile(*testing.T) {
	parser.Marshal()
	//home := os.Getenv("HOME")
	//jsontype := parser.Unmarshal(home + "/config.json")
	jsontype := parser.Unmarshal("/tmp/config.json")
	ReadFile(jsontype)
	fmt.Printf("Unmarshalled %s: %v\n", reflect.TypeOf(jsontype), jsontype)
}
