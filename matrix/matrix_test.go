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
func TestMatrixFromString(t *testing.T) {

	matString := `
| 1 | 2 | 3 | 4 |
| 5 | 6 | 7 | 8 |
| 9 | 8 | 7 | 6 |
| 5 | 4 | 3 | 2 |
`

	expected := matrix{[]float64{1, 2, 3, 4}, []float64{5, 6, 7, 8}, []float64{9, 8, 7, 6}, []float64{5, 4, 3, 2}}

	m, err := MatrixFromString(matString)
	if err != nil {
		t.Error(err)
	}

	if !AreEqual(m, expected) {
		t.Errorf("expected %v, but got %v", expected, m)
	}
}
