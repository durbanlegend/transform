package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Transformation struct {
	ID      int
	Name    string
	Infile  string
	Outfile string
	Parms   []string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func Marshal() {
	home := os.Getenv("HOME")

	group := Transformation{
		ID:      1,
		Name:    "Transform",
		Infile:  home + "/dat",
		Outfile: "/tmp/out.dat",
		Parms:   []string{"Crimson", "Red", "Ruby", "Maroon", "Scarlet"},
	}
	b, err := json.MarshalIndent(group, "", "  ")
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
func Unmarshal() {
	file, e := ioutil.ReadFile("/tmp/config.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(file))

	var jsontype Transformation
	json.Unmarshal(file, &jsontype)
	fmt.Printf("Results: %v\n", jsontype)

}
