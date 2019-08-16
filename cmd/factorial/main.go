package factorial

// n! 1부터 n까지의 곱샘을 한다.
func factorial(n uint64) uint64 {

	if n == 1 {
		return 1
	}

	return n * factorial(n-1)
}
