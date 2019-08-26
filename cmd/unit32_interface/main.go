package unit32_interface

type hello interface {
}

type Rectangle struct {
	width  int
	height int
}

func (rect *Rectangle) area() int {
	return rect.width * rect.height
}

type Triangle struct {
	width  int
	height int
}

func (rect *Triangle) area() int {
	return (rect.width * rect.height) / 2
}

type AreaCalculator interface {
	area() int
}

type Duck struct {
}

func (d Duck) quack() string {
	return "quack"
}

type Person struct{}

func (p Person) quack() string {
	return "꽥"
}

type Quacker interface {
	quack() string
}

func vocalCord(q Quacker) string {
	return q.quack()
}
