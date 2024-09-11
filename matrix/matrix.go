package matrix

import (
	"math"
	"strconv"
	"strings"

	tpv "ray-tracer/tuplespointsvectors"
)

type Matrix [][]float64

func (m Matrix) Shape() (int, int) {
	lenRows := len(m)
	lenColumns := 0

	if lenRows > 0 {
		lenColumns = len(m[0])
	}

	return lenRows, lenColumns
}

func NewMatrix(rows, columns int) Matrix {
	if rows <= 0 || columns <= 0 {
		return Matrix{}
	}

	mat := make([][]float64, rows)

	for i := range mat {
		mat[i] = make([]float64, columns)
	}

	return mat
}

func NewIdentityMatrix(rows, columns int) Matrix {
	if rows != columns {
		return Matrix{}
	}

	mat := NewMatrix(rows, columns)

	for i, row := range mat {
		for j := range row {
			if i == j {
				mat[i][j] = (1)
			}
		}
	}

	return mat
}

func AreEqual(a, b Matrix) bool {
	if len(a) == 0 && len(b) == 0 {
		return true
	}

	if len(a) != len(b) {
		return false
	}

	if len(a[0]) != len(b[0]) {
		return false
	}

	for i, row := range a {
		for j := range row {
			if !tpv.Equals(a[i][j], b[i][j]) {
				return false
			}
		}
	}

	return true
}

func Multiply(a, b Matrix) Matrix {
	newRows, _ := a.Shape()
	_, newCols := b.Shape()

	new := NewMatrix(newRows, newCols)

	for rowIndex := 0; rowIndex < 4; rowIndex++ {
		for colIndex := 0; colIndex < 4; colIndex++ {
			new[rowIndex][colIndex] = a[rowIndex][0]*b[0][colIndex] +
				a[rowIndex][1]*b[1][colIndex] +
				a[rowIndex][2]*b[2][colIndex] +
				a[rowIndex][3]*b[3][colIndex]
		}
	}

	return new
}

func TupleMultiply(a tpv.Tuple, b Matrix) tpv.Tuple {
	var tempSlice [4]float64

	for rowIndex, row := range b {
		tupleB := tpv.Tuple{X: row[0], Y: row[1], Z: row[2], W: row[3]}

		tempSlice[rowIndex] = tpv.Dot(a, tupleB)
	}

	return tpv.Tuple{X: tempSlice[0], Y: tempSlice[1], Z: tempSlice[2], W: tempSlice[3]}
}

func Transpose(m Matrix) Matrix {
	rows, cols := m.Shape()

	new := NewMatrix(cols, rows)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			new[j][i] = m[i][j]
		}
	}

	return new
}

func Submatrix(m Matrix, delRow, delCol int) Matrix {
	new := Matrix{}

	for rowIndex := 0; rowIndex < len(m); rowIndex++ {
		if rowIndex == delRow {
			continue
		}

		rowCopy := append([]float64{}, m[rowIndex][:delCol]...)

		newRow := append(rowCopy, m[rowIndex][delCol+1:]...)

		new = append(new, newRow)
	}

	return new
}

func Minor(m Matrix, delRow, delCol int) float64 {
	return Submatrix(m, delRow, delCol).Determinant()
}

func Cofactor(m Matrix, delRow, delCol int) float64 {
	multiplier := 1.0

	if (delRow+delCol)%2 != 0 {
		multiplier = -1.0
	}

	return (multiplier) * (Minor(m, delRow, delCol))
}

func (m Matrix) Determinant() float64 {
	rows, cols := m.Shape()

	if rows == 2 && cols == 2 {
		return m[0][0]*(m[1][1]) - (m[0][1] * (m[1][0]))
	}

	det := 0.0

	for colIndex, num := range m[0] {
		det = det + (num * Cofactor(m, 0, colIndex))
	}

	return det
}

func (m Matrix) IsInvertible() bool {
	return !tpv.Equals(m.Determinant(), 0.0)
}

func (m Matrix) Inverse() Matrix {
	det := m.Determinant()

	rows, cols := m.Shape()

	new := NewMatrix(rows, cols)

	for i := range new {
		for j := range new[0] {
			new[i][j] = Cofactor(m, i, j)
		}
	}

	new = Transpose(new)

	for i := range new {
		for j := range new[0] {
			new[i][j] = new[i][j] / (det)
		}
	}

	return new
}

