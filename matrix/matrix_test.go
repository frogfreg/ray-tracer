package matrix

import (
	"fmt"
	"math"
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

	expected := Matrix{[]float64{1, 2, 3, 4}, []float64{5, 6, 7, 8}, []float64{9, 8, 7, 6}, []float64{5, 4, 3, 2}}

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
	aString := `
| 1 | 5 |
| -3 | 2 |
`

	a, err := MatrixFromString(aString)
	if err != nil {
		t.Error(err)
	}

	if a.Determinant() != 17 {
		t.Error("the determinant should be 17")
	}

	aString = `
| 1 | 2 | 6 |
| -5 | 8 | -4 |
| 2 | 6 | 4 |`
	a, err = MatrixFromString(aString)
	if err != nil {
		t.Error(err)
	}

	if Cofactor(a, 0, 0) != 56 {
		t.Error("wrong value")
	}
	if Cofactor(a, 0, 1) != 12 {
		t.Error("wrong value")
	}
	if Cofactor(a, 0, 2) != -46 {
		t.Error("wrong value")
	}
	if a.Determinant() != -196 {
		t.Error("wrong value")
	}

	aString = `
| -2 | -8 | 3 | 5 |
| -3 | 1 | 7 | 3 |
| 1 | 2 | -9 | 6 |
| -6 | 7 | 7 | -9 |`

	a, err = MatrixFromString(aString)
	if err != nil {
		t.Error(err)
	}

	if Cofactor(a, 0, 0) != 690 {
		t.Error("wrong value")
	}
	if Cofactor(a, 0, 1) != 447 {
		t.Error("wrong value")
	}
	if Cofactor(a, 0, 2) != 210 {
		t.Error("wrong value")
	}
	if Cofactor(a, 0, 3) != 51 {
		t.Error("wrong value")
	}
	if a.Determinant() != -4071 {
		t.Error("wrong value")
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

func TestMinor(t *testing.T) {
	aString := `
| 3 | 5 | 0 |
| 2 | -1 | -7 |
| 6 | -1 | 5 |`

	a, err := MatrixFromString(aString)
	if err != nil {
		t.Error(err)
	}

	b := Submatrix(a, 1, 0)

	if b.Determinant() != 25 {
		t.Error("determinant is wrong!")
	}

	if Minor(a, 1, 0) != b.Determinant() {
		t.Error("values should be the same")
	}
}

func TestCofactor(t *testing.T) {
	aString := `
| 3 | 5 | 0 |
| 2 | -1 | -7 |
| 6 | -1 | 5 |`

	a, err := MatrixFromString(aString)
	if err != nil {
		t.Error(err)
	}

	if Minor(a, 0, 0) != -12 {
		t.Error("wrong value")
	}
	if Cofactor(a, 0, 0) != -12 {
		t.Error("wrong value")
	}
	if Minor(a, 1, 0) != 25 {
		t.Error("wrong value")
	}
	if Cofactor(a, 1, 0) != -25 {
		t.Error("wrong value")
	}
}

func TestIsInvertible(t *testing.T) {
	aString := `
| 6 | 4 | 4 | 4 |
| 5 | 5 | 7 | 6 |
| 4 | -9 | 3 | -7 |
| 9 | 1 | 7 | -6 |`

	a, err := MatrixFromString(aString)
	if err != nil {
		t.Error(err)
	}

	if a.Determinant() != -2120 || !a.IsInvertible() {
		t.Error("matrix should be invertible")
	}
	aString = `
| -4 | 2 | -2 | -3 |
| 9 | 6 | 2 | 6 |
| 0 | -5 | 1 | -5 |
| 0 | 0 | 0 | 0 |`

	a, err = MatrixFromString(aString)
	if err != nil {
		t.Error(err)
	}

	if a.Determinant() != 0 || a.IsInvertible() {
		t.Error("matrix should not be invertible")
	}
}

func TestInverse(t *testing.T) {
	aString := `
| -5 | 2 | 6 | -8 |
| 1 | -5 | 1 | 8 |
| 7 | 7 | -6 | -7 |
| 1 | -3 | 7 | 4 |`

	a, err := MatrixFromString(aString)
	if err != nil {
		t.Error(err)
	}

	b := a.Inverse()

	if a.Determinant() != 532 {
		t.Error("wrong value")
	}
	if Cofactor(a, 2, 3) != -160 {
		t.Error("wrong value")
	}

	if b[3][2] != -(160.0 / 532.0) {
		t.Error("wrong value")
	}
	if Cofactor(a, 3, 2) != 105 {
		t.Error("wrong value")
	}
	if b[2][3] != 105.0/532.0 {
		t.Error("wrong value")
	}

	expectedString := `
| 0.21805 | 0.45113 | 0.24060 | -0.04511 |
| -0.80827 | -1.45677 | -0.44361 | 0.52068 |
| -0.07895 | -0.22368 | -0.05263 | 0.19737 |
| -0.52256 | -0.81391 | -0.30075 | 0.30639 |`

	expected, err := MatrixFromString(expectedString)
	if err != nil {
		t.Error(err)
	}

	if !AreEqual(b, expected) {
		t.Error("matrices should be equal")
	}

	aString = `
| 8 | -5 | 9 | 2 |
| 7 | 5 | 6 | 1 |
| -6 | 0 | 9 | 6 |
| -3 | 0 | -9 | -4 |`

	expectedString = `
| -0.15385 | -0.15385 | -0.28205 | -0.53846 |
| -0.07692 | 0.12308 | 0.02564 | 0.03077 |
| 0.35897 | 0.35897 | 0.43590 | 0.92308 |
| -0.69231 | -0.69231 | -0.76923 | -1.92308 |`

	a, err = MatrixFromString(aString)
	if err != nil {
		t.Error(err)
	}
	expected, err = MatrixFromString(expectedString)
	if err != nil {
		t.Error(err)
	}

	if !AreEqual(expected, a.Inverse()) {
		t.Error("matrices should be equal")
	}

	aString = `
| 9 | 3 | 0 | 9 |
| -5 | -2 | -6 | -3 |
| -4 | 9 | 6 | 4 |
| -7 | 6 | 6 | 2 |`

	expectedString = `
| -0.04074 | -0.07778 | 0.14444 | -0.22222 |
| -0.07778 | 0.03333 | 0.36667 | -0.33333 |
| -0.02901 | -0.14630 | -0.10926 | 0.12963 |
| 0.17778 | 0.06667 | -0.26667 | 0.33333 |`

	a, err = MatrixFromString(aString)
	if err != nil {
		t.Error(err)
	}
	expected, err = MatrixFromString(expectedString)
	if err != nil {
		t.Error(err)
	}

	if !AreEqual(expected, a.Inverse()) {
		t.Error("matrices should be equal")
	}

	aString = `
| 3 | -9 | 7 | 3 |
| 3 | -8 | 2 | -9 |
| -4 | 4 | 4 | 1 |
| -6 | 5 | -1 | 1 |`

	bString := `
| 8 | 2 | 2 | 2 |
| 3 | -1 | 7 | 0 |
| 7 | 0 | 5 | 4 |
| 6 | -2 | 0 | 5 |`

	a, err = MatrixFromString(aString)
	if err != nil {
		t.Error(err)
	}
	b, err = MatrixFromString(bString)
	if err != nil {
		t.Error(err)
	}

	c := Multiply(a, b)

	if !AreEqual(a, Multiply(c, b.Inverse())) {
		t.Error("matrices should be equal")
	}
}

func TestTranslation(t *testing.T) {
	trans := Translation(5, -3, 2)
	p := tpv.Point(-3, 4, 5)

	if !tpv.SameTuple(TupleMultiply(p, trans), tpv.Point(2, 1, 7)) {
		t.Error("tuples should be equal")
	}

	inv := trans.Inverse()
	if !tpv.SameTuple(TupleMultiply(p, inv), tpv.Point(-8, 7, 3)) {
		t.Error("tuples should be equal")
	}

	v := tpv.Vector(-3, 4, 5)

	if !tpv.SameTuple(TupleMultiply(v, trans), v) {
		t.Error("tuples should be equal")
	}
}

func TestScaling(t *testing.T) {
	scaleMat := Scaling(2, 3, 4)

	p := tpv.Point(-4, 6, 8)

	if !tpv.SameTuple(tpv.Point(-8, 18, 32), TupleMultiply(p, scaleMat)) {
		t.Error("points should equal")
	}

	v := tpv.Vector(-4, 6, 8)

	if !tpv.SameTuple(tpv.Vector(-8, 18, 32), TupleMultiply(v, scaleMat)) {
		t.Error("points should equal")
	}

	inv := scaleMat.Inverse()
	if !tpv.SameTuple(tpv.Vector(-2, 2, 2), TupleMultiply(v, inv)) {
		t.Error("points should equal")
	}

	// reflection
	scaleMat = Scaling(-1, 1, 1)
	p = tpv.Point(2, 3, 4)

	if !tpv.SameTuple(tpv.Point(-2, 3, 4), TupleMultiply(p, scaleMat)) {
		t.Error("points should equal")
	}
}

func TestRotationX(t *testing.T) {
	p := tpv.Point(0, 1, 0)
	halfQuarter := RotationX(math.Pi / 4)
	fullQuarter := RotationX(math.Pi / 2)

	sqrt2Over2 := math.Sqrt(2) / 2

	// Test half-quarter rotation
	expected := tpv.Point(0, sqrt2Over2, sqrt2Over2)
	result := TupleMultiply(p, halfQuarter)
	if !tpv.SameTuple(expected, result) {
		t.Errorf("Half-quarter rotation failed. Expected %v, got %v", expected, result)
	}

	// Test full-quarter rotation
	expected = tpv.Point(0, 0, 1)
	result = TupleMultiply(p, fullQuarter)
	if !tpv.SameTuple(expected, result) {
		t.Errorf("Full-quarter rotation failed. Expected %v, got %v", expected, result)
	}

	inv := halfQuarter.Inverse()
	expected = tpv.Point(0, sqrt2Over2, -sqrt2Over2)

	result = TupleMultiply(p, inv)

	if !tpv.SameTuple(result, expected) {
		t.Errorf("Full-quarter rotation failed. Expected %v, got %v", expected, result)
	}
}

func TestRotationY(t *testing.T) {
	p := tpv.Point(0, 0, 1)
	halfQuarter := RotationY(math.Pi / 4)
	fullQuarter := RotationY(math.Pi / 2)

	sqrt2Over2 := math.Sqrt(2) / 2

	// Test half-quarter rotation
	expected := tpv.Point(sqrt2Over2, 0, sqrt2Over2)
	result := TupleMultiply(p, halfQuarter)
	if !tpv.SameTuple(expected, result) {
		t.Errorf("Half-quarter rotation failed. Expected %v, got %v", expected, result)
	}

	// Test full-quarter rotation
	expected = tpv.Point(1, 0, 0)
	result = TupleMultiply(p, fullQuarter)
	if !tpv.SameTuple(expected, result) {
		t.Errorf("Full-quarter rotation failed. Expected %v, got %v", expected, result)
	}

	// Test inverse rotation
	inv := halfQuarter.Inverse()
	expected = tpv.Point(-sqrt2Over2, 0, sqrt2Over2)
	result = TupleMultiply(p, inv)
	if !tpv.SameTuple(result, expected) {
		t.Errorf("Inverse rotation failed. Expected %v, got %v", expected, result)
	}
}

func TestRotationZ(t *testing.T) {
	p := tpv.Point(0, 1, 0)
	halfQuarter := RotationZ(math.Pi / 4)
	fullQuarter := RotationZ(math.Pi / 2)

	sqrt2Over2 := math.Sqrt(2) / 2

	// Test half-quarter rotation
	expected := tpv.Point(-sqrt2Over2, sqrt2Over2, 0)
	result := TupleMultiply(p, halfQuarter)
	if !tpv.SameTuple(expected, result) {
		t.Errorf("Half-quarter rotation failed. Expected %v, got %v", expected, result)
	}

	// Test full-quarter rotation
	expected = tpv.Point(-1, 0, 0)
	result = TupleMultiply(p, fullQuarter)
	if !tpv.SameTuple(expected, result) {
		t.Errorf("Full-quarter rotation failed. Expected %v, got %v", expected, result)
	}

	// Test inverse rotation
	inv := halfQuarter.Inverse()
	expected = tpv.Point(sqrt2Over2, sqrt2Over2, 0)
	result = TupleMultiply(p, inv)
	if !tpv.SameTuple(result, expected) {
		t.Errorf("Inverse rotation failed. Expected %v, got %v", expected, result)
	}
}

func TestShearing(t *testing.T) {
	p := tpv.Point(2, 3, 4)

	tests := []struct {
		shearTrans Matrix
		expected   tpv.Tuple
	}{
		{shearTrans: Shearing(1, 0, 0, 0, 0, 0), expected: tpv.Point(5, 3, 4)},
		{shearTrans: Shearing(0, 1, 0, 0, 0, 0), expected: tpv.Point(6, 3, 4)},
		{shearTrans: Shearing(0, 0, 1, 0, 0, 0), expected: tpv.Point(2, 5, 4)},
		{shearTrans: Shearing(0, 0, 0, 1, 0, 0), expected: tpv.Point(2, 7, 4)},
		{shearTrans: Shearing(0, 0, 0, 0, 1, 0), expected: tpv.Point(2, 3, 6)},
		{shearTrans: Shearing(0, 0, 0, 0, 0, 1), expected: tpv.Point(2, 3, 7)},
	}

	for i, ts := range tests {
		t.Run(fmt.Sprintf("case %v", i), func(t *testing.T) {
			result := TupleMultiply(p, ts.shearTrans)

			if !tpv.SameTuple(result, ts.expected) {
				t.Errorf("expected %v, but got %v", ts.expected, result)
			}
		})
	}
}

func TestChaining(t *testing.T) {
	p := tpv.Point(1, 0, 1)
	aMat := RotationX(math.Pi / 2)
	bMat := Scaling(5, 5, 5)
	cMat := Translation(10, 5, 7)

	p2 := TupleMultiply(p, aMat)
	if !tpv.SameTuple(p2, tpv.Point(1, -1, 0)) {
		t.Error("tuples should be equal")
	}

	p3 := TupleMultiply(p2, bMat)
	if !tpv.SameTuple(p3, tpv.Point(5, -5, 0)) {
		t.Error("tuples should be equal")
	}

	p4 := TupleMultiply(p3, cMat)
	if !tpv.SameTuple(p4, tpv.Point(15, 0, 7)) {
		t.Error("tuples should be equal")
	}

	tMat := aMat.Scale(5, 5, 5).Translate(10, 5, 7)

	result := TupleMultiply(p, tMat)

	if !tpv.SameTuple(result, p4) {
		t.Error("tuples should be equal")
	}
}
