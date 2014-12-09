package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func Marshal() {
	home := os.Getenv("HOME")

	group := TransType{
		ID:      1,
		Name:    "Transform",
		Infile:  home + "/dat",
		Outfile: "/tmp/out.dat",
		//Parms:   []string{"Crimson", "Red", "Ruby", "Maroon", "Scarlet"},
		Inrec: InrecType{
			Name: "InRecname",
			Infields: []InfieldsType{
				{
					Name: "itemNum",
					Type: reflect.Int.String(),
				},
				{
					Name: "siteNum",
					Type: "int",
				},
				{
					Name: "C",
					Type: "float64",
				},

				{
					Name: "D",
					Type: "int",
				},

				{
					Name: "E",
					Type: "float64",
				},

				{
					Name: "F",
					Type: "string",
				},

				{
					Name: "G",
					Type: "string",
				},

				{
					Name: "H",
					Type: "string",
				},

				{
					Name: "I",
					Type: "string",
				},

				{
					Name: "J",
					Type: "string",
				},

				{
					Name: "K",
					Type: "string",
				},
			},
		},
	}
	b, err := json.MarshalIndent(group, "", "  ")
	//b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	//os.Stdout.Write(b)
	err = ioutil.WriteFile("/tmp/config.json", b, 0644)
	check(err)
	//type Trans struct {
	//	Page   int      `json:"page"`
	//	Fruits []string `json:"fruits"`
	//}
}
func Unmarshal(fileName string) TransType {
	file, e := ioutil.ReadFile(fileName)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	fmt.Printf("Marshalled file %s:\n", fileName)
	fmt.Printf("%s\n", string(file))

	var jsontype TransType
	json.Unmarshal(file, &jsontype)
	return jsontype
}
