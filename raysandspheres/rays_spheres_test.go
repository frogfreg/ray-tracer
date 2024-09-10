package raysandspheres

import (
	"fmt"
	"testing"

	tpv "ray-tracer/tuplespointsvectors"
)

func TestNew(t *testing.T) {
	origin := tpv.Point(1, 2, 3)
	direction := tpv.Vector(4, 5, 6)

	r, err := NewRay(origin, direction)
	if err != nil {
		t.Error(err)
	}

	if !tpv.SameTuple(origin, r.Origin) {
		t.Error("tuples should be equal")
	}
	if !tpv.SameTuple(direction, r.Direction) {
		t.Error("tuples should be equal")
	}
}

func TestPosition(t *testing.T) {
	r, err := NewRay(tpv.Point(2, 3, 4), tpv.Vector(1, 0, 0))
	if err != nil {
		t.Error(err)
	}
	tests := []struct {
		time     float64
		expected tpv.Tuple
	}{
		{time: 0, expected: tpv.Point(2, 3, 4)},
		{time: 1, expected: tpv.Point(3, 3, 4)},
		{time: -1, expected: tpv.Point(1, 3, 4)},
		{time: 2.5, expected: tpv.Point(4.5, 3, 4)},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("time = %v", test.time), func(t *testing.T) {
			got := r.Position(test.time)
			if !tpv.SameTuple(got, test.expected) {
				t.Errorf("expected %v, but got %v", test.expected, got)
			}
		})
	}
}

func TestIntersect(t *testing.T) {
	tests := []struct {
		p               tpv.Tuple
		v               tpv.Tuple
		s               sphere
		expected        []float64
		shouldIntersect bool
	}{
		{p: tpv.Point(0, 0, -5), v: tpv.Vector(0, 0, 1), s: *NewSphere(), expected: []float64{4, 6}, shouldIntersect: true},
		{p: tpv.Point(0, 1, -5), v: tpv.Vector(0, 0, 1), s: *NewSphere(), expected: []float64{5, 5}, shouldIntersect: true},
		{p: tpv.Point(0, 2, -5), v: tpv.Vector(0, 0, 1), s: *NewSphere(), expected: []float64{}, shouldIntersect: false},
		{p: tpv.Point(0, 0, 0), v: tpv.Vector(0, 0, 1), s: *NewSphere(), expected: []float64{-1.0, 1, 0}, shouldIntersect: true},
		{p: tpv.Point(0, 0, 5), v: tpv.Vector(0, 0, 1), s: *NewSphere(), expected: []float64{-6, -4}, shouldIntersect: true},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test #%v", i+1), func(t *testing.T) {
			r, err := NewRay(tt.p, tt.v)
			if err != nil {
				t.Error(err)
			}

			xs := tt.s.Intersect(r)

			if !tt.shouldIntersect && len(xs) > 0 {
				t.Error("no intersection should happen")
			}

			if !tt.shouldIntersect {
				return
			}

			if len(xs) != 2 {
				t.Error("there should be 2 intersections")
			}

			if xs[0].T != tt.expected[0] {
				t.Error("wrong value")
			}
			if xs[0].T != tt.expected[0] {
				t.Error("wrong value")
			}
		})
	}
}
