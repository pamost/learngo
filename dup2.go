// Prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
// Выводит текст каждой строки, которая появляется во входных данных более одного раза.
// Программа читает стандартный ввод или список именованных файлов.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	counts := make(map[string]int)

	file := "file.txt"

	f, err := os.Open(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
	}

	countLines(f, counts)
	f.Close()

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if input.Text() == "" {
			break
		}
		counts[input.Text()]++
	}
}
