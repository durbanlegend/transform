package parser

import (
	"fmt"
	"testing"
)

func TestDNA(*testing.T) {
	x := New(A, C, G, T)
	fmt.Println(x)
}
