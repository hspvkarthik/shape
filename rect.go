package shape

type Rectangle struct {
	length float64
	width  float64
}

func NewRectangle(length float64, width float64) Rectangle {
	if length < 0 || width < 0 {
		panic("length or width should be positive")
	}
	return Rectangle{length: length, width: width}
}

func (rect Rectangle) FindPerimeter() float64 {
	if rect.length == rect.width {
		return rect.length * 4
	} else if rect.length == 1 {
		return 2 + 2*rect.width
	} else if rect.width == 1 {
		return 2 + 2*rect.length
	} else {
		return 2 * (rect.length + rect.width)
	}
}
