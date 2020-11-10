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

