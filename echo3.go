// Echo3 prints its command-line arguments through strings.Join
// a more productive operation
package main

import (
	"fmt"
	"strings"
)

//!+
func main() {
	a := []string{"g", "h", "i"}
	fmt.Println(strings.Join(a[0:], " "))
}
