package unit34_channel

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
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
			}()

			n, ok := <-c
			Expect(n).Should(Equal(1))
			Expect(ok).Should(Equal(true))

			close(c) // 채널 닫음

			n, ok = <-c

			Expect(n).Should(Equal(0))
			Expect(ok).Should(Equal(false))
		})
	})
})
