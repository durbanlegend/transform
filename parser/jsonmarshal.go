package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func MarshalJSON(jsontype TransType, indent bool) {
	//home := os.Getenv("HOME")
	var err error
	var b []byte
	if indent {
		b, err = json.MarshalIndent(jsontype, "", "  ")
	} else {
		b, err = json.Marshal(jsontype)
	}
	check(err)

	err = ioutil.WriteFile("/tmp/config.json", b, 0644)
	check(err)
	//type Trans struct {
	//	Page   int      `json:"page"`
	//	Fruits []string `json:"fruits"`
	//}
}
func UnmarshalJSON(fileName string) TransType {
	file, e := ioutil.ReadFile(fileName)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	//fmt.Printf("Unmarshalling file %s:\n", fileName)
	//fmt.Printf("%s\n", string(file))

	var jsontype TransType
	json.Unmarshal(file, &jsontype)
	return jsontype
}
