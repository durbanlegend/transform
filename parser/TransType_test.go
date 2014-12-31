package parser

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestTransType(*testing.T) {
	var i TypeId
	for i = 0; i <= String; i++ {
		fmt.Printf("%d=%s: GoType=%s; TypeName=%s; GetType(TypeName)=%s\n", i, i, GoTypes[i], TypeNames[i], GetTypeId(TypeNames[i]))

		dt := DataType{
			Id:     i,
			GoType: GoTypes[i],
		}
		//txt, err := dt.MarshalJSON()
		//txt, err := dt.MarshalText()
		txt, err := json.Marshal(dt)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%d=%s\n", i, txt)
	}
	fmt.Printf("Type of GoTypes=%s\n", reflect.TypeOf(GoTypes))

}
