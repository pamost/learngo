package main

import (
	"github.com/vova616/screenshot"
	"fmt"
	"os"
	"image/png"
)

func main() {
	img, err := screenshot.CaptureScreen()
	if err != nil {
		fmt.Println(err)
	}

	out, err := os.Create("./screen.png")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Write out the data into the new PNG file
	err = png.Encode(out, img)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	out.Close()
	fmt.Println("Screen save to screen.png \n")
}
