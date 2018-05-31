// Fetchall выполняет параллельную выборку URL и сообщает
// о затраченном времени и размере ответа для каждого из них.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	u := []string{
		"http://golang.org",
		"http://megaweb.kz",
		"http://smartgadget.kz",
		"http://google.com",
	}

	start := time.Now()
	ch := make(chan string)
	for _, url := range u[0:] {
		go fetch10ex(url, ch)
	}
	// Время старта новой проверки
	copyFile(start.Format("2006-01-02 15:04:05"))

	for range u[0:] {
		copyFile(<-ch) // Получение из канала ch
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

	copyFile("\n")
}

func fetch10ex(url string, ch chan<- string) {
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

func copyFile(fo string) {
	fmt.Println(fo)

	// Если файла не существует, создаем его
	if _, err := os.Stat("text10ex.txt"); os.IsNotExist(err) {
		o, err := os.OpenFile("text10ex.txt", os.O_RDONLY|os.O_CREATE, 0666)
		if err != nil {
			panic(err)
		}
		o.Close()
	}

	// Производим запись в существующий файл
	f, err := os.OpenFile("text10ex.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	f.WriteString(fo + "\n")
	f.Close()
}
