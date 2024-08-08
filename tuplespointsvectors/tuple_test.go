package tuplespointsvectors

import "testing"

func TestPointAndVector(t *testing.T) {

	point := Tuple{4.3, -4.2, 3.1, 1.0}
	if !point.IsPoint() {
		t.Error("that should be a point!")
	}

	vector := Tuple{4.3, -4.2, 3.1, 0.0}
	if !vector.IsVector() {
		t.Error("that should be a vector!")
	}

	point = NewPoint(0, 0, 0)
	if !point.IsPoint() {
		t.Error("that should be a point!")
	}

	vector = NewVector(0, 0, 0)
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
	a := NewPoint(3, 2, 1)
	b := NewPoint(5, 6, 7)

	if !SameTuple(Tuple{-2, -4, -6, 0}, Subtract(a, b)) {
		t.Error("not equal!")
	}

	b = NewVector(5, 6, 7)

	if !SameTuple(Tuple{-2, -4, -6, 1}, Subtract(a, b)) {
		t.Error("not equal!")
	}

	a = NewVector(3, 2, 1)
	if !SameTuple(Tuple{-2, -4, -6, 0}, Subtract(a, b)) {
		t.Error("not equal!")
	}

	zeroTuple := NewVector(0, 0, 0)
	a = NewVector(1, -2, 3)

	if !SameTuple(NewVector(-1, 2, -3), Subtract(zeroTuple, a)) {
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
	v := NewVector(1, 0, 0)

	if v.Magnitude() != 1 {
		t.Error("magnitude should be 1")
	}
}
