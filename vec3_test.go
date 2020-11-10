package goraytracer

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	for _, test := range []struct {
		v1   Vec3
		v2   Vec3
		want Vec3
	}{
		// test add when there is an unitialized vector
		{Vec3{}, Vec3{[3]float64{1.0, 2.0, 3.0}}, Vec3{[3]float64{1.0, 2.0, 3.0}}},
		// test basic addition
		{Vec3{[3]float64{-1.0, 2.0, -3.0}}, Vec3{[3]float64{1.0, 2.0, 3.0}}, Vec3{[3]float64{0.0, 4.0, 0.0}}},
	} {
		got := Add(&test.v1, &test.v2)
		if !reflect.DeepEqual(test.want, got) {
			t.Errorf("want %v, got %v", test.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	for _, test := range []struct {
		v1   Vec3
		v2   Vec3
		want Vec3
	}{
		// test add when there is an unitialized vector
		{Vec3{}, Vec3{[3]float64{1.0, 2.0, 3.0}}, Vec3{[3]float64{-1.0, -2.0, -3.0}}},
		// test basic addition
		{Vec3{[3]float64{-1.0, 2.0, -3.0}}, Vec3{[3]float64{-1.5, -2.0, -3.5}}, Vec3{[3]float64{0.5, 4.0, 0.5}}},
	} {
		got := Subtract(&test.v1, &test.v2)
		if !reflect.DeepEqual(test.want, got) {
			t.Errorf("want %v, got %v", test.want, got)
		}
	}
}
