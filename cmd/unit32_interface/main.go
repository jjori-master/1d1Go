package unit32_interface

import "strconv"

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

type Person struct {
	name string
	age  int
}

func (p Person) quack() string {
	return "꽥"
}

type Quacker interface {
	quack() string
}

func vocalCord(q Quacker) string {
	return q.quack()
}

func formatString(arg interface{}) string {

	switch arg.(type) {
	case int:
		i := arg.(int)
		return strconv.Itoa(i)

	case float32:
		f := arg.(float32)
		return strconv.FormatFloat(float64(f), 'f', -1, 32)

	case float64:
		f := arg.(float64)
		return strconv.FormatFloat(f, 'f', -1, 64)

	case string:
		s := arg.(string)
		return s

	case Person:
		p := arg.(Person)
		return p.name + " " + strconv.Itoa(p.age)

	case *Person:
		p := arg.(*Person)
		return p.name + " " + strconv.Itoa(p.age) + " 1"

	default:
		return "Error"
	}
}
