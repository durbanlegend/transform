package parser

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func Parse() {
	f, _ := os.Open("/home/donf/transform.json")
	dec := json.NewDecoder(f)
	defer f.Close()
	var v map[string]interface{}
	if err := dec.Decode(&v); err != nil {
		log.Println("Decoding error")
		log.Println(err)
		return
	}
	for k, x := range v {
		fmt.Println(k, ":", x)
	}
}
