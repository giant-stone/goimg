package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/giant-stone/goimg"
)

func main() {
	exampleResizePNG()
	exampleResizeJPG()
}

func exampleResizeJPG() {
	// fetch sample image form https://en.wikipedia.org/wiki/Lenna
	fn := "Lena.jpg"
	dat, err := ioutil.ReadFile(fn)
	exitOnErr(err)

	img, format, err := image.Decode(bytes.NewReader(dat))
	exitOnErr(err)

	widthOrigin, heightOrigin := img.Bounds().Max.X, img.Bounds().Max.Y
	width, height := uint(widthOrigin/2), uint(heightOrigin/2)
	imgResized := goimg.Resize(width, height, img, goimg.Lanczos3)

	output, err := os.Create("small.jpg")
	exitOnErr(err)
	defer output.Close()

	err = jpeg.Encode(output, imgResized, nil)
	exitOnErr(err)

	fmt.Printf("resize format=%s (%d,%d)=>(%d,%d) \n", format, widthOrigin, heightOrigin, width, height)
}

func exampleResizePNG() {
	// fetch sample image form https://en.wikipedia.org/wiki/Lenna
	fn := "Lena.png"
	dat, err := ioutil.ReadFile(fn)
	exitOnErr(err)

	img, format, err := image.Decode(bytes.NewReader(dat))
	exitOnErr(err)

	widthOrigin, heightOrigin := img.Bounds().Max.X, img.Bounds().Max.Y
	width, height := uint(widthOrigin/2), uint(heightOrigin/2)
	imgResized := goimg.Resize(width, height, img, goimg.Lanczos3)

	output, err := os.Create("small.png")
	exitOnErr(err)
	defer output.Close()

	err = png.Encode(output, imgResized)
	exitOnErr(err)

	fmt.Printf("resize format=%s (%d,%d)=>(%d,%d) \n", format, widthOrigin, heightOrigin, width, height)
}

func exitOnErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
