package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func Marshal() {
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
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

	var jsontype ColorGroup
	json.Unmarshal(file, &jsontype)
	fmt.Printf("Results: %v\n", jsontype)

}
