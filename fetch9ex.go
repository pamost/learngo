// Fetch выводит ответ на запрос по заданному URL
package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	u := []string{"golang.org", "http://megaweb.kz", "smartgadget.kz"}
	for _, url := range u[0:] {
		resp, err := http.Get(checkUrl9ex(url))
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		}
		fmt.Println(resp.Status)
	}
}

func checkUrl9ex(url string) string {
	prefixUrl := "http://"

	if strings.HasPrefix(url, prefixUrl) {
		return url
	}
	return prefixUrl + url
}
