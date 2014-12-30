package parser

import (
	"fmt"
	"testing"
)

func TestTransType(*testing.T) {
	var i TypeId
	for i = 0; i <= String; i++ {
		fmt.Printf("%d=%s: GoType %s\n", i, i, GoTypes[i])
	}
}
