package unit34_channel

import "fmt"

func sum(a int, b int, c chan int) {
	c <- a + b
}

func sumReturnChan(a, b int) <-chan int {
	out := make(chan int)

	go func() {
		out <- a + b
	}()

	return out
}

// 보내기 전용 채널
func producer(c chan<- int) {

	for i := 0; i < 5; i++ {
		c <- i
	}

	c <- 100
}

// 받기 전용 채널
func consumer(c <-chan int) {
	for i := range c {
		fmt.Println(i)
	}

	fmt.Println(<-c)
}

// 보내기 전용 채널 num 함수
func num(a int, b int) <-chan int {
	out := make(chan int)

	go func() {
		out <- a
		out <- b
		close(out)
	}()

	return out
}

// 받기 전용 함수를 이용하여 채널을 보내는 sum 함수
func sumReciveChanelReturnChanel(c <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		r := 0

		for i := range c {
			r += i
		}

		out <- r
	}()

	return out
}
