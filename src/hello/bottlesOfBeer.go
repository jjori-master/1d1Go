package main

import "fmt"

func main() {

	for bottlesCount := 99; bottlesCount >= 0 ; bottlesCount-- {

		switch  {
		case bottlesCount > 1:
			s := "bottles"
			fmt.Printf("%d %s of beer on the wail, %d %s of beer\n", bottlesCount, s, bottlesCount, s)

			if bottlesCount - 1 == 1 {
				s = "bottle"
			}

			fmt.Printf("Take one down, pass it around, %d %s of beer on the wall\n", bottlesCount - 1, s)

		case bottlesCount == 1:
			fmt.Println("1 bottle of beer on the wail, 1 bottle of beer")
			fmt.Println("Take one down, pass it around, 1 bottle of beer on the wall")

		default:
			fmt.Println("No more bottles of beer on the wall, No more bottles of beer.")
			fmt.Println("Go to the store and buy some more, 99 bottles of beer on the wall.")
		}
	}

}
