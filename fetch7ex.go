// Fetch выводит ответ на запрос по заданному URL
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	u := []string{"http://golang.org", "http://megaweb.kz"}
	for _, url := range u[0:] {

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		}

		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			fmt.Fprintf(os.Stderr, "fetch: чтение %s: %v\n", url, err)
			os.Exit(1)
		}

		resp.Body.Close()
	}
}
