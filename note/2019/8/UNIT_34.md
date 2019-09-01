#### UNIT 34 채널 사용하기

> 채널은 고루틴낄 데이터를 주고 받고,  실행 흐름을 제어하는 기능
>
> 모든 타입을 채널로 사용 가능
> 채널 자체는 값이 아닌 레퍼런스 타입



- 모든 타입을 채널로 사용가능

  ```go
  c := make(chan int) // int형 채널 생성
  
  var c2 chan int
  c2 = make(chan int) // 채널은 사용하기전 반드시 make 함수로 공간 할당
  ```



- 고루틴으로 실행한 함수는 값을 반환 받을 수 없다.
  채널을 통해 값을 받을 수 있다.

  ```go
  func sum(a int, b int, c chan int) {
  	c <- a + b // 채널에 a 와 b의 합산 결과를 내보냅
  }
  
  c := make(chan int) // int 형 채널 생성
  
  go sum(1, 2, c)
  
  n := <-c // 채널을 넘겼으면 받을 때까지 대기 즉 동기화
  
  fmt.Println(n) // 3
  ```



- 채널에 값을 집어 넣었을때 그 값을 사용하지 않으면 대기한다.

  ```go
  var slice []int
  slice = make([]int, 0)
  
  c := make(chan int)
  
  go func() {
      for i := 0; i < 3; i++ {
          c <- i // 채널의 값을 가져가기 까지 반복문 대기
      }
  }()
  
  for i := 10; i < 13; i++ {
      n := <- c // 값을 가져가랏!!
      slice = append(slice, n)
      slice = append(slice, i)
  }
  
  Expect(slice[0]).Should(Equal(0)) // 고루틴
  Expect(slice[1]).Should(Equal(10)) // 다른 반복문 
  Expect(slice[2]).Should(Equal(1)) // 고루틴
  Expect(slice[3]).Should(Equal(11)) // 다른 반복문
  Expect(slice[4]).Should(Equal(2)) // 고루틴
  Expect(slice[5]).Should(Equal(12)) // 다른 반복문
  ```

- 채널 버퍼 사용

  > 채널 생성시 두번째 인자에 숫자를 기입하면 그 사이즈 대로 버퍼가 사이즈가 된다.
  > 채널 버퍼를 1개 이상 설정시 비동기 채널이 생성이 된다. 

  

  ```go
  done := make(chan bool, 2) // 채널 버퍼 사이즈 2로 설정
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
  
  보냈어 : 0
  보냈어 : 1
  보냈어 : 2
  받았어 : 0
  받았어 : 1
  받았어 : 2
  받았어 : 3
  보냈어 : 3
  ```




- range 와 close

  > for문 안에서 range  키워드를 사용하면 채널이 닫힐 때까지 반복하면서 값을 꺼낸다.

  ```go
  c := make(chan int, 1)
  
  go func() {
    for i:=0;i<5;i++ {
      c <- i
  	}
  
    close(c)
  }()
  
  for i:=range c {
    fmt.Println(i)
  }
  
  
  // 0
  // 1
  // 2
  // 3
  // 4
  
  ```

  

  > 채널이 닫혀 있는지 여부를 확인 할수 있다.

  ```go
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
  ```
  


- 보내기 전용, 읽기 전용 채널

  > 읽기만 전용할 수 있고 받기만 전용할수 있는 함수를 만든다.
  > 방법은 간단한데 채널을 인자로 받을때 `chan`이 어디에 위치하느냐의 차이
  > func(c chan<- int) //  보내기 전용 채널
  > func(c <-chan int) // 받기 전용
  >
  > 보내기 전용은 보내기만 가능하며, 읽기 전용은 읽기만 가능

  ```go
  // 보내기 전용 채널
  func producer(c chan<- int) {
  
  	for i := 0; i < 5; i++ {
  		c <- i
  	}
  
  	c <- 100
      
      // <- c 보내기 전용이라 받을 수 없음 에러!!
      
  }
  
  // 받기 전용 채널
  func consumer(c <-chan int) {
  	for i := range c {
  		fmt.Println(i)
  	}
  
  	fmt.Println(<-c)
      
      // c <- 1 받기 전용이라 보낼 수 없음 에러!!
  }
  
  ```



- 채널 리턴

  > 채널도 인자로 리턴 받을 수 있다.
  > 그리리고 그 채널은 받기 전까지는 둥둥 떠다닐뿐~~ 아니 클로저로 그 채널을 잡고 받으면 문제가 없으나
  > 만약 함수내에서 생성한 채널이고 채널로 데이터를 보냈다면, 그 채널이 받기 전까지는 둥둥 떠다닌다.
  > 아니 대기중~!!

  ```go
  func sumReturnChan(a, b int) <-chan int {
  	out := make(chan int)
  
  	go func() {
  		out <- a + b
  	}()
  
  	return out
  }
  
  
  Expect(<-c).Should(Equal(3))
  ```

  