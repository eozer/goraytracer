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
	aspectRatio float64,
	aperture float64,
	focusDistance float64) Camera {
	var (
		theta          = DegreesToRadians(vfov)
		h              = math.Tan(theta / 2.0)
		viewportHeight = 2.0 * h
		viewportWidth  = aspectRatio * viewportHeight
		w              = UnitVector(Subtract(lookfrom, lookat))
		u              = UnitVector(Cross(vup, w))
		v              = Cross(w, u)
		origin         = lookfrom
		horizontal     = Mult(focusDistance*viewportWidth, u)
		vertical       = Mult(focusDistance*viewportHeight, v)
		lensRadius     = aperture / 2.0
	)
	lowerLeftCorner := origin.Clone()
	lowerLeftCorner.Sub(Div(2.0, horizontal))
	lowerLeftCorner.Sub(Div(2.0, vertical))
	lowerLeftCorner.Sub(Mult(focusDistance, w))
	return Camera{origin, lowerLeftCorner, horizontal, vertical, u, v, w, lensRadius}
}

type camera interface {
	GetRay(float64, float64) Ray
	GetOrigin() Point3
}

type Camera struct {
	origin          Point3
	lowerLeftCorner Point3
	horizontal      Vec3
	vertical        Vec3
	u               Vec3
	v               Vec3
	w               Vec3
	lensRadius      float64
}

func (c *Camera) GetRay(s, t float64) Ray {
	rd := Mult(c.lensRadius, NewRandomVec3InUnitDisk())
	offset := Add(Mult(rd.GetX(), c.u), Mult(rd.GetY(), c.v))
	rayOrigin := Add(c.origin, offset)
	rayDirection := Subtract(c.lowerLeftCorner, c.origin)
	rayDirection.Sub(offset)
	rayDirection.Add(Mult(s, c.horizontal))
	rayDirection.Add(Mult(t, c.vertical))
	ray := NewRay(rayOrigin, rayDirection)
	return ray
}

func (c *Camera) GetOrigin() Point3 {
	return c.origin
}
