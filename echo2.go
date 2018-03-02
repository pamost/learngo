// Echo2 prints its command-line arguments through range
package main

import (
	"fmt"
)

func main() {

	d := []string{"g", "h", "i"}

	s, sep := "", ""
	for _, value := range d[0:] {
		s += sep + value
		sep = " "
	}
	fmt.Println(s)
}
