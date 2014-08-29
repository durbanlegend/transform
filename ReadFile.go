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
	start := time.Now()
	i := 0
	for scanner.Scan() {
		i++
		// fmt.Println(scanner.Text()) // Println will add back the final '\n'
		fmt.Printf("%q\n", strings.Split(scanner.Text(), ","))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	end := time.Now()
	delta := end.Sub(start)
	// fmt.Printf("Full read took this amount of time: %s\n”, delta)
	fmt.Println("Delta is", delta)
	fmt.Println("Value of i is now:", i)
}
