package ops

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/durbanlegend/transform/parser"
)

type T struct {
	A int
	B int
	C big.Rat
	D int
	E float64
	F string
	G string
	H string
	I string
	J string
	K string
}

// Example of how to use Go's reflection
// Print the attributes of a Data Model
func attributes(m interface{}) map[string]reflect.Type {
	typ := reflect.TypeOf(m)
	// if a pointer to a struct is passed, get the type of the dereferenced object
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	// create an attribute data structure as a map of types keyed by a string.
	attrs := make(map[string]reflect.Type)
	// Only structs are supported so return an empty result if the passed object
	// isn't a struct
	if typ.Kind() != reflect.Struct {
		fmt.Printf("%v type can't have attributes inspected\n", typ.Kind())
		return attrs
	}

	// loop through the struct's fields and set the map
	for i := 0; i < typ.NumField(); i++ {
		p := typ.Field(i)
		if !p.Anonymous {
			attrs[p.Name] = p.Type
		}
	}

	return attrs
}

func ReadFile(trans parser.TransType) {
	for name, mtype := range attributes(&parser.TransType{}) {
		//for name, mtype := range attributes(&T{}) {
		fmt.Printf("Name: %s, Type %s\n", name, mtype.Name())
	}

	f, err := os.Open(trans.Infile)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	defer f.Close()

	// open output file
	fo, err := os.Create(trans.Outfile)
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	// make a write buffer
	w := bufio.NewWriter(fo)
	defer w.Flush()

	t := T{}
	start := time.Now()
	i := 0
	for scanner.Scan() {
		i++
		if scanner.Err() != nil {
			fmt.Println(scanner.Err())
			break
		}
		fmt.Println(scanner.Text()) // Println will add back the final '\n'
		fmt.Printf("%q\n", strings.Split(scanner.Text(), ","))
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading input:", err)
		}
		arr := strings.Split(scanner.Text(), ",")
		j := 0
		t.A, err = strconv.Atoi(arr[j])
		j++
		t.B, err = strconv.Atoi(arr[j])
		j++
		//t.C, err = strconv.ParseFloat(arr[j], 64)
		_, err = fmt.Sscan(arr[j], &t.C)
		if err != nil {
			fmt.Fprintln(os.Stderr, "reading input:", err)
		} else {
			fmt.Printf("ReadFile: C=%v\n", t.C.FloatString(3))
		}

		j++
		t.D, err = strconv.Atoi(arr[j])
		j++
		t.E, err = strconv.ParseFloat(arr[j], 64)
		j++
		t.F = arr[j]
		j++
		t.G = arr[j]
		j++
		t.H = arr[j]
		j++
		t.I = arr[j]
		j++
		t.J = arr[j]
		j++
		t.K = arr[j]
		j++
		// _` := w.WriteString(strings.Join(strings.Split(scanner.Text(), ","), ";") + "\n")
		_, err := w.WriteString(fmt.Sprintf("ReadFile: structure T: %v\n", t))
		if err != nil {
			fmt.Fprintln(os.Stderr, "writing output:", err)
		}
		//_, err := w.WriteString("\n")
		//if err != nil {
		//	fmt.Fprintln(os.Stderr, "writing output:", err)
		//}
	}
	end := time.Now()
	delta := end.Sub(start)
	//fmt.Println("Full transform took”, delta)
	fmt.Println("Elapsed:", delta)
	fmt.Println("Rows processed:", i)
}
