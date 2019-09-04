#### UNIT 33 고루틴 Goroutine

> 고루틴은 함수를 동시에 실행시키는 기능
>
> 테스트를 어떻게 하징 ㅠㅠ



- 동시에 실행할 함수 앞에 `go`라고  붙여주면 끝

  ```go
  func hello() {
      fmt.Println('hello')
  }
  
  for i := 0; i < 10; i++ {
  	go hello();
  }
  
  // hello * 10
  ```

  

- 고루틴을 종료하고 싶다면 `return`  또는 `runtime.Goexit`함수를 사용
  단 return에 값이나 변수를 지정하더라도 무시



- 클로저를 고루틴으로 실행

  - 함수에서 클로저 변수를 그냥 사용하는 경우
    `고루틴으로 실행한 클로저는 반복문이 끝난 뒤에 고루틴이 실행`

    그래서 마지막에 변경된 i의 값 10이 계속 합산되어 100이라는 숫자가 생김
    원래 원했던 건 0 부터 9까지의 합 45가 나와야함

    ```go
    runtime.GOMAXPROCS(1)
    
    n := 0
    
    for i := 0; i < 10; i++ {
        go func() {
            n += i
        }()
    }
    
    time.Sleep(time.Duration(1000))
    
    Expect(n).Should(Equal(100))
    ```

  - 고루틴 함수를 실행할때 인자로 클로저 함수는 주는 경우
    클로저가 값복사되어 실행함으로 0 부터 9까지의 합인 45가 나올 수 있음

    ```go
    runtime.GOMAXPROCS(1)
    
    n := 0
    
    for i := 0; i < 10; i++ {
        go func(x int) {
            n += x
        }(x)
    }
    
    time.Sleep(time.Duration(1000))
    
    Expect(n).Should(Equal(45))
    ```

    

