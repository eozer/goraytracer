package goraytracer

// Hittable represents that a material is ray-hittable.
type Hittable interface {
	// Hit(, double t_min, double t_max, hit_record& rec) bool
	Hit(*Ray, float64, float64, *HitRecord) bool
}

type hitrecord interface {
	SetFaceNormal(*Ray, *Vec3)
}

// HitRecord is a type used for storing the results of Hittable.Hit.
type HitRecord struct {
	P         Point3
	Normal    Vec3
	T         float64
	FrontFace bool
}

func NewHitRecord(p Point3, n Vec3, t float64, ff bool) HitRecord {
	return HitRecord{p, n, t, ff}
}

func (hr *HitRecord) SetFaceNormal(ray *Ray, outwardNormal *Vec3) {
	hr.FrontFace = Dot(*ray.GetDirection(), *outwardNormal) < 0.0
	c := outwardNormal.Clone()
	if hr.FrontFace {
		hr.Normal = c
	} else {
		hr.Normal = *c.Neg()
	}
}
