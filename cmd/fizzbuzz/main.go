package main

import "fmt"

func main() {
	// 1에서 100까지 출력
	// 3의 배수는 Fizz 출력
	// 5의 배수는 Buzz 출력
	// 3과 5의 공배수는 FizzBuzz

	for i := 1; i <= 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz ", i)
			break
		case i%3 == 0:
			fmt.Println("Fizz ", i)
			break

		case i%5 == 0:
			fmt.Println("Buzz ", i)
			break
		default:
			fmt.Println(i)
		}
	}
}
