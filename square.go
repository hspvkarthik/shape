package shape

type Square struct {
	Rectangle
}

func NewSquare(length float64) Square {
	if length < 0 {
		panic("length cannot be negative")
	}
	return Square{NewRectangle(length, length)}
}
