package matrix

import (
	"testing"

	tpv "ray-tracer/tuplespointsvectors"
)

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

func TestAreEqual(t *testing.T) {
	aMatString := `
| 1 | 2 | 3 | 4 |
| 5 | 6 | 7 | 8 |
| 9 | 8 | 7 | 6 |
| 5 | 4 | 3 | 2 |
`
	bMatString := `| 1 | 2 | 3 | 4 |
| 5 | 6 | 7 | 8 |
| 9 | 8 | 7 | 6 |
| 5 | 4 | 3 | 2 |`

	cMatString := `
| 2 | 3 | 4 | 5 |
| 6 | 7 | 8 | 9 |
| 8 | 7 | 6 | 5 |
| 4 | 3 | 2 | 1 |
`
	a, err := MatrixFromString(aMatString)
	if err != nil {
		t.Error(err)
	}
	b, err := MatrixFromString(bMatString)
	if err != nil {
		t.Error(err)
	}
	c, err := MatrixFromString(cMatString)
	if err != nil {
		t.Error(err)
	}

	if !AreEqual(a, b) {
		t.Error("both matrices should be equal!")
	}
	if AreEqual(c, b) {
		t.Error("those matrices should be different")
	}
}

func TestMultiply(t *testing.T) {
	aMatString := `
| 1 | 2 | 3 | 4 |
| 5 | 6 | 7 | 8 |
| 9 | 8 | 7 | 6 |
| 5 | 4 | 3 | 2 |
`
	bMatString := `
| -2 | 1 | 2 | 3 |
| 3 | 2 | 1 | -1 |
| 4 | 3 | 6 | 5 |
| 1 | 2 | 7 | 8 |
`

	expectedMatString := `
| 20| 22 | 50 | 48 |
| 44| 54 | 114 | 108 |
| 40| 58 | 110 | 102 |
| 16| 26 | 46 | 42 |
`

	a, err := MatrixFromString(aMatString)
	if err != nil {
		t.Error(err)
	}
	b, err := MatrixFromString(bMatString)
	if err != nil {
		t.Error(err)
	}
	expected, err := MatrixFromString(expectedMatString)
	if err != nil {
		t.Error(err)
	}

	prod := Multiply(a, b)

	if !AreEqual(expected, prod) {
		t.Error("those should be equal")
	}
}

func TestTupleMultiply(t *testing.T) {
	aMatString := `
| 1 | 2 | 3 | 4 |
| 2 | 4 | 4 | 2 |
| 8 | 6 | 4 | 1 |
| 0 | 0 | 0 | 1 |
`

	a, err := MatrixFromString(aMatString)
	if err != nil {
		t.Error(err)
	}

	tup := tpv.Tuple{X: 1, Y: 2, Z: 3, W: 1}

	prod := TupleMultiply(tup, a)

	if !tpv.SameTuple(prod, tpv.Tuple{X: 18, Y: 24, Z: 33, W: 1}) {
		t.Error("should be the same tuple")
	}
}

func TestNewIdentityMatrix(t *testing.T) {
	aMatString := `
| 0 | 1 | 2 | 4 |
| 1 | 2 | 4 | 8 |
| 2 | 4 | 8 | 16 |
| 4 | 8 | 16 | 32 |
`

	a, err := MatrixFromString(aMatString)
	if err != nil {
		t.Error(err)
	}
	iMat := NewIdentityMatrix(4, 4)

	if !AreEqual(a, Multiply(a, iMat)) {
		t.Error("should be equal!")
	}
}

func TestTranspose(t *testing.T) {
	aMatString := `
| 0 | 9 | 3 | 0 |
| 9 | 8 | 0 | 8 |
| 1 | 8 | 5 | 3 |
| 0 | 0 | 5 | 8 |`

	expectedMatString := `
| 0 | 9 | 1 | 0 |
| 9 | 8 | 8 | 0 |
| 3 | 0 | 5 | 5 |
| 0 | 8 | 3 | 8 |`

	a, err := MatrixFromString(aMatString)
	if err != nil {
		t.Error(err)
	}
	expected, err := MatrixFromString(expectedMatString)
	if err != nil {
		t.Error(err)
	}
	aTrans := Transpose(a)

	if !AreEqual(aTrans, expected) {
		t.Error("should be equal!")
	}

	if !AreEqual(NewIdentityMatrix(4, 4), Transpose(NewIdentityMatrix(4, 4))) {
		t.Error("should be equal!")
	}
}

func TestDeterminant(t *testing.T) {
	aMatString := `
| 1 | 5 |
| -3 | 2 |
`

	a, err := MatrixFromString(aMatString)
	if err != nil {
		t.Error(err)
	}

	if a.Determinant() != 17 {
		t.Error("the determinant should be 17")
	}
}

func TestSubmatrix(t *testing.T) {
	aString := `
| 1 | 5 | 0 |
| -3 | 2 | 7 |
| 0 | 6 | -3 |`

	expectedString := `
| -3 | 2 |
| 0 | 6 |`

	a, err := MatrixFromString(aString)
	if err != nil {
		t.Error(err)
	}

	expected, err := MatrixFromString(expectedString)
	if err != nil {
		t.Error(err)
	}

	if !AreEqual(Submatrix(a, 0, 2), expected) {
		t.Error(err)
	}

	aString = `
| -6 | 1 | 1 | 6 |
| -8 | 5 | 8 | 6 |
| -1 | 0 | 8 | 2 |
| -7 | 1 | -1 | 1 |`

	expectedString = `
| -6 | 1 | 6 |
| -8 | 8 | 6 |
| -7 | -1 | 1 |
`

	a, err = MatrixFromString(aString)
	if err != nil {
		t.Error(err)
	}

	expected, err = MatrixFromString(expectedString)
	if err != nil {
		t.Error(err)
	}

	if !AreEqual(Submatrix(a, 2, 1), expected) {
		t.Error(err)
	}
}