func Translation(x, y, z float64) Matrix {
	new := NewIdentityMatrix(4, 4)

	new[0][3] = (x)
	new[1][3] = (y)
	new[2][3] = (z)

	return new
}

func (m Matrix) Translate(x, y, z float64) Matrix {
	new := NewIdentityMatrix(4, 4)

	new[0][3] = (x)
	new[1][3] = (y)
	new[2][3] = (z)

	res := Multiply(new, m)

	return res
}

func Scaling(x, y, z float64) Matrix {
	new := NewIdentityMatrix(4, 4)

	new[0][0] = (x)
	new[1][1] = (y)
	new[2][2] = (z)

	return new
}

func (m Matrix) Scale(x, y, z float64) Matrix {
	new := NewIdentityMatrix(4, 4)

	new[0][0] = (x)
	new[1][1] = (y)
	new[2][2] = (z)

	res := Multiply(new, m)

	return res
}

func RotationX(radians float64) Matrix {
	new := NewIdentityMatrix(4, 4)

	new[1][1] = (math.Cos(radians))
	new[1][2] = (-math.Sin(radians))
	new[2][1] = (math.Sin(radians))
	new[2][2] = (math.Cos(radians))

	return new
}

func (m Matrix) RotateX(radians float64) Matrix {
	new := NewIdentityMatrix(4, 4)

	new[1][1] = (math.Cos(radians))
	new[1][2] = (-math.Sin(radians))
	new[2][1] = (math.Sin(radians))
	new[2][2] = (math.Cos(radians))

	res := Multiply(new, m)

	return res
}

func RotationY(radians float64) Matrix {
	new := NewIdentityMatrix(4, 4)

	new[0][0] = (math.Cos(radians))
	new[0][2] = (math.Sin(radians))
	new[2][0] = (-math.Sin(radians))
	new[2][2] = (math.Cos(radians))

	return new
}

func (m Matrix) RotateY(radians float64) Matrix {
	new := NewIdentityMatrix(4, 4)

	new[0][0] = (math.Cos(radians))
	new[0][2] = (math.Sin(radians))
	new[2][0] = (-math.Sin(radians))
	new[2][2] = (math.Cos(radians))

	res := Multiply(new, m)

	return res
}

func RotationZ(radians float64) Matrix {
	new := NewIdentityMatrix(4, 4)

	new[0][0] = (math.Cos(radians))
	new[0][1] = (-math.Sin(radians))
	new[1][0] = (math.Sin(radians))
	new[1][1] = (math.Cos(radians))

	return new
}

func (m Matrix) RotateZ(radians float64) Matrix {
	new := NewIdentityMatrix(4, 4)

	new[0][0] = (math.Cos(radians))
	new[0][1] = (-math.Sin(radians))
	new[1][0] = (math.Sin(radians))
	new[1][1] = (math.Cos(radians))

	res := Multiply(new, m)

	return res
}

func Shearing(xy, xz, yx, yz, zx, zy float64) Matrix {
	new := NewIdentityMatrix(4, 4)

	new[0][1] = (xy)
	new[0][2] = (xz)
	new[1][0] = (yx)
	new[1][2] = (yz)
	new[2][0] = (zx)
	new[2][1] = (zy)

	return new
}

func (m Matrix) Shear(xy, xz, yx, yz, zx, zy float64) Matrix {
	new := NewIdentityMatrix(4, 4)

	new[0][1] = (xy)
	new[0][2] = (xz)
	new[1][0] = (yx)
	new[1][2] = (yz)
	new[2][0] = (zx)
	new[2][1] = (zy)

	res := Multiply(new, m)

	return res
}

func MatrixFromString(matString string) (Matrix, error) {
	m := Matrix{}

	for _, row := range strings.Split(matString, "\n") {
		if strings.TrimSpace(row) == "" {
			continue
		}
		var nums []float64

		for _, numString := range strings.Split(strings.ReplaceAll(row, " ", ""), "|") {
			if numString == "" {
				continue
			}
			num, err := strconv.ParseFloat(numString, 64)
			if err != nil {
				return nil, err
			}

			nums = append(nums, num)
		}

		m = append(m, nums)
	}

	return m, nil
}

func (m Matrix) String() string {
	s := ""

	for _, row := range m {
		line := ""
		for _, col := range row {
			line += strconv.FormatFloat(col, 'f', -1, 64) + " "
		}
		s += line + "\n"
	}

	return "\n" + strings.TrimSpace(s)
}
