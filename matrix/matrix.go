package matrix

type matrix [][]float64

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
