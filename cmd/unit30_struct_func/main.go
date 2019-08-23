package unit30_struct_func

type Rectangle struct {
	width  int
	height int
}

func (rect *Rectangle) area() int {
	return rect.width * rect.height
}

func (rect *Rectangle) scaleA(factor int) {
	rect.width = rect.width * factor
	rect.height = rect.height * factor
}

func (rect Rectangle) scaleB(factor int) {
	rect.width = rect.width * factor
	rect.height = rect.height * factor
}
