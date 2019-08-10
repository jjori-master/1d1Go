package main

import "fmt"

func printSlice(slice []int) {
	fmt.Println("--------------------------------------")
	for _, value := range slice {
		fmt.Println(value)
	}

}

func main()  {

	var slice []int

	// make를 통해 공간을 확보 한다.
	// 슬라이스는 make를 통해 공간을 확보해야 값을 넣을 수 있다.
	slice = make([]int, 5)
	printSlice(slice)

	slice2 := []int{1, 2, 3, 4, 5} // 바로 값을 할당
	printSlice(slice2)


	// 슬라이스는 append 함수를 이용해서 값을 추가한다.
	slice3 := []int{1, 2, 3, 4, 5}
	slice3 = append(slice3, 6, 7, 8, 9)
	printSlice(slice3)

	// 슬라이스에 슬라이스를 추가하려면 추가할 슬라이스에 ...을 덧붙인다.
	slice4 := []int{1, 2, 3}
	slice5 := []int{4, 5, 6}

	slice4 = append(slice4, slice5...)

	printSlice(slice4)


}
