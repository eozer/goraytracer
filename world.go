package goraytracer

import "math/rand"

// HittableList is a type that represents
type HittableList interface {
	Hittable
	Clear()
	Add(Hittable)
}

func MakeRandomWorld() HittableList {
	world := World{}
	groundMat := NewLambertian(NewColor(.5, .5, .5))
	groundSphere := NewSphere(NewPoint3(0, -1000.0, 0), 1000.0, &groundMat)
	world.Add(&groundSphere)
	// Create a world of random objects
	for i := -5; i < 5; i++ {
		for j := -5; j < 5; j++ {
			objCenter := NewPoint3(float64(i)+0.9*rand.Float64(), 0.2, float64(j)+0.9*rand.Float64())
			chooseMat := rand.Float64()
			if chooseMat < 0.45 {
				mat := NewLambertian(NewRandomVec3())
				sphere := NewSphere(objCenter, 0.2, &mat)
				world.Add(&sphere)
			} else if chooseMat < 0.85 {
				mat := NewMetal(NewRandomVec3From(0.5, 1.0), rand.Float64())
				sphere := NewSphere(objCenter, 0.2, &mat)
				world.Add(&sphere)
			} else {
				mat := NewDielectric(1.5)
				sphere := NewSphere(objCenter, 0.2, &mat)
				world.Add(&sphere)
			}
		}
	}

	mat1 := NewDielectric(1.5)
	sph1 := NewSphere(NewPoint3(0, 1, 0), 1.0, &mat1)
	world.Add(&sph1)

	mat2 := NewMetal(NewColor(0.7, 0.6, 0.5), 0)
	sph2 := NewSphere(NewPoint3(3.5, 1, -0.5), 1.0, &mat2)
	world.Add(&sph2)

	mat3 := NewLambertian(NewColor(0.4, 0.2, 0.1))
	sph3 := NewSphere(NewPoint3(-4, 1, 0), 1.0, &mat3)
	world.Add(&sph3)

	return &world
}

// World is a type that stores slice of hittable objects
type World struct {
	// objects holds a list of Hittable objects.
	objects []Hittable
}

// Hit returns true if ray in [tMin, tMax) hits to an object in the world.
func (w *World) Hit(ray *Ray, tMin float64, tMax float64, rec *HitRecord) bool {
	isHit := false
	closestSoFar := tMax
	//
	for _, obj := range w.objects {
		if obj.Hit(ray, tMin, closestSoFar, rec) {
			isHit = true
			closestSoFar = rec.T
		}
	}

	return isHit
}

func (w *World) Clear() {
	w.objects = nil
}

func (w *World) Add(h Hittable) {
	w.objects = append(w.objects, h)
}
