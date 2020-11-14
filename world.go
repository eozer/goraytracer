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
	tmphr := HitRecord{}
	isHit := false
	closestSoFar := tMax
	//
	for _, obj := range w.objects {
		if obj.Hit(ray, tMin, closestSoFar, &tmphr) {
			isHit = true
			closestSoFar = rec.T
			// NOTE: We can directl use rec instead of tmphr
			rec.Normal = tmphr.Normal
			rec.P = tmphr.P
			rec.T = tmphr.T
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
