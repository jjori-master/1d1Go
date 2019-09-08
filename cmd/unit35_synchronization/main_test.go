package unit35_synchronization

import (
	"fmt"
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

		It("읽기 쓰기 Mutex", func() {
			runtime.GOMAXPROCS(runtime.NumCPU())

			data := 0
			rwMutex := new(sync.RWMutex)

			go func() {
				for i := 0; i < 3; i++ {
					rwMutex.Lock()

					data = i

					rwMutex.Unlock()
				}
			}()

			go func() {
				for i := 0; i < 3; i++ {
					rwMutex.RLock()

					Expect(data).Should(Equal(i))

					rwMutex.RUnlock()
				}
			}()

			go func() {
				for i := 0; i < 3; i++ {
					rwMutex.RLock()
					Expect(data).Should(Equal(i))
					rwMutex.RUnlock()
				}
			}()
		})

		It("조건 변수 하나씩 깨우기", func() {
			runtime.GOMAXPROCS(runtime.NumCPU())

			var mutex = new(sync.Mutex)

			var cond = sync.NewCond(mutex)

			c := make(chan bool, 3)

			slice := []int{1, 2, 3}

			for _, s := range slice {
				go func(n int) {
					mutex.Lock()

					c <- true

					fmt.Println("Wait begin : ", n)

					cond.Wait()

					fmt.Println("Wait end : ", n)

					mutex.Unlock()
				}(s)
			}

			for i := 0; i < 3; i++ {
				<-c
			}

			for i := 0; i < 3; i++ {
				mutex.Lock()

				fmt.Println("signal : ", i)

				cond.Signal()

				mutex.Unlock()
			}
		})

		It("조건 변수 모두 깨우기", func() {
			runtime.GOMAXPROCS(runtime.NumCPU())

			var mutex = new(sync.Mutex)

			var cond = sync.NewCond(mutex)

			c := make(chan bool, 3)

			slice := []int{1, 2, 3}

			for _, s := range slice {
				go func(n int) {
					mutex.Lock()

					c <- true

					fmt.Println("Wait begin : ", n)

					cond.Wait()

					fmt.Println("Wait end : ", n)

					mutex.Unlock()
				}(s)
			}

			for i := 0; i < 3; i++ {
				<-c
			}

			mutex.Lock()

			fmt.Println("broadcast")

			cond.Broadcast()

			mutex.Unlock()
		})

		It("함수 한번만 실행하기", func() {
			runtime.GOMAXPROCS(runtime.NumCPU())

			once := new(sync.Once)

			var hello *Hello   // 구조체 포인터 선언
			hello = new(Hello) // 구조체 메모리 할당

			for i := 0; i < 3; i++ {
				go func() {
					once.Do(hello.sayHello)
				}()
			}

			time.Sleep(1 * time.Second)

			Expect(len(hello.messages)).Should(Equal(1))
		})

		FIt("풀 사용", func() {
			usePool()

			time.Sleep(3 * time.Second)
		})
	})
})
