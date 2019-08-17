package unit25_closure

func calc() func(x int) int {
	a, b := 3, 5

	return func(x int) int {
		return (a * x) + b
	}
}

func sayHelloTo() func(name string) string {
	base := "Hello "

	return func(name string) string {
		return base + name
	}
}
