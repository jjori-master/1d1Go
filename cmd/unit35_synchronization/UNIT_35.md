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

  