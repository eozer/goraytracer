package main

import (
	"fmt"
	// TODO: Change GOPATH in .devcontainer
	"goraytracer"
	rand "math/rand"
	"os"
)

const (
	// Image constants
	imageWidth      = 400
	aspectRatio     = 16.0 / 9.0
	imageHeight     = int(imageWidth / aspectRatio)
	maxColorValue   = 255
	samplesPerPixel = 100
	maxDepth        = 50
)

func main() {
	// Set the world
	world := goraytracer.World{}
	sphere := goraytracer.NewSphere(goraytracer.NewPoint3(0, 0, -1.0), 0.5)
	world.Add(&sphere)
	ground := goraytracer.NewSphere(goraytracer.NewPoint3(0, -100.5, -1.0), 100.0)
	world.Add(&ground)
	// Set the camera
	camera := goraytracer.NewCamera(aspectRatio)
	// See PPM specification: http://netpbm.sourceforge.net/doc/ppm.html
	fmt.Printf("P3\n%d %d\n%d\n", imageWidth, imageHeight, maxColorValue)
	for j := imageHeight - 1; j >= 0; j-- {
		fmt.Fprintf(os.Stderr, "\nScanlines remaining: %d", j)
		for i := 0; i < imageWidth; i++ {
			// Accumulate the pixel color
			pixelColor := goraytracer.Color{}
			for s := 0; s < samplesPerPixel; s++ {
				u := (float64(i) + rand.Float64()) / (imageWidth - 1)
				v := (float64(j) + rand.Float64()) / float64(imageHeight-1)
				ray := camera.GetRay(u, v)
				color := goraytracer.RayColor(&ray, &world, maxDepth)
				pixelColor.Add(color)
			}
			goraytracer.WriteColor(os.Stdout, pixelColor, samplesPerPixel)
		}
	}
	fmt.Fprintf(os.Stderr, "\nDone\n")
}
