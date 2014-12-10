package ops

import (
	"bufio"
	"fmt"
	"github.com/durbanlegend/transform/parser"
	"os"
	//"reflect"
	//"strconv"
	"strings"
	"time"
)

func Transform(trans parser.TransType) {
	//for name, mtype := range attributes(&parser.TransType{}) {
	//	fmt.Printf("Name: %s, Type %s\n", name, mtype.Name())
	//}

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

	// open output file
	fe, err := os.Create("/dev/stderr")
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fe.Close(); err != nil {
			panic(err)
		}
	}()

	ew := bufio.NewWriter(fe)
	defer ew.Flush()

	start := time.Now()
	i := 0
	for scanner.Scan() {
		i++
		if scanner.Err() != nil {
			fmt.Println(scanner.Err())
			break
		}
		//fmt.Println(scanner.Text()) // Println will add back the final '\n'
		//fmt.Printf("%q\n", strings.Split(scanner.Text(), ","))
		//if err := scanner.Err(); err != nil {
		//	fmt.Fprintln(os.Stderr, "reading input:", err)
		//}

		arr := strings.Split(scanner.Text(), ",")

		j := 0
		fieldArray := trans.Inrec.Infields
		fmt.Printf("trans.Infields: %s\n", fieldArray)
		for pos, infield := range fieldArray {
			fmt.Printf("Pos: %d, Name: %s, Type: %s, Value %s\n", pos, infield.Name, infield.Type, arr[j])
			j++
		}

		//_, err := w.WriteString(fmt.Sprintf("%v\n", t))
		_, err := w.WriteString(strings.Join(arr, ",") + "\n")
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
	fmt.Fprintln(ew, "Elapsed:", delta)
	fmt.Fprintln(ew, "Rows processed:", i)
}
