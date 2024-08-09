package tuplespointsvectors

import "testing"

func TestEquals(t *testing.T) {
	if equals(2.0, 1.9) {
		t.Error("2.0 should not be equal to 1.9!")
	}

	if equals(1.9, 2.0) {
		t.Error("2.0 should not be equal to 1.9!")
	}

	if !equals(2.0, 1.999999) {
		t.Error("they should be equal for this library")
	}
	if !equals(1.999999, 2.0) {
		t.Error("they should be equal for this library")
	}
	if !equals(0.26726, 0.2672612419124244) {
		t.Error("they should be equal for this library")
	}
}