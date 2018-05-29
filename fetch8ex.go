// Fetch выводит ответ на запрос по заданному URL
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	u := []string{"golang.org", "http://megaweb.kz"}
	for _, url := range u[0:] {

		resp, err := http.Get(checkUrl(url))
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		}

		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		resp.Body.Close()
	}
}

func checkUrl(url string) string {
	prefixUrl := "http://"

	if strings.HasPrefix(url, prefixUrl) {
		return url
	}
	return prefixUrl + url
}
