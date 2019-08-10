package main

import "fmt"

func printArray(arr [5]int) {
	fmt.Println("--------------------------------------")
	for _, value := range arr {
		fmt.Println(value)
	}

}

func main() {

	var arr [5]int
	printArray(arr)

	arr = [5]int{1, 2, 3, 4, 5}
	printArray(arr)

	var arr2 = [5]int{1, 2, 3, 4, 5}
	printArray(arr2)

	arr3 := [5]int{6, 7, 8, 9, 10}
	printArray(arr3)


	arr4 := [5]int{1, 2, 3, 4, 5}
	arr5 := arr4

	arr4[1] = 7

	printArray(arr4)
	printArray(arr5)

}
