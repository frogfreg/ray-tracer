package tuplespointsvectors

import (
	"math"
	"testing"
)

func TestPointAndVector(t *testing.T) {

	point := Tuple{4.3, -4.2, 3.1, 1.0}
	if !point.IsPoint() {
		t.Error("that should be a point!")
	}

	vector := Tuple{4.3, -4.2, 3.1, 0.0}
	if !vector.IsVector() {
		t.Error("that should be a vector!")
	}

	point = Point(0, 0, 0)
	if !point.IsPoint() {
		t.Error("that should be a point!")
	}

	vector = Vector(0, 0, 0)
	if !vector.IsVector() {
		t.Error("that should be a vector!")
	}
}

func TestAdd(t *testing.T) {
	a := Tuple{3, -2, 5, 1}
	b := Tuple{-2, 3, 1, 0}

	if !SameTuple(Tuple{1, 1, 6, 1}, Add(a, b)) {
		t.Error("not equal!")
	}

}
func TestSubtract(t *testing.T) {
	a := Point(3, 2, 1)
	b := Point(5, 6, 7)

	if !SameTuple(Tuple{-2, -4, -6, 0}, Subtract(a, b)) {
		t.Error("not equal!")
	}

	b = Vector(5, 6, 7)

	if !SameTuple(Tuple{-2, -4, -6, 1}, Subtract(a, b)) {
		t.Error("not equal!")
	}

	a = Vector(3, 2, 1)
	if !SameTuple(Tuple{-2, -4, -6, 0}, Subtract(a, b)) {
		t.Error("not equal!")
	}

	zeroTuple := Vector(0, 0, 0)
	a = Vector(1, -2, 3)

	if !SameTuple(Vector(-1, 2, -3), Subtract(zeroTuple, a)) {
		t.Error("not equal!")
	}

}

func TestNegated(t *testing.T) {
	a := Tuple{1, -2, 3, -4}

	if !SameTuple(Tuple{-1, 2, -3, 4}, Negated(a)) {
		t.Error("not equal!")
	}
}

func TestScMult(t *testing.T) {
	a := Tuple{1, -2, 3, -4}
	if !SameTuple(Tuple{0.5, -1, 1.5, -2}, ScMult(a, 0.5)) {
		t.Error("not equal!")
	}
	if !SameTuple(Tuple{3.5, -7, 10.5, -14}, ScMult(a, 3.5)) {
		t.Error("not equal!")
	}
}

func TestScDiv(t *testing.T) {
	a := Tuple{1, -2, 3, -4}
	if !SameTuple(Tuple{0.5, -1, 1.5, -2}, ScDiv(a, 2)) {
		t.Error("not equal!")
	}
}

func TestMagnitude(t *testing.T) {
	v := Vector(1, 0, 0)

	if v.Magnitude() != 1 {
		t.Error("magnitude should be 1")
	}
	v = Vector(0, 1, 0)
	if v.Magnitude() != 1 {
		t.Error("magnitude should be 1")
	}
	v = Vector(0, 0, 1)
	if v.Magnitude() != 1 {
		t.Error("magnitude should be 1")
	}
	v = Vector(1, 2, 3)
	if v.Magnitude() != math.Sqrt(14) {
		t.Errorf("magnitude should be %v", math.Sqrt(14))
	}
	v = Vector(-1, -2, -3)
	if v.Magnitude() != math.Sqrt(14) {
		t.Errorf("magnitude should be %v", math.Sqrt(14))
	}
}

func TestNormalized(t *testing.T) {
	v := Vector(4, 0, 0)
	if !SameTuple(Vector(1, 0, 0), Normalized(v)) {
		t.Errorf("vectors are not equal!")
	}
	v = Vector(1, 2, 3)
	nv := Normalized(v)
	testV := Vector(0.26726, 0.53452, 0.80178)

	if !SameTuple(testV, nv) {
		t.Errorf("vectors are not equal!")
	}
	if nv.Magnitude() != 1 {
		t.Error("magnitude is not 1")
	}
}

func TestDot(t *testing.T) {
	a := Vector(1, 2, 3)
	b := Vector(2, 3, 4)

	if Dot(a, b) != 20 {
		t.Error("dot product should be 20")
	}
}

func TestCross(t *testing.T) {
	a := Vector(1, 2, 3)
	b := Vector(2, 3, 4)

	if !SameTuple(Cross(a, b), Vector(-1, 2, -1)) {
		t.Error("not equal")
	}
	if !SameTuple(Cross(b, a), Vector(1, -2, 1)) {
		t.Error("not equal")
	}
}

func TestColors(t *testing.T) {
	c1 := Color(0.9, 0.6, 0.75)
	c2 := Color(0.7, 0.1, 0.25)

	if !SameTuple(Add(c1, c2), Color(1.6, 0.7, 1.0)) {
		t.Error("not the same color!")
	}

	if !SameTuple(Subtract(c1, c2), Color(0.2, 0.5, 0.5)) {
		t.Error("not the same color!")
	}

	c1 = Color(0.2, 0.3, 0.4)
	if !SameTuple(ScMult(c1, 2), Color(0.4, 0.6, 0.8)) {
		t.Error("not the same color!")
	}

}

func TestHadamardProd(t *testing.T) {
	c1 := Vector(1, 0.2, 0.4)
	c2 := Vector(0.9, 1, 0.1)

	expected := Vector(0.9, 0.2, 0.04)

	if !SameTuple(expected, HadamardProd(c1, c2)) {
		t.Error("not the same color!")
	}
}
