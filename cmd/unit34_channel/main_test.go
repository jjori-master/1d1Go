package unit34_channel

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"time"
)

func TestGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test functions Suite")
}

var _ = Describe("Unit 34 channel", func() {
	Context("channel 실습", func() {
		It("channel 로 합산 데이터 받기", func() {
			c := make(chan int) // int 형 채널 생성

			go sum(1, 2, c)

			n := <-c

			Expect(n).Should(Equal(3))
		})

		It("채널을 := 대신 var로 선언해서 사용하기", func() {
			var c chan int
			c = make(chan int)

			go sum(12, 7, c)

			n := <-c

			Expect(n).Should(Equal(19))
		})

		It("채널 동기화 테스트", func() {

			var slice []int
			slice = make([]int, 0)

			c := make(chan int)

			go func() {
				for i := 0; i < 3; i++ {
					c <- i
				}
			}()

			for i := 10; i < 13; i++ {
				n := <-c
				slice = append(slice, n)
				slice = append(slice, i)
			}

			Expect(slice[0]).Should(Equal(0))
			Expect(slice[1]).Should(Equal(10))
			Expect(slice[2]).Should(Equal(1))
			Expect(slice[3]).Should(Equal(11))
			Expect(slice[4]).Should(Equal(2))
			Expect(slice[5]).Should(Equal(12))
		})

		It("채널 버퍼 테스트", func() {
			done := make(chan bool, 2)
			count := 4

			go func() {
				for i := 0; i < count; i++ {
					done <- true
					fmt.Println("보냈어 :", i)
				}
			}()

			for i := 0; i < count; i++ {
				<-done
				fmt.Println("받았어 :", i)
			}
		})

		It("채널이 닫혀 있는지 확인", func() {
			c := make(chan int)

			go func() {
				c <- 1
				close(c) // 채널 닫음
			}()

			n, ok := <-c
			Expect(n).Should(Equal(1))
			Expect(ok).Should(Equal(true))

			n, ok = <-c

			Expect(n).Should(Equal(0))
			Expect(ok).Should(Equal(false))
		})

		It("보내기 채널 받기 채널", func() {
			c := make(chan int)

			go producer(c)

			go consumer(c)
		})

		It("채널을 인자로 받기", func() {

			c := sumReturnChan(1, 2)

			Expect(<-c).Should(Equal(3))
		})

		It("채널을 인자로 받아 결과 값을 다시 채널로 보내는 sum 함수 테스트", func() {
			c := num(1, 2)

			out := sumReciveChanelReturnChanel(c)

			Expect(<-out).Should(Equal(3))
		})

		It("채널에 값을 넣고 close하고 다른 곳에서 값을 받을 수 있나?", func() {
			out := make(chan int)

			go func() {
				out <- 1

				close(out)
			}()

			for r := range out {
				Expect(r).Should(Equal(1))
			}
		})

		It("채널을 colse 하지 않으면 range에서 무한 대기!!", func() {
			out := make(chan int)

			go func() {
				out <- 1

				// close(out) 무한대기
			}()

			go func() {
				time.Sleep(2 * time.Second)
				close(out)
			}()

			for r := range out {
				Expect(r).Should(Equal(1))
			}
		})

		It("간단한 select 사용하기 ", func() {
			c1 := make(chan int)
			c2 := make(chan string)

			go func() {
				c1 <- 1
			}()

			go func() {
				c2 <- "안녕하세요"
			}()

			go func() {
				for {
					select {
					case i := <-c1:
						Expect(i).Should(Equal(1))

					case s := <-c2:
						Expect(s).Should(Equal("안녕하세요"))
					}
				}
			}()
		})

		It("select에서 채널에 값을 보내기", func() {
			c1 := make(chan int)

			go func() {
				i := <-c1
				Expect(i).Should(Equal(1))
			}()

			go func() {
				select {
				case c1 <- 1:
				}
			}()
		})

		It("select 에서 하나의 채널로 주고 받기", func() {
			c1 := make(chan int)
			i := 0
			var arr []int
			arr = append(arr, []int{1, 2, 3}...)

			go func() {
				for ; i < 3; i++ {
					select {
					case c1 <- arr[i]:
					case j := <-c1:
						Expect(j).Should(Equal(arr[i]))
					}
				}
			}()
		})
	})
})
