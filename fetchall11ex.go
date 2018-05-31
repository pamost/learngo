// Fetchall выполняет параллельную выборку URL и сообщает
// о затраченном времени и размере ответа для каждого из них.
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {

	u := []string{}
	fileText := "site11ex.txt"

	fileHandle, _ := os.Open(fileText)
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	for fileScanner.Scan() {
		u = append(u, fileScanner.Text())
	}

	start := time.Now()
	ch := make(chan string)
	for _, url := range u[0:] {
		go fetch11ex("http://"+url, ch)
	}

	// Время старта новой проверки
	copyFile11ex(start.Format("2006-01-02 15:04:05"))

	for range u[0:] {
		copyFile11ex(<-ch) // Получение из канала ch
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

	copyFile11ex("\n")
}

func fetch11ex(url string, ch chan<- string) {
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

func copyFile11ex(fo string) {
	fmt.Println(fo)

	fileInput := "text11ex.txt"

	// Если файла не существует, создаем его
	if _, err := os.Stat(fileInput); os.IsNotExist(err) {
		o, err := os.OpenFile(fileInput, os.O_RDONLY|os.O_CREATE, 0666)
		if err != nil {
			panic(err)
		}
		o.Close()
	}

	// Производим запись в существующий файл
	f, err := os.OpenFile(fileInput, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	f.WriteString(fo + "\n")
	f.Close()
}
