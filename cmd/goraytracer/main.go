package main

import (
	"fmt"
	// TODO: Change GOPATH in .devcontainer
	"goraytracer"
	"os"
)

const (
	imageHeight   = 250
	imageWidth    = 250
	maxColorValue = 255
)

func main() {
	// See PPM specification: http://netpbm.sourceforge.net/doc/ppm.html
	fmt.Printf("P3\n%d %d\n%d\n", imageWidth, imageHeight, maxColorValue)
	for j := 0; j < imageHeight; j++ {
		fmt.Fprintf(os.Stderr, "\nScanlines remaining: %d", j)
		for i := 0; i < imageWidth; i++ {
			c := goraytracer.NewColor(float64(j)/(imageWidth-1), float64(i)/(imageHeight-1), 0.25)
			goraytracer.WriteColor(os.Stdout, c)
		}
	}
	fmt.Fprintf(os.Stderr, "\nDone\n")
}
