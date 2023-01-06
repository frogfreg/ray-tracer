package tuplespointsvectors

import "math"

type Tuple struct {
	X, Y, Z, W float64
}

func (t *Tuple) IsPoint() bool {

	return t.W == 1

}
func (t *Tuple) IsVector() bool {

	return t.W == 0

}

func NewPoint(x, y, z float64) Tuple {
	return Tuple{X: x, Y: y, Z: z, W: 1.0}

}
func NewVector(x, y, z float64) Tuple {
	return Tuple{X: x, Y: y, Z: z, W: 0.0}

}

func SameTuple(a, b Tuple) bool {

	return a.X == b.X && a.Y == b.Y && a.Z == b.Z && a.W == b.W

}

// func Add(a, b Tuple) (Tuple, error) {

// 	if a.W == 1 && b.W == 1 {
// 		return Tuple{}, errors.New("invalid operation")
// 	}

// 	return Tuple{a.X + b.X, a.Y + b.Y, a.Z + b.Z, a.W + b.W}, nil

// }

// func Subtract(a, b Tuple) (Tuple, error) {

// 	if a.W <= b.W {
// 		return Tuple{}, errors.New("invalid operation")
// 	}

// 	return Tuple{a.X - b.X, a.Y - b.Y, a.Z - b.Z, a.W - b.W}, nil

// }

func Add(a, b Tuple) Tuple {

	return Tuple{a.X + b.X, a.Y + b.Y, a.Z + b.Z, a.W + b.W}

}

func Subtract(a, b Tuple) Tuple {

	return Tuple{a.X - b.X, a.Y - b.Y, a.Z - b.Z, a.W - b.W}

}

func Negated(t Tuple) Tuple {
	return Tuple{-t.X, -t.Y, -t.Z, -t.W}
}

func (t *Tuple) Magnitude() float64 {
	return math.Sqrt(math.Pow(t.X, 2) + math.Pow(t.Y, 2) + math.Pow(t.Z, 2) + math.Pow(t.W, 2))

}

func ScMult(t Tuple, num float64) Tuple {

	return Tuple{t.X * num, t.Y * num, t.Z * num, t.W * num}
}

func ScDiv(t Tuple, num float64) Tuple {

	return Tuple{t.X / num, t.Y / num, t.Z / num, t.W / num}
}

func Normalized(t Tuple) Tuple {

	mag := t.Magnitude()

	return Tuple{t.X / mag, t.Y / mag, t.Z / mag, t.W / mag}
}

func Dot(a, b Tuple) float64 {

	return a.X*b.X + a.Y*b.Y + a.Z*b.Z + a.W*b.W

}

func Cross(a, b Tuple) Tuple {
	return NewVector(a.Y*b.Z-a.Z*b.Y,
		a.Z*b.X-a.X*b.Z,
		a.X*b.Y-a.Y*b.X)

}

func Newrgb(a, b, c float64) Tuple {
	return Tuple{a, b, c, 0}
}

func HadamardProd(c1, c2 Tuple) Tuple {

	newColor := Tuple{}

	newColor.X = c1.X * c2.X
	newColor.Y = c1.Y * c2.Y
	newColor.Z = c1.Z * c2.Z

	return newColor

}
