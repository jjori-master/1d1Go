package unit35_synchronization

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

type Hello struct {
	messages []string
}

func (hello *Hello) sayHello() {
	hello.messages = append(hello.messages, "hello world")
}

type Data struct {
	tag    string
	buffer []int
}

func usePool() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	pool := sync.Pool{
		New: func() interface{} {
			data := new(Data)

			data.tag = "new"

			data.buffer = make([]int, 10)

			return data
		},
	}

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get().(*Data)

			for index := range data.buffer {
				data.buffer[index] = rand.Intn(100) // 랜던값 지정
			}

			fmt.Println(data)

			data.tag = "used"

			pool.Put(data)
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get().(*Data)

			n := 0

			for index := range data.buffer {
				data.buffer[index] = n

				n += 2
			}

			fmt.Println(data)

			data.tag = "used"

			pool.Put(data)
		}()
	}
}
