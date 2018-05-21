// Echo2ex prints its command-line arguments through range
package main

import "fmt"

func main() {
	d := []string{"g", "h", "i"}

	for idx, value := range d[0:] {
		fmt.Println(idx, value)
	}
}
