#### UNIT 35  동기화 객체 사용하기



-  Mutex

  > Mutex는 여러 고루틴이 공유하는 데이터를 보호할때 사용

  ```go
  // 데이터가 보호되지 않는 상황
  
  runtime.GOMAXPROCS(runtime.NumCPU()) // CPU 맘껏 사용
  
  var data []int
  
  go func() {
      for i := 0; i < 1000; i++ {
          data = append(data, 1)
          
          runtime.Gosched()
      }
  }()
  
  go func() {
      for i := 0; i < 1000; i++ {
          data = append(data, 1)
  
          runtime.Gosched()
      }
  }()
  
  time.Sleep(2 * time.Second)
  
  fmt.Println(len(data)) // 무작위 1800 ~ 1990 값예상
  ```

  

  > 두개의 고루틴에서 하나의 데이터를 두고 경합이 발생되어 동시에 접근하기 때문에 이런
  >
  > 현상이 발생

   ```go
  // mutex를 사용하여 동시에 접근하는것을 방지
  
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
   ```

  



- 읽기, 쓰기 Mutex

  > Mutex는 읽기 동작 전용, 쓰기 동작 전용으로 락을 걸 수 있다.
  >
  > `읽기 전용 락`은 서로 락을 잡지 않지만 읽는 도중 데이터의 변조가 있으면 안되기 때문에
  > `쓰기 전용 락`은 막는다.
  >
  > `쓰기 전용 락`은 모든 락을 막는다.

  ```go
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
          rwMutex.RLock() // 읽기 전용 락끼리는 락을 걸리 않는다.
          
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
  ```


- 조건 변수 사용하기

  > 조건 변수는 대기하고 있는 객체 하나만 깨우거나 여러개를 깨울때 사용

  - 조건 한개씩 깨우기

    ```GO
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
    
    // 결과값 
    
    Wait begin :  3
    Wait begin :  1
    Wait begin :  2
    
    signal :  0
    signal :  1
    signal :  2
    
    Wait end :  3
    Wait end :  1
    Wait end :  2
    ```




- 함수 한번만 실행하기

  > 고루틴을 여러번 호출하나 고루틴 안의 실행 함수는 단 한번만 실행

  ```go
  type Hello struct {
  	messages []string
  }
  
  func (hello *Hello) sayHello()  {
  	hello.messages = append(hello.messages, "hello world")
  }
  
  runtime.GOMAXPROCS(runtime.NumCPU())
  
  once := new(sync.Once)
  
  var hello *Hello   // 구조체 포인터 선언
  hello = new(Hello) // 구조체 메모리 할당
  
  for i := 0; i < 3; i++ {
      go func() {
          once.Do(hello.sayHello) // 함수 단 한번만 실행
      }()
  }
  
  time.Sleep(1 * time.Second)
  
  Expect(len(hello.messages)).Should(Equal(1))
  ```

  

- 풀 사용하기

  > 풀은 객체(메모리)를 사용한 후 보관해 두었다가 다시 사용하게 해주는 기능
  > 즉 새로운 객체를 메모리에 할당할 필요 없이 풀에서 가져다 사용하는 방법으로
  > 메모리 생성, 해제의 가비지 컬렉터에게 부담을 줄여주는 방법. 일종의 캐시

  ```go
  type Data struct {
  	tag string
  	buffer []int
  }
  
  func main()  {
  	runtime.GOMAXPROCS(runtime.NumCPU())
  
  	pool := sync.Pool{
  		New: func() interface{} {
  			data := new(Data)
  
  			data.tag = "new"
  
  			data.buffer = make([]int, 10)
  
  			return data
  		},
  	}
  
  	for i :=0; i< 10; i++ {
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
  
  	for i :=0; i< 10; i++ {
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
  
  출력값은 그때마다 다르다. pool에서 객체를 가져올때 없으면 tag에 new를 하기로 했음
  출력에는 2개의 new가 있으며 나머지는 pool에서 있는 객체를 그대로 이용.
  
  출력 : 
  &{new [0 2 4 6 8 10 12 14 16 18]}
  &{used [94 11 62 89 28 74 11 45 37 6]}
  &{used [95 66 28 58 47 47 87 88 90 15]}
  &{used [41 8 87 31 29 56 37 31 85 26]}
  &{used [13 90 94 63 33 47 78 24 59 53]}
  &{used [57 21 89 99 0 5 88 38 3 55]}
  &{used [51 10 5 56 66 28 61 2 83 46]}
  &{used [63 76 2 18 47 94 77 63 96 20]}
  &{used [23 53 37 33 41 59 33 43 91 2]}
  &{used [78 36 46 7 40 3 52 43 5 98]}
  &{used [0 2 4 6 8 10 12 14 16 18]}
  &{used [0 2 4 6 8 10 12 14 16 18]}
  &{used [0 2 4 6 8 10 12 14 16 18]}
  &{used [0 2 4 6 8 10 12 14 16 18]}
  &{used [0 2 4 6 8 10 12 14 16 18]}
  &{used [0 2 4 6 8 10 12 14 16 18]}
  &{used [0 2 4 6 8 10 12 14 16 18]}
  &{new [81 87 47 59 81 18 25 40 56 0]}
  &{used [0 2 4 6 8 10 12 14 16 18]}
  &{used [0 2 4 6 8 10 12 14 16 18]}
  ```

  