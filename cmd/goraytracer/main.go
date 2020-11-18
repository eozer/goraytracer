package main

import (
	"fmt"
	grt "github.com/eozer/goraytracer"
	"math/rand"
	"os"
)

const (
	// Image constants
	imageWidth      = 1200
	aspectRatio     = 16.0 / 9.0
	imageHeight     = int(imageWidth / aspectRatio)
	maxColorValue   = 255
	samplesPerPixel = 500
	maxDepth        = 50
)

func main() {
	// Set the camera
	lookfrom := grt.NewPoint3(9.0, 2.0, 3.0)
	lookat := grt.NewPoint3(0.0, 0.0, 0.0)
	vup := grt.NewVec3(0.0, 1.0, 0.0)
	vfov := 25.0
	aperture := 0.1
	distToFocus := 10.0
	camera := grt.NewCamera(lookfrom, lookat, vup, vfov, aspectRatio, aperture, distToFocus)
	world := grt.MakeRandomWorld()
	// See PPM specification: http://netpbm.sourceforge.net/doc/ppm.html
	fmt.Printf("P3\n%d %d\n%d\n", imageWidth, imageHeight, maxColorValue)
	for j := imageHeight - 1; j >= 0; j-- {
		fmt.Fprintf(os.Stderr, "\nScanlines remaining: %d", j)
		for i := 0; i < imageWidth; i++ {
			// Accumulate the pixel color
			pixelColor := grt.Color{}
			for s := 0; s < samplesPerPixel; s++ {
				u := (float64(i) + rand.Float64()) / (imageWidth - 1)
				v := (float64(j) + rand.Float64()) / float64(imageHeight-1)
				ray := camera.GetRay(u, v)
				color := grt.RayColor(&ray, world, maxDepth)
				pixelColor.Add(color)
			}
			grt.WriteColor(os.Stdout, pixelColor, samplesPerPixel)
		}
	}
	fmt.Fprintf(os.Stderr, "\nDone\n")
}
