package raysandspheres

import (
	"fmt"
	"testing"

	"ray-tracer/matrix"
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
		transformMat    matrix.Matrix
	}{
		{p: tpv.Point(0, 0, -5), v: tpv.Vector(0, 0, 1), s: *NewSphere(), expected: []float64{4, 6}, shouldIntersect: true},
		{p: tpv.Point(0, 1, -5), v: tpv.Vector(0, 0, 1), s: *NewSphere(), expected: []float64{5, 5}, shouldIntersect: true},
		{p: tpv.Point(0, 2, -5), v: tpv.Vector(0, 0, 1), s: *NewSphere(), expected: []float64{}, shouldIntersect: false},
		{p: tpv.Point(0, 0, 0), v: tpv.Vector(0, 0, 1), s: *NewSphere(), expected: []float64{-1.0, 1, 0}, shouldIntersect: true},
		{p: tpv.Point(0, 0, 5), v: tpv.Vector(0, 0, 1), s: *NewSphere(), expected: []float64{-6, -4}, shouldIntersect: true},
		{p: tpv.Point(0, 0, -5), v: tpv.Vector(0, 0, 1), s: *NewSphere(), expected: []float64{3, 7}, shouldIntersect: true, transformMat: matrix.Scaling(2, 2, 2)},
		{p: tpv.Point(0, 0, -5), v: tpv.Vector(0, 0, 1), s: *NewSphere(), expected: []float64{}, shouldIntersect: false, transformMat: matrix.Translation(5, 0, 0)},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test #%v", i+1), func(t *testing.T) {
			r, err := NewRay(tt.p, tt.v)
			if err != nil {
				t.Error(err)
			}

			if tt.transformMat != nil {
				tt.s.SetTransform(tt.transformMat)
			}

			xs := tt.s.Intersect(r)

			if r.Origin != tt.p || r.Direction != tt.v {
				t.Errorf("expected %v, but got %v", ray{tt.p, tt.v}, r)
			}

			if !tt.shouldIntersect && len(xs) > 0 {
				t.Error("no intersection should happen")
			}

			if !tt.shouldIntersect {
				return
			}

			if len(xs) != 2 {
				t.Error("there should be 2 intersections")
			}

			if xs[0].TValue != tt.expected[0] {
				t.Error("wrong value")
			}
			if xs[0].TValue != tt.expected[0] {
				t.Error("wrong value")
			}
		})
	}
}

func TestNewIntersection(t *testing.T) {
	s := NewSphere()

	i := NewIntersection(3.5, s)

	if i.TValue != 3.5 {
		t.Errorf("expected 3.5, but got %v", i.TValue)
	}

	if i.Object != s {
		t.Error("objects are not equal")
	}
}

func TestIntersections(t *testing.T) {
	s := NewSphere()

	i1 := NewIntersection(1, s)
	i2 := NewIntersection(2, s)

	xs := NewCollection(i1, i2)

	expected := 2

	if len(xs) != expected {
		t.Errorf("expected %v, but got %v", expected, len(xs))
	}

	if xs[0].TValue != 1 {
		t.Error("wrong value")
	}
	if xs[1].TValue != 2 {
		t.Error("wrong value")
	}
}

func TestHit(t *testing.T) {
	s := NewSphere()

	tests := []struct {
		col         collection
		expected    intersection
		errExpected bool
	}{
		{col: NewCollection(NewIntersection(1, s), NewIntersection(2, s)), expected: NewIntersection(1, s)},
		{col: NewCollection(NewIntersection(-1, s), NewIntersection(1, s)), expected: NewIntersection(1, s)},
		{col: NewCollection(NewIntersection(-2, s), NewIntersection(-1, s)), expected: intersection{}, errExpected: true},
		{col: NewCollection(NewIntersection(5, s), NewIntersection(7, s), NewIntersection(-3, s), NewIntersection(2, s)), expected: NewIntersection(2, s)},
	}

	for ind, tt := range tests {
		t.Run(fmt.Sprintf("test #%v", ind+1), func(t *testing.T) {
			i, err := tt.col.Hit()
			if err != nil && !tt.errExpected {
				t.Error(err)
			}

			if tt.errExpected && err == nil {
				t.Error("expected error, but got nil")
			}

			if tt.errExpected {
				return
			}

			if i.TValue != tt.expected.TValue {
				t.Errorf("expected %v, but got %v", tt.expected, i)
			}
		})
	}
}

func TestTransform(t *testing.T) {
	tests := []struct {
		matTransform matrix.Matrix
		origin       tpv.Tuple
		direction    tpv.Tuple
		expected     ray
	}{
		{matTransform: matrix.Translation(3, 4, 5), origin: tpv.Point(1, 2, 3), direction: tpv.Vector(0, 1, 0), expected: ray{tpv.Point(4, 6, 8), tpv.Vector(0, 1, 0)}},
		{matTransform: matrix.Scaling(2, 3, 4), origin: tpv.Point(1, 2, 3), direction: tpv.Vector(0, 1, 0), expected: ray{tpv.Point(2, 6, 12), tpv.Vector(0, 3, 0)}},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test #%v", i+1), func(t *testing.T) {
			r, err := NewRay(tt.origin, tt.direction)
			if err != nil {
				t.Error(err)
			}

			got := r.Transform(tt.matTransform)

			if got != tt.expected {
				t.Errorf("expected %v, but got %v", tt.expected, got)
			}
		})
	}
}

func TestSetTransform(t *testing.T) {
	s := NewSphere()

	if !matrix.AreEqual(s.transformMat, matrix.NewIdentityMatrix(4, 4)) {
		t.Error("default matrix should be identity matrix")
	}

	transMat := matrix.Translation(2, 3, 4)
	s.SetTransform(transMat)

	if !matrix.AreEqual(s.transformMat, transMat) {
		t.Errorf("expected %v, but got %v", transMat, s.transformMat)
	}
}
