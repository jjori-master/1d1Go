package unit34_channel

func sum(a int, b int, c chan int) {
	c <- a + b
}
