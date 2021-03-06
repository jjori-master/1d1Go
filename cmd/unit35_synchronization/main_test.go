package unit35_synchronization

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
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

			wg := new(sync.WaitGroup)

			for i := 0; i < 100; i++ {
				wg.Add(1)
				go func() {
					rwMutex.Lock()
					data += 1
					rwMutex.Unlock()
					wg.Done()
				}()
			}

			for i := 0; i < 100; i++ {
				wg.Add(1)
				go func() {
					rwMutex.Lock()
					data += 1
					rwMutex.Unlock()
					wg.Done()
				}()
			}

			wg.Wait()
			Expect(data).Should(Equal(200))
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

		It("풀 사용", func() {
			usePool()

			time.Sleep(3 * time.Second)
		})
	})

	Context("대기 그룹 사용", func() {
		It("기본적인 대기 그룹 사용 고루틴이 모두 끝날때가지 대기", func() {

			var data int = 0

			wg := new(sync.WaitGroup)

			for i := 0; i < 10; i++ {
				wg.Add(1)
				go func(n int) {
					if n > data {
						data = n
					}
					wg.Done()
				}(i)
			}

			wg.Wait()

			Expect(data).Should(Equal(9))
		})
	})

	Context("원자적 연산", func() {

		It("2000더하고 1000빼기", func() {
			var data int64 = 0
			wg := new(sync.WaitGroup)

			// 2000번 더하기
			for i := 0; i < 2000; i++ {
				wg.Add(1)

				go func() {
					atomic.AddInt64(&data, 1)
					wg.Done()
				}()
			}

			for i := 0; i < 1000; i++ {
				wg.Add(1)

				go func() {
					atomic.AddInt64(&data, -1)
					wg.Done()
				}()
			}

			wg.Wait()

			var expectData int64 = 1000
			Expect(data).Should(Equal(expectData))
		})
	})
})
