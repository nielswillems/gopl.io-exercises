// Exercise 1.4 prints the count and text of lines which appear more than once in the input files.
// It also prints the files the line is duplicated in.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	foundIn := make(map[string]string)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts, foundIn)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, foundIn)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t %s\n", n, line, foundIn[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, foundIn map[string]string) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		counts[input.Text()]++
		foundIn[input.Text()] += " " + f.Name()
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
