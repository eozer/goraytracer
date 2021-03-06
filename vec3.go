package goraytracer

import (
	"math"
	"math/rand"
)

type Vec3 struct {
	axes [3]float64
}

// Point3 is an alias for Vec3
type Point3 = Vec3

func NewVec3(x, y, z float64) Vec3 {
	return Vec3{axes: [3]float64{x, y, z}}
}

func NewPoint3(x, y, z float64) Vec3 {
	return NewVec3(x, y, z)
}

// NewRandomVec3 returns a Vec3 in [0.0, 1.0) range.
func NewRandomVec3() Vec3 {
	return NewVec3(rand.Float64(), rand.Float64(), rand.Float64())
}

// NewRandomVec3From returns a Vec3 in [min, max) range.
func NewRandomVec3From(min, max float64) Vec3 {
	x := min + (max-min)*rand.Float64()
	y := min + (max-min)*rand.Float64()
	z := min + (max-min)*rand.Float64()
	return NewVec3(x, y, z)
}

// NewRandomVec3InUnitSphere returns a Vec3 that is inside the unit sphere.
func NewRandomVec3InUnitSphere() Vec3 {
	for {
		v := NewRandomVec3From(-1.0, 1.0)
		if v.SqrLen() >= 1.0 {
			continue
		}
		return v
	}
}

func NewRandomVec3Unit() Vec3 {
	return UnitVector(NewRandomVec3InUnitSphere())
}

func NewRandomVec3InHemisphere(normal Vec3) Vec3 {
	unitsph := NewRandomVec3InUnitSphere()
	// if the unit vector is in the same direction as normal vector
	if Dot(unitsph, normal) > 0.0 {
		// empty on purpose
	} else {
		unitsph.Neg()
	}
	return unitsph
}

func NewRandomVec3InUnitDisk() Vec3 {
	for {
		x := -1.0 + 2.0*rand.Float64()
		y := -1.0 + 2.0*rand.Float64()
		z := 0.0
		v := NewVec3(x, y, z)
		if v.SqrLen() >= 1 {
			continue
		}
		return v
	}
}

func Add(v1, v2 Vec3) Vec3 {
	v1.Add(v2)
	return v1
}

func Subtract(v1, v2 Vec3) Vec3 {
	v1.Sub(v2)
	return v1
}

func Mult(t float64, v Vec3) Vec3 {
	v.Mult(t)
	return v
}

func Div(t float64, v Vec3) Vec3 {
	v.Div(t)
	return v
}

func ElemMult(v1, v2 Vec3) Vec3 {
	return Vec3{[3]float64{v1.axes[0] * v2.axes[0], v1.axes[1] * v2.axes[1], v1.axes[2] * v2.axes[2]}}
}

func Dot(v1, v2 Vec3) float64 {
	return (v1.axes[0] * v2.axes[0]) + (v1.axes[1] * v2.axes[1]) + (v1.axes[2] * v2.axes[2])
}

func Cross(v1, v2 Vec3) Vec3 {
	return Vec3{[3]float64{
		(v1.axes[1]*v2.axes[2] - v1.axes[2]*v2.axes[1]),
		(v1.axes[2]*v2.axes[0] - v1.axes[0]*v2.axes[2]),
		(v1.axes[0]*v2.axes[1] - v1.axes[1]*v2.axes[0]),
	}}
}

func UnitVector(v Vec3) Vec3 {
	v.Div(v.Len())
	return v
}

func Reflect(v, n Vec3) Vec3 {
	return Subtract(v, Mult(2*Dot(v, n), n))
}

func Refract(uv, n Vec3, etaI_over_etaT float64) Vec3 {
	negUv := uv.Clone()
	negUv.Neg()
	cosTheta := math.Min(1.0, Dot(negUv, n))
	rPerpendicular := Mult(etaI_over_etaT, Add(uv, Mult(cosTheta, n)))
	rParallel := Mult(-math.Sqrt(math.Abs(1.0-rPerpendicular.SqrLen())), n)
	return Add(rPerpendicular, rParallel)
}

type Vector interface {
	Clone() Vec3
	GetX() float64
	GetY() float64
	GetZ() float64
	GetAxe(i int) float64
	Neg() *Vec3
	Add(Vec3) *Vec3
	Sub(Vec3) *Vec3
	Mult(float64) *Vec3
	Div(float64) *Vec3
	Len() float64
	SqrLen() float64
	Unit(*Vec3) Vec3
	IsNearZero() bool
}

func (v *Vec3) Clone() Vec3 {
	return Vec3{[3]float64{v.axes[0], v.axes[1], v.axes[2]}}
}

func (v *Vec3) GetX() float64 {
	return v.axes[0]
}

func (v *Vec3) GetY() float64 {
	return v.axes[1]
}

func (v *Vec3) GetZ() float64 {
	return v.axes[2]
}

func (v *Vec3) GetAxe(i int) float64 {
	// TODO: Handle if i given outside of ranges in [0,2]
	return v.axes[i]
}

func (v *Vec3) Neg() *Vec3 {
	v.axes[0] = -v.axes[0]
	v.axes[1] = -v.axes[1]
	v.axes[2] = -v.axes[2]
	return v
}

func (v *Vec3) Add(v2 Vec3) *Vec3 {
	v.axes[0] += v2.axes[0]
	v.axes[1] += v2.axes[1]
	v.axes[2] += v2.axes[2]
	return v
}

func (v *Vec3) Sub(v2 Vec3) *Vec3 {
	v.axes[0] -= v2.axes[0]
	v.axes[1] -= v2.axes[1]
	v.axes[2] -= v2.axes[2]
	return v
}

func (v *Vec3) Mult(t float64) *Vec3 {
	v.axes[0] *= t
	v.axes[1] *= t
	v.axes[2] *= t
	return v
}

func (v *Vec3) Div(t float64) *Vec3 {
	v.Mult(1 / t)
	return v
}

func (v *Vec3) Unit() Vec3 {
	clone := v.Clone()
	clone.Div(v.Len())
	return clone
}

func (v *Vec3) Len() float64 {
	return math.Sqrt(v.SqrLen())
}

func (v *Vec3) SqrLen() float64 {
	return v.axes[0]*v.axes[0] + v.axes[1]*v.axes[1] + v.axes[2]*v.axes[2]
}

// IsNearZero returns true if all of the axes of v are near to zero.
func (v *Vec3) IsNearZero() bool {
	const tol float64 = 1e-8
	return (math.Abs(v.axes[0]) < tol) && (math.Abs(v.axes[1]) < tol) && (math.Abs(v.axes[2]) < tol)
}
