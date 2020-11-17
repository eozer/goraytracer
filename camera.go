package goraytracer

import "math"

// DegreesToRadians converts degrees to radians
func DegreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180.0
}

// NewCamera creates a new camera with vfov (vertical field of view in degrees) and aspectRatio.
func NewCamera(
	lookfrom Point3,
	lookat Point3,
	vup Vec3,
	vfov float64,
	aspectRatio float64) Camera {
	var (
		// Camera constants
		theta          = DegreesToRadians(vfov)
		h              = math.Tan(theta / 2.0)
		viewportHeight = 2.0 * h
		viewportWidth  = aspectRatio * viewportHeight
		w              = UnitVector(Subtract(lookfrom, lookat))
		u              = UnitVector(Cross(vup, w))
		v              = Cross(w, u)
	)

	origin := lookfrom
	horizontal := Mult(viewportWidth, u)
	vertical := Mult(viewportHeight, v)
	lowerLeftCorner := origin.Clone()
	lowerLeftCorner.Sub(Div(2.0, horizontal))
	lowerLeftCorner.Sub(Div(2.0, vertical))
	lowerLeftCorner.Sub(w)
	return Camera{origin, lowerLeftCorner, horizontal, vertical}
}

type camera interface {
	GetRay(float64, float64) Ray
}

type Camera struct {
	origin          Point3
	lowerLeftCorner Point3
	horizontal      Vec3
	vertical        Vec3
}

func (c *Camera) GetRay(s, t float64) Ray {
	rayDirection := Subtract(c.lowerLeftCorner, c.origin)
	rayDirection.Add(Mult(s, c.horizontal))
	rayDirection.Add(Mult(t, c.vertical))
	ray := NewRay(c.origin, rayDirection)
	return ray
}
