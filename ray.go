package goraytracer

type Ray struct {
	origin    Point3
	direction Vec3
}

type IRay interface {
	GetOrigin() Point3
	GetDirection() Vec3
	At(float64) Point3
}

func NewRay(origin Point3, direction Vec3) Ray {
	return Ray{origin, direction}
}

func (r *Ray) GetOrigin() Point3 {
	return r.origin
}

func (r *Ray) GetDirection() Vec3 {
	return r.direction
}

func (r *Ray) At(t float64) Point3 {
	// origin + t * direction
	tdotd := Mult(t, r.direction)
	optdotd := Add(r.origin, tdotd)
	return optdotd
}
