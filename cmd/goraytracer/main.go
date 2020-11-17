package main

import (
	"fmt"
	// TODO: Change GOPATH in .devcontainer
	grt "goraytracer"
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
	world := grt.World{}

	// Draw the world
	groundMat := grt.NewLambertian(grt.NewColor(0.8, 0.8, 0.0))
	groundSphere := grt.NewSphere(grt.NewPoint3(0, -100.5, -1.0), 100.0, &groundMat)
	world.Add(&groundSphere)

	// Put a lambertian material to the center of the camera
	centerMat := grt.NewLambertian(grt.NewColor(0.7, 0.3, 0.3))
	centerSphere := grt.NewSphere(grt.NewPoint3(0, 0, -1.0), 0.5, &centerMat)
	world.Add(&centerSphere)

	// Put a reflective metal material to the left
	// leftMat := grt.NewMetal(grt.NewColor(0.8, 0.8, 0.8), 0.3)
	leftMat := grt.NewDielectric(1.5)
	leftSphere := grt.NewSphere(grt.NewPoint3(-1.0, 0, -1.0), -0.5, &leftMat)
	world.Add(&leftSphere)

	// Put second reflective metal material to the right
	// rightMat := grt.NewDielectric(1.5)
	rightMat := grt.NewMetal(grt.NewColor(0.8, 0.6, 0.2), 0.0)
	rightSphere := grt.NewSphere(grt.NewPoint3(1.0, 0, -1.0), 0.5, &rightMat)
	world.Add(&rightSphere)

	// Set the camera
	lookfrom := grt.NewPoint3(-2.0, 2.0, 1.0)
	lookat := grt.NewPoint3(0.0, 0.0, -1.0)
	vup := grt.NewVec3(0.0, 1.0, 0.0)
	dist := grt.Subtract(lookfrom, lookat)
	vfov := 20.0
	aperture := 2.0
	distToFocus := dist.Len()
	camera := grt.NewCamera(lookfrom, lookat, vup, vfov, aspectRatio, aperture, distToFocus)
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
				color := grt.RayColor(&ray, &world, maxDepth)
				pixelColor.Add(color)
			}
			grt.WriteColor(os.Stdout, pixelColor, samplesPerPixel)
		}
	}
	fmt.Fprintf(os.Stderr, "\nDone\n")
}
