package main

import (
	"fmt"
	// TODO: Change GOPATH in .devcontainer
	"goraytracer"
	"os"
)

const (
	// Image constants
	imageWidth    = 400
	aspectRatio   = 16.0 / 9.0
	imageHeight   = int(imageWidth / aspectRatio)
	maxColorValue = 255
	// Camera constants
	viewportHeight = 2.0
	viewportWidth  = aspectRatio * viewportHeight
	focalLength    = 1.0
)

func main() {
	// Set the world
	world := goraytracer.World{}
	sphere := goraytracer.NewSphere(goraytracer.NewPoint3(0, 0, -1.0), 0.5)
	world.Add(&sphere)
	ground := goraytracer.NewSphere(goraytracer.NewPoint3(0, -100.5, -1.0), 100.0)
	world.Add(&ground)
	// Set the camera
	rayOrigin := goraytracer.Point3{}
	horizontal := goraytracer.NewVec3(viewportWidth, 0.0, 0.0)
	vertical := goraytracer.NewVec3(0.0, viewportHeight, 0.0)
	lowerLeftCorner := rayOrigin.Clone()
	lowerLeftCorner.Sub(goraytracer.Div(2.0, horizontal))
	lowerLeftCorner.Sub(goraytracer.Div(2.0, vertical))
	lowerLeftCorner.Sub(goraytracer.NewVec3(0.0, 0.0, focalLength))
	// See PPM specification: http://netpbm.sourceforge.net/doc/ppm.html
	fmt.Printf("P3\n%d %d\n%d\n", imageWidth, imageHeight, maxColorValue)
	for j := imageHeight - 1; j >= 0; j-- {
		fmt.Fprintf(os.Stderr, "\nScanlines remaining: %d", j)
		for i := 0; i < imageWidth; i++ {
			u := float64(i) / (imageWidth - 1)
			v := float64(j) / float64(imageHeight-1)
			rayDirection := goraytracer.Subtract(lowerLeftCorner, rayOrigin)
			rayDirection.Add(goraytracer.Mult(u, horizontal))
			rayDirection.Add(goraytracer.Mult(v, vertical))
			ray := goraytracer.NewRay(rayOrigin, rayDirection)
			pixelColor := goraytracer.RayColor(&ray, &world)
			goraytracer.WriteColor(os.Stdout, pixelColor)
		}
	}
	fmt.Fprintf(os.Stderr, "\nDone\n")
}
