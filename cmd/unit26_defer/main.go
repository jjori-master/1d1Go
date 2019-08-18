package unit26_defer

func sayHello() func() string {
	var greeting string

	world := func() {
		greeting += " world!!"
	}

	hello := func() {
		defer world()
		greeting += "hello"
	}

	return func() string {
		hello()
		return greeting
	}
}
