package goraytracer

func HitSphere(center Point3, radius float64, ray *Ray) bool {
	oc := Subtract(ray.GetOrigin(), center)
	a := Dot(ray.GetDirection(), ray.GetDirection())
	b := 2.0 * Dot(oc, ray.GetDirection())
	c := Dot(oc, oc) - radius*radius
	discriminant := b*b - 4*a*c
	return (discriminant > 0)
}
