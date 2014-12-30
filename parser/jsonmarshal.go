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
func Marshal() {
	home := os.Getenv("HOME")

	group := TransType{
		ID:      1,
		Name:    "Transform",
		Infile:  home + "/dat",
		Outfile: "/tmp/dat",
		//Parms:   []string{"Crimson", "Red", "Ruby", "Maroon", "Scarlet"},
		Inrec: InrecType{
			Name: "InRecname",
			Infields: []InfieldsType{
				{
					Name: "itemNum",
					Type: Integer,
				},
				{
					Name: "siteNum",
					Type: Integer,
				},
				{
					Name: "C",
					Type: Decimal,
				},

				{
					Name: "D",
					Type: Integer,
				},

				{
					Name: "E",
					Type: Decimal,
				},

				{
					Name: "F",
					Type: String,
				},

				{
					Name: "G",
					Type: String,
				},

				{
					Name: "H",
					Type: String,
				},

				{
					Name: "I",
					Type: String,
				},

				{
					Name: "J",
					Type: String,
				},

				{
					Name: "K",
					Type: String,
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
	//fmt.Printf("Unmarshalling file %s:\n", fileName)
	//fmt.Printf("%s\n", string(file))

	var jsontype TransType
	json.Unmarshal(file, &jsontype)
	return jsontype
}
