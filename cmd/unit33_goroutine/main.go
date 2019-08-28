package unit33_goroutine

type Rectangle struct {
	width  int
	height int
}

func change(rect *Rectangle) {
	rect.width++
}
