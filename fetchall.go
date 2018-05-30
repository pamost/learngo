// Fetchall выполняет параллельную выборку URL и сообщает
// о затраченном времени и размере ответа для каждого из них.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	u := []string{
		"http://golang.org",
		"http://megaweb.kz",
		"http://smartgadget.kz",
	}

	start := time.Now()
	ch := make(chan string)
	for _, url := range u[0:] {
		go fetch(url, ch)
	}

	for range u[0:] {
		fmt.Println(<-ch) // Получение из канала ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // Отправка в канал ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // Исключение утечки ресурсов
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
