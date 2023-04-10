# About

Pure golang image resizing.

This package is a fork of https://github.com/nfnt/resize  , required Go >=1.12+.

## Usage

Fetch package `go get github.com/giant-store/goimg`

Example: resize a JPEG image

```go
import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"

	"github.com/giant-stone/goimg"
)

func main() {
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

func exitOnErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
```

The goimg package provides 2 functions:

* `goimg.Resize` creates a scaled image with new dimensions (`width`, `height`) using the interpolation function `i`.
  If either `width` or `height` is set to 0, it will be set to an aspect ratio preserving value.
* `goimg.Thumbnail` down-scales an image preserving its aspect ratio to the maximum dimensions (`maxWidth`, `maxHeight`).
  It will return the original image if original sizes are smaller than the provided dimensions.

```go
goimg.Resize(width, height uint, img image.Image, i goimg.InterpolationFunction) image.Image
goimg.Thumbnail(maxWidth, maxHeight uint, img image.Image, i goimg.InterpolationFunction) image.Image
```

The provided interpolation functions are (from fast to slow execution time)

- `NearestNeighbor`: [Nearest-neighbor interpolation](http://en.wikipedia.org/wiki/Nearest-neighbor_interpolation)
- `Bilinear`: [Bilinear interpolation](http://en.wikipedia.org/wiki/Bilinear_interpolation)
- `Bicubic`: [Bicubic interpolation](http://en.wikipedia.org/wiki/Bicubic_interpolation)
- `MitchellNetravali`: [Mitchell-Netravali interpolation](http://dl.acm.org/citation.cfm?id=378514)
- `Lanczos2`: [Lanczos resampling](http://en.wikipedia.org/wiki/Lanczos_resampling) with a=2
- `Lanczos3`: [Lanczos resampling](http://en.wikipedia.org/wiki/Lanczos_resampling) with a=3

Which of these methods gives the best results depends on your use case.
