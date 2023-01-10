package matrix

type FbFmatrix [4][4]float64
type TbTmatrix [3][3]float64
type SquareMatrix [2][2]float64

func (m *FbFmatrix) Shape() (int, string) {
	return 4, "4x4"
}
func (m *TbTmatrix) Shape() (int, string) {
	return 3, "3x3"
}
func (m *SquareMatrix) Shape() (int, string) {
	return 2, "2x2"
}

type Matrix interface {
	Shape() (int, string)
}

func AreEqual[M Matrix](a, b M) bool {

	aSize, aShape := a.Shape()
	bSize, bShape := b.Shape()

	if aShape != bShape {
		return false
	}

	for i := 0; i < aSize; i++ {
		for j := 0; j < bSize; j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}

	return true

}

// func NewFbFMatrix() FbFmatrix {

// 	var nm [4][4]floaany

// 	return nm
// }
// func NewTbTMatrix() TbTmatrix {

// 	var nm [3][3]float64

// 	return nm
// }
// func NewSquareMatrix() Matrix {

// 	var nm [2][2]float64

// 	return nm
// }
