package main

import "fmt"

func printBootPercentage(a int, b int) {
	r := (a * 100) / b
	fmt.Println(r, "%")
}

func main() {
	printBootPercentage(134, 528)
}
