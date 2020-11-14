package goraytracer

import "math"

func HitSphere(center Point3, radius float64, ray *Ray) float64 {
	oc := Subtract(ray.GetOrigin(), center)
	a := Dot(ray.GetDirection(), ray.GetDirection())
	b := 2.0 * Dot(oc, ray.GetDirection())
	c := Dot(oc, oc) - radius*radius
	discriminant := b*b - 4*a*c
	if discriminant < 0.0 {
		return -1.0
	}
	return (-b - math.Sqrt(discriminant)) / (2.0 * a)
}
