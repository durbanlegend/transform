package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type T struct {
	A int
	B int
	C float64
	D int
	E float64
	F string
	G string
	H string
	I string
	J string
	K string
}

func main() {
	f, _ := os.Open("/tmp/dat")
	scanner := bufio.NewScanner(f)
	defer f.Close()
	for i := 0; i < 5; i++ {
		fmt.Println("Value of i is now:", i)
	}

	// open output file
	fo, err := os.Create("/tmp/output.txt")
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
		// fmt.Println(scanner.Text()) // Println will add back the final '\n'
		// fmt.Printf("%q\n", strings.Split(scanner.Text(), ","))
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading input:", err)
		}
		arr := strings.Split(scanner.Text(), "|")
		i := 0
		t.A, err = strconv.Atoi(arr[i])
		i++
		t.B, err = strconv.Atoi(arr[i])
		i++
		t.C, err = strconv.ParseFloat(arr[i], 64)
		i++
		t.D, err = strconv.Atoi(arr[i])
		i++
		t.E, err = strconv.ParseFloat(arr[i], 64)
		i++
		t.F = arr[i]
		i++
		t.G = arr[i]
		i++
		t.H = arr[i]
		i++
		t.I = arr[i]
		i++
		t.J = arr[i]
		i++
		t.K = arr[i]
		i++
		// _` := w.WriteString(strings.Join(strings.Split(scanner.Text(), ","), ";") + "\n")
		_, err := w.WriteString(fmt.Sprintf("%v\n", t))
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
	fmt.Println("Value of i is now:", i)
}
