package matrix

import (
	tpv "ray-tracer/tuplespointsvectors"
)

type matrix [][]float64

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

	mat := make([][]float64, rows)

	for i := range mat {
		mat[i] = make([]float64, columns)
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
				mat[i][j] = 1
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
			if a[i][j] != b[i][j] {
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
			new[rowIndex][colIndex] = a[rowIndex][0]*b[0][colIndex] +
				a[rowIndex][1]*b[1][colIndex] +
				a[rowIndex][2]*b[2][colIndex] +
				a[rowIndex][3]*b[3][colIndex]

		}
	}

	return new

}

func TupleMultiply(a tpv.Tuple, b matrix) tpv.Tuple {

	var tempSlice [4]float64

	for rowIndex, row := range b {
		tupleB := tpv.Tuple{X: row[0], Y: row[1], Z: row[2], W: row[3]}

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

func simpleDeterminant(m matrix) float64 {
	return m[0][0]*m[1][1] - (m[0][1] * m[1][0])
}

func Submatrix(m matrix, delRow, delCol int) matrix {

	new := matrix{}

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

func Minor(m matrix, delRow, delCol int) float64 {

	return simpleDeterminant(Submatrix(m, delRow, delCol))

}

func Cofactor(m matrix, delRow, delCol int) float64 {
	multiplier := 1.0

	if (delRow+delCol)%2 != 0 {
		multiplier = -1.0
	}

	return multiplier * Minor(m, delRow, delCol)
}

// func (m matrix) Determinant() float64 {

// 	rows, cols := m.Shape()

// 	if rows == 2 && cols == 2 {
// 		return simpleDeterminant(m)
// 	}

// 	det := 0.0

// 	for colIndex, num := range m[0] {

// 		det += (num * Cofactor(m, 0, colIndex))
// 	}

// 	return det

// }
