package goraytracer

func NewCamera(aspectRatio float64) Camera {
	var (
		// Camera constants
		viewportHeight = 2.0
		viewportWidth  = aspectRatio * viewportHeight
		focalLength    = 1.0
	)
	o := Point3{}
	horizontal := NewVec3(viewportWidth, 0.0, 0.0)
	vertical := NewVec3(0.0, viewportHeight, 0.0)
	lowerLeftCorner := Point3{}
	lowerLeftCorner.Sub(Div(2.0, horizontal))
	lowerLeftCorner.Sub(Div(2.0, vertical))
	lowerLeftCorner.Sub(NewVec3(0.0, 0.0, focalLength))
	return Camera{o, lowerLeftCorner, horizontal, vertical}
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

func (c *Camera) GetRay(u, v float64) Ray {
	rayDirection := Subtract(c.lowerLeftCorner, c.origin)
	rayDirection.Add(Mult(u, c.horizontal))
	rayDirection.Add(Mult(v, c.vertical))
	ray := NewRay(c.origin, rayDirection)
	return ray
}
