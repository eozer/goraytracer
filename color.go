package goraytracer

import (
	"fmt"
	"io"
	"math"
)

type Color = Vec3

func NewColor(red, green, blue float64) Color {
	return NewVec3(red, green, blue)
}

func ClampColor(val, min, max float64) float64 {
	if val < min {
		return min
	}
	if val > max {
		return max
	}
	return val
}

// WriteColor writes pixelColor in [0, 255] scale
func WriteColor(writer io.Writer, pixelColor Color, samplesPerPixel int) {
	// NOTE: We accumulate pixel colors * samplesPerPixel
	// Divide the color by the number of samples.
	scale := 1.0 / float64(samplesPerPixel)
	r := pixelColor.GetX() * scale
	g := pixelColor.GetY() * scale
	b := pixelColor.GetZ() * scale
	// Scale colors in [0.0,1.0] back to [0, 255]
	ir := int(256 * ClampColor(r, 0.0, 0.999))
	ig := int(256 * ClampColor(g, 0.0, 0.999))
	ib := int(256 * ClampColor(b, 0.0, 0.999))
	fmt.Fprintf(writer, "%d %d %d\n", ir, ig, ib)
}

// RayColor creates a gradient color from c1 to c2 along Y axesray's unit vector.
func RayColor(ray *Ray, world Hittable) Color {
	rec := HitRecord{}
	pInf := math.Inf(1)
	if world.Hit(ray, 0, pInf, &rec) {
		g := Add(rec.Normal, NewColor(1.0, 1.0, 1.0))
		g.Mult(0.5)
		return g
	}
	unitdir := ray.GetDirection().Unit()
	t := 0.5 * (unitdir.GetY() + 1.0)
	// fmt.Fprintf(os.Stderr, "\n%f \n", unitdir.GetY())
	c1 := NewColor(1.0, 1.0, 1.0)
	c2 := NewColor(0.5, 0.7, 1.0)
	c1.Mult(1.0 - t)
	c2.Mult(t)
	return Add(c1, c2)
}
