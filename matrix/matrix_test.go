package matrix

import "testing"

func TestMatrixSizes(t *testing.T) {
	m := NewMatrix(2, 2)

	if len(m[0]) != 2 {
		t.Error("wrong size")
	}
	if len(m) != 2 {
		t.Error("wrong size")
	}

	m = NewMatrix(4, 4)

	if len(m[0]) != 4 {
		t.Error("wrong size")
	}
	if len(m) != 4 {
		t.Error("wrong size")
	}
}
