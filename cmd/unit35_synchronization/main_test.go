package unit35_synchronization

import (
	"runtime"
	"sync"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test functions Suite")
}

var _ = Describe("Unit 35 동기화 객체 사용", func() {
	Context("동기화 객체 실습", func() {
		It("Mutex를 사용하여 고루틴에서 공유하는 데이터 보호", func() {
			runtime.GOMAXPROCS(runtime.NumCPU()) // CPU 맘껏 사용

			var data []int

			var mutex = new(sync.Mutex)

			go func() {
				for i := 0; i < 1000; i++ {
					mutex.Lock()

					data = append(data, 1)

					mutex.Unlock()

					runtime.Gosched()
				}
			}()

			go func() {
				for i := 0; i < 1000; i++ {
					mutex.Lock()

					data = append(data, 1)

					mutex.Unlock()

					runtime.Gosched()
				}
			}()

			time.Sleep(2 * time.Second)

			Expect(len(data)).Should(Equal(2000))
		})
	})
})
