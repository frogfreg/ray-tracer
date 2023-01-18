package matrix

import (
	"math"
	tpv "ray-tracer/tuplespointsvectors"

	"github.com/shopspring/decimal"
)

type matrix [][]decimal.Decimal

func same(a, b decimal.Decimal) bool {
	epsilon := decimal.NewFromFloat(0.0000001)

	return a.Sub(b).Abs().LessThan(epsilon)
}

func (m matrix) Shape() (int, int) {
	lenRows := len(m)
	lenColumns := 0

	if lenRows > 0 {
		lenColumns = len(m[0])
	}

	return lenRows, lenColumns

}

func NewMatrix(rows, columns int) matrix {

	if rows <= 0 || columns <= 0 {
		return matrix{}
	}

	mat := make([][]decimal.Decimal, rows)

	for i := range mat {
		mat[i] = make([]decimal.Decimal, columns)
	}

	return mat
}

func NewIdentityMatrix(rows, columns int) matrix {

	if rows != columns {
		return matrix{}
	}

	mat := NewMatrix(rows, columns)

	for i, row := range mat {
		for j := range row {
			if i == j {
				mat[i][j] = decimal.NewFromInt(1)
			}
		}
	}

	return mat
}

func AreEqual(a, b matrix) bool {

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

			if !same(a[i][j], b[i][j]) {
				return false
			}
		}
	}

	return true

}

func Multiply(a, b matrix) matrix {

	newRows, _ := a.Shape()
	_, newCols := b.Shape()

	new := NewMatrix(newRows, newCols)

	for rowIndex := 0; rowIndex < 4; rowIndex++ {
		for colIndex := 0; colIndex < 4; colIndex++ {
			new[rowIndex][colIndex] = decimal.Sum(a[rowIndex][0].Mul(b[0][colIndex]),
				a[rowIndex][1].Mul(b[1][colIndex]),
				a[rowIndex][2].Mul(b[2][colIndex]),
				a[rowIndex][3].Mul(b[3][colIndex]))

		}
	}

	return new

}

func TupleMultiply(a tpv.Tuple, b matrix) tpv.Tuple {

	var tempSlice [4]float64

	for rowIndex, row := range b {
		tupleB := tpv.Tuple{X: row[0].InexactFloat64(), Y: row[1].InexactFloat64(), Z: row[2].InexactFloat64(), W: row[3].InexactFloat64()}

		tempSlice[rowIndex] = tpv.Dot(a, tupleB)
	}

	return tpv.Tuple{X: tempSlice[0], Y: tempSlice[1], Z: tempSlice[2], W: tempSlice[3]}

}

func Transpose(m matrix) matrix {

	rows, cols := m.Shape()

	new := NewMatrix(cols, rows)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			new[j][i] = m[i][j]
		}
	}

	return new
}

func Submatrix(m matrix, delRow, delCol int) matrix {

	new := matrix{}

	for rowIndex := 0; rowIndex < len(m); rowIndex++ {
		if rowIndex == delRow {
			continue
		}

		rowCopy := append([]decimal.Decimal{}, m[rowIndex][:delCol]...)

		newRow := append(rowCopy, m[rowIndex][delCol+1:]...)

		new = append(new, newRow)
	}

	return new

}

func Minor(m matrix, delRow, delCol int) decimal.Decimal {

	return Submatrix(m, delRow, delCol).Determinant()

}

func Cofactor(m matrix, delRow, delCol int) decimal.Decimal {
	multiplier := 1.0

	if (delRow+delCol)%2 != 0 {
		multiplier = -1.0
	}

	return decimal.NewFromFloat(multiplier).Mul(Minor(m, delRow, delCol))
}

func (m matrix) Determinant() decimal.Decimal {

	rows, cols := m.Shape()

	if rows == 2 && cols == 2 {
		return m[0][0].Mul(m[1][1]).Sub(m[0][1].Mul(m[1][0]))

	}

	det := decimal.Zero

	for colIndex, num := range m[0] {

		det = det.Add((num.Mul(Cofactor(m, 0, colIndex))))
	}

	return det

}

func (m matrix) IsInvertible() bool {
	return !m.Determinant().IsZero()
}

func (m matrix) Inverse() matrix {

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
			new[i][j] = new[i][j].Div(det)
		}
	}

	return new

}

func Translation(x, y, z float64) matrix {
	new := NewIdentityMatrix(4, 4)

	new[0][3] = decimal.NewFromFloat(x)
	new[1][3] = decimal.NewFromFloat(y)
	new[2][3] = decimal.NewFromFloat(z)

	return new
}

func Scaling(x, y, z float64) matrix {
	new := NewIdentityMatrix(4, 4)

	new[0][0] = decimal.NewFromFloat(x)
	new[1][1] = decimal.NewFromFloat(y)
	new[2][2] = decimal.NewFromFloat(z)

	return new

}

func RotateX(radians float64) matrix {
	new := NewIdentityMatrix(4, 4)

	new[1][1] = decimal.NewFromFloat(math.Cos(radians))
	new[1][2] = decimal.NewFromFloat(-math.Sin(radians))
	new[2][1] = decimal.NewFromFloat(math.Sin(radians))
	new[2][2] = decimal.NewFromFloat(math.Cos(radians))

	return new

}

func RotateY(radians float64) matrix {
	new := NewIdentityMatrix(4, 4)

	new[0][0] = decimal.NewFromFloat(math.Cos(radians))
	new[0][2] = decimal.NewFromFloat(math.Sin(radians))
	new[2][0] = decimal.NewFromFloat(-math.Sin(radians))
	new[2][2] = decimal.NewFromFloat(math.Cos(radians))

	return new

}

func RotateZ(radians float64) matrix {
	new := NewIdentityMatrix(4, 4)

	new[0][0] = decimal.NewFromFloat(math.Cos(radians))
	new[0][1] = decimal.NewFromFloat(-math.Sin(radians))
	new[1][0] = decimal.NewFromFloat(math.Sin(radians))
	new[1][1] = decimal.NewFromFloat(math.Cos(radians))

	return new

}

func Shearing(xy, xz, yx, yz, zx, zy float64) matrix {
	new := NewIdentityMatrix(4, 4)

	new[0][1] = decimal.NewFromFloat(xy)
	new[0][2] = decimal.NewFromFloat(xz)
	new[1][0] = decimal.NewFromFloat(yx)
	new[1][2] = decimal.NewFromFloat(yz)
	new[2][0] = decimal.NewFromFloat(zx)
	new[2][1] = decimal.NewFromFloat(zy)

	return new

}
