// Dup3 prints the count and text of lines that
// appear more than once in the named input files.
//
//
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)

	fileText := "file.txt"

	data, err := ioutil.ReadFile(fileText)

	if err != nil {
		fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
	}

	for _, line := range strings.Split(string(data), "\n") {
		counts[line]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
