package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

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

	start := time.Now()
	i := 0
	for scanner.Scan() {
		i++
		// fmt.Println(scanner.Text()) // Println will add back the final '\n'
		// fmt.Printf("%q\n", strings.Split(scanner.Text(), ","))
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading input:", err)
		}
		_, err := w.WriteString(strings.Join(strings.Split(scanner.Text(), ","), ";") + "\n")
		if err != nil {
			fmt.Fprintln(os.Stderr, "writing output:", err)
		}
	}
	end := time.Now()
	delta := end.Sub(start)
	//fmt.Println("Full transform took”, delta)
	fmt.Println("Elapsed:", delta)
	fmt.Println("Value of i is now:", i)
}
