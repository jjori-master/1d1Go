package unit27_panic

import "fmt"

func ExamplePanic() {
	defer func() {
		r := recover()

		fmt.Println(r)
		// Output:
		// asdf
	}()

	panic("Error !!!")
}

func panic1() {
	defer func() {
		r := recover()
		fmt.Println(r)

		panic("panic!!!")
	}()
	a := [...]int{1, 2}
	for i := 0; i < 3; i++ {
		fmt.Println(a[i])
	}
}
