// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
)

func main() {

	t := []string{"g", "h", "i"}

	var s, sep string
	for i := 0; i < len(t); i++ {
		s += sep + t[i]
		sep = " "
	}
	fmt.Println(s)
}
