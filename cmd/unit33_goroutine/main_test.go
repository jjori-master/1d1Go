package unit33_goroutine

import (
	"runtime"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test functions Suite")
}

var _ = Describe("Unit 33 goroutine", func() {
	Context("goroutine 실습", func() {
		It("goroutine 으로 10번 호출 그런데 뭔가가  이상한 테스트얌... ", func() {
			rect := Rectangle{}

			for i := 0; i < 10; i++ {
				go change(&rect)
			}

			time.Sleep(time.Duration(1000))

			Expect(rect.width).Should(Equal(10))
		})

		It("goroutine 일반 클로저로 실행", func() {
			runtime.GOMAXPROCS(1)

			n := 0

			for i := 0; i < 10; i++ {
				go func() {
					n += i
				}()
			}

			time.Sleep(time.Duration(1000))

			Expect(n).Should(Equal(100))
		})

		It("goroutine 클로저를 매개 변수로 실행", func() {
			runtime.GOMAXPROCS(1)

			n := 0

			for i := 0; i < 10; i++ {
				go func(x int) {
					n += x
				}(i)
			}

			time.Sleep(time.Duration(1000))

			Expect(n).Should(Equal(45))
		})
	})
})
