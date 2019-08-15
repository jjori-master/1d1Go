package unit24_func

func sum(a int, b int) (r int) {
	r = a + b
	return
}

func sumNDiff(a int, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return
}

func sumAll(n ...int) int {
	total := 0

	for _, value := range n {
		total += value
	}

	return total
}
