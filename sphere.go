package goraytracer

import "math"

type sphere interface {
	Hittable
	GetCenter() *Point3
	GetRadius() float64
}

type Sphere struct {
	center Point3
	radius float64
	mat    Material
}

// NewSphere returns a new sphere at center with radius
func NewSphere(center Point3, radius float64, mat Material) Sphere {
	return Sphere{center, radius, mat}
}

func (s *Sphere) Hit(ray *Ray, tMin float64, tMax float64, rec *HitRecord) bool {
	oc := Subtract(*ray.GetOrigin(), s.center)
	a := ray.GetDirection().SqrLen()
	halfb := Dot(oc, *ray.GetDirection())
	c := oc.SqrLen() - s.radius*s.radius
	//
	discriminant := halfb*halfb - a*c
	if discriminant < 0.0 {
		return false
	}
	sqrtd := math.Sqrt(discriminant)
	// Find the nearest root that lies in the acceptable range.
	root := (-halfb - sqrtd) / a
	if root < tMin || tMax < root {
		root = (-halfb + sqrtd) / a
		if root < tMin || tMax < root {
			return false
		}
	}
	// Record
	rec.T = root
	rec.P = ray.At(rec.T)
	// NOTE: normal vector is set at geometry time, i.e., we set the normals
	// according to the ray that is passed.
	outwardNormal := Subtract(rec.P, s.center)
	outwardNormal.Div(s.radius)
	rec.SetFaceNormal(ray, &outwardNormal)
	rec.Mat = s.mat
	return true
}

// GetCenter returns center of Sphere
func (s *Sphere) GetCenter() *Point3 {
	return &s.center
}

// GetRadius returns radius of Sphere
func (s *Sphere) GetRadius() float64 {
	return s.radius
}
