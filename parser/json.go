package parser

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
)

type Trans struct {
	Name    string
	Op      string
	Age     int
	Parents []interface{}
}

//type Trans struct {
//	Page   int      `json:"page"`
//	Fruits []string `json:"fruits"`
//}

func Parse() {
	home := os.Getenv("HOME")
	//f, _ := os.Open("/home/donf/transform.json")
	f, _ := os.Open(home + "/transform.json")
	dec := json.NewDecoder(f)
	defer f.Close()
	var v map[string]interface{}
	if err := dec.Decode(&v); err != nil {
		log.Println("Decoding error")
		log.Println(err)
		return
	}

	// To store the keys in slice in sorted order
	var keys []string
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println(k, ":", v[k])
	}
}
