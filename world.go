package goraytracer

// HittableList is a type that represents
type HittableList interface {
	Hittable
	Clear()
	Add(Hittable)
}

// World is a type that stores slice of hittable objects
type World struct {
	objects []Hittable
}

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
