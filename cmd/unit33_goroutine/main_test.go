package unit33_goroutine

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
)

func TestGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test functions Suite")
}

var _ = Describe("Unit 33 goroutine", func() {
	Context("goroutine 실습", func() {
		It("goroutine 으로 10번 호출 그런데 뭔가가  이상한 테스트얌... ", func() {
			rect := Rectangle{}

			wg := new(sync.WaitGroup)

			for i := 0; i < 10; i++ {
				wg.Add(1)
				go change(&rect, wg)
			}

			wg.Wait()

			Expect(rect.width).Should(Equal(10))
		})

		It("goroutine 일반 클로저로 실행", func() {
			runtime.GOMAXPROCS(runtime.NumCPU())

			var data int64 = 0
			wg := new(sync.WaitGroup)

			for i := 0; i < 10; i++ {
				wg.Add(1)
				go func() {
					atomic.AddInt64(&data, 1)
					wg.Done()
				}()
			}

			wg.Wait()

			Expect(data).Should(Equal(int64(10)))
		})

		It("goroutine 클로저를 매개 변수로 실행", func() {
			runtime.GOMAXPROCS(1)

			n := 0
			wg := new(sync.WaitGroup)

			for i := 0; i < 10; i++ {
				wg.Add(1)
				go func(x int) {
					n += x
					wg.Done()
				}(i)
			}

			wg.Wait()

			Expect(n).Should(Equal(45))
		})
	})
})
