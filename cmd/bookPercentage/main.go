package main

import "fmt"

func printBootPercentage(a int, b int) {
	r := (a * 100) / b
	fmt.Println(a, "/", b, " ", r, "%")
}

func main() {
	printBootPercentage(116, 420)
}
