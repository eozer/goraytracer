package goraytracer

type Material interface {
	Scatter(*Ray, *HitRecord) (bool, Color, Ray)
}

type ILambertian interface {
	Material
}

type Lambertian struct {
	albedo Color
}

func NewLambertian(color Color) Lambertian {
	return Lambertian{albedo: color}
}

func (l *Lambertian) Scatter(in *Ray, rec *HitRecord) (bool, Color, Ray) {
	// scatterDirection := Add(rec.Normal, NewRandomVec3())
	scatterDirection := Add(rec.Normal, NewRandomVec3InHemisphere(rec.Normal))
	// normalize degenerate scatterDirection vector direction.
	if scatterDirection.IsNearZero() {
		scatterDirection = rec.Normal
	}
	scatteredRay := NewRay(rec.P, scatterDirection)
	return true, l.albedo, scatteredRay
}

type IMetal interface {
	Material
}

type Metal struct {
	albedo Color
	fuzz   float64
}

func NewMetal(albedo Color, fuzz float64) Metal {
	f := 1.0
	if fuzz < 1.0 {
		f = fuzz
	}
	return Metal{albedo, f}
}

func (m *Metal) Scatter(in *Ray, rec *HitRecord) (bool, Color, Ray) {
	reflected := Reflect(in.GetDirection().Unit(), rec.Normal)
	scatteredRay := NewRay(rec.P, Add(reflected, Mult(m.fuzz, NewRandomVec3InUnitSphere())))
	isScattered := (Dot(*scatteredRay.GetDirection(), rec.Normal) > 0.0)
	return isScattered, m.albedo, scatteredRay
}
