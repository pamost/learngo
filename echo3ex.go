package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	start := time.Now()
	for i := 1000000000; i >= 0; i-- {
	}
	fmt.Printf("%.10fs \n", time.Since(start).Seconds())

	start1 := time.Now()
	echo1()
	fmt.Printf("%.10fs \n", time.Since(start1).Seconds())

	start2 := time.Now()
	echo2()
	fmt.Printf("%.10fs \n", time.Since(start2).Seconds())

	start3 := time.Now()
	echo3()
	fmt.Printf("%.10fs \n", time.Since(start3).Seconds()) // or: time.Now() - startTime
}

func echo1() {
	m := []string{"g", "h", "i"}

	var s, sep string
	for i := 0; i < len(m); i++ {
		s += sep + m[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	d := []string{"g", "h", "i"}

	s, sep := "", ""
	for _, value := range d[0:] {
		s += sep + value
		sep = " "
	}
	fmt.Println(s)
}

func echo3() {
	a := []string{"g", "h", "i"}
	fmt.Println(strings.Join(a[0:], " "))
}
