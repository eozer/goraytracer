package goraytracer

import (
	"math"
	"math/rand"
)

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

type IDielectric interface {
	Material
	GetReflectance(cos, refIdx float64) float64
}

type Dielectric struct {
	refractionIndex float64
}

func NewDielectric(refractionIndex float64) Dielectric {
	return Dielectric{refractionIndex: refractionIndex}
}

func (d *Dielectric) Scatter(in *Ray, rec *HitRecord) (bool, Color, Ray) {
	attenuation := NewColor(1.0, 1.0, 1.0)
	refractionRatio := d.refractionIndex
	if rec.FrontFace {
		refractionRatio = (1.0 / d.refractionIndex)
	}
	unitDir := in.GetDirection().Unit()
	negUnitDir := in.GetDirection().Unit()
	negUnitDir.Neg()
	cosTheta := math.Min(Dot(negUnitDir, rec.Normal), 1.0)
	sinTheta := math.Sqrt(1.0 - cosTheta*cosTheta)
	cannotRefract := refractionRatio*sinTheta > 1.0
	direction := Vec3{}
	if cannotRefract || d.GetReflectance(cosTheta, refractionRatio) > rand.Float64() {
		direction = Reflect(unitDir, rec.Normal)
	} else {
		direction = Refract(unitDir, rec.Normal, refractionRatio)
	}
	scattered := NewRay(rec.P, direction)
	return true, attenuation, scattered
}

func (d *Dielectric) GetReflectance(cos, refIdx float64) float64 {
	// Use Schlick's approximation for reflectance.
	// https://en.wikipedia.org/wiki/Schlick%27s_approximation
	r0 := (1 - refIdx) / (1 + refIdx)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow(1-cos, 5)
}
