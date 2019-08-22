package unit29_struct

type Rectangle struct {
	width  int
	height int
}

func NewRectangle(width, height int) *Rectangle {
	return &Rectangle{width, height}
}

func rectangleScaleA(rect *Rectangle, factor int) {
	rect.width = rect.width * factor
	rect.height = rect.height * factor
}

func rectangleScaleB(rect Rectangle, factor int) {
	rect.width = rect.width * factor
	rect.height = rect.height * factor
}
