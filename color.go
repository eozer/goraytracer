package goraytracer

import (
	"fmt"
	"io"
)

type Color = Vec3

func NewColor(red, green, blue float64) Color {
	return NewVec3(red, green, blue)
}

// WriteColor writes pixelColor in [0, 255] scale
func WriteColor(writer io.Writer, pixelColor Color) {
	ir := int(255.999 * pixelColor.GetX())
	ig := int(255.999 * pixelColor.GetY())
	ib := int(255.999 * pixelColor.GetZ())
	fmt.Fprintf(writer, "%d %d %d\n", ir, ig, ib)
}

// RayColor creates a gradient color from c1 to c2 along Y axesray's unit vector.
func RayColor(ray *Ray) Color {
	if HitSphere(NewPoint3(0.0, 0.0, -1.0), 0.5, ray) {
		return NewColor(1, 0, 0)
	}
	unitdir := UnitVector(ray.GetDirection())
	t := 0.5 * (unitdir.GetY() + 1.0)
	// fmt.Fprintf(os.Stderr, "\n%f \n", unitdir.GetY())
	c1 := NewColor(1.0, 1.0, 1.0)
	c2 := NewColor(0.5, 0.7, 1.0)
	c1.Mult(1.0 - t)
	c2.Mult(t)
	return Add(c1, c2)
}
