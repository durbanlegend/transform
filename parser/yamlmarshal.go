package parser

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func MarshalYAML(transtype TransType) {
	//home := os.Getenv("HOME")
	var err error
	var b []byte
	b, err = yaml.Marshal(transtype)
	check(err)

	err = ioutil.WriteFile("/tmp/config.yaml", b, 0644)
	check(err)
	//type Trans struct {
	//	Page   int      `yaml:"page"`
	//	Fruits []string `yaml:"fruits"`
	//}
}
func UnmarshalYAML(fileName string) TransType {
	file, e := ioutil.ReadFile(fileName)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	//fmt.Printf("Unmarshalling file %s:\n", fileName)
	//fmt.Printf("%s\n", string(file))

	var transtype TransType
	yaml.Unmarshal(file, &transtype)
	return transtype
}
