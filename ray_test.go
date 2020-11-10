package goraytracer

import (
	"reflect"
	"testing"
)

func TestRayAt(t *testing.T) {
	for _, test := range []struct {
		ray  Ray
		t    float64
		want Point3
	}{
		{NewRay(Vec3{}, NewVec3(1.0, 1.0, 1.0)), 2.0, NewVec3(2.0, 2.0, 2.0)},
	} {
		got := test.ray.At(test.t)
		if !reflect.DeepEqual(test.want, got) {
			t.Errorf("want %v, got %v", test.want, got)
		}
	}
}
