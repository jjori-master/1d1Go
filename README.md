# 1d1Go

> 하루 30분 [가장 빨리 만나는 Go(이재홍)](http://pyrasis.com/go.html) 책을 읽는다.
>
> 책을 읽은 후 10분 정도 가볍게 정리를 한다.
>
> 예제가 있으면 우선 풀어 본다.



- Go 환경 설정

  > go 1.11버전 이후부터는 `dep`이 아니라 `mod`를 사용한다.

  ```bash
  $ export GO111MODULE=on
  
  $ go mod init
  $ go mod download
  
  사용 안하는 패키지 정리 및 재 다운로드
  $ go mod tidy
  
  private repository 실패 시 로그인 요청
  
  $ export GIT_TERMINAL_PROMPT=1
  $ go mod download
  
  
  set GO111MODULE=on
  set GIT_TERMINAL_PROMPT=1
  go mod download
  ```




[2019-08-14 수]

- 오늘 회식이라 정신이 멍함... function에 대해 조금 풀어봄
  

[2019-08-13 화]

- 책 진도 못나감 ㅠㅠ
- 뭔가 go 환경이 박살났다~!!!
- 원인
  - test를 위해 `ginkgo`가 필요해 `go get`명령어를 통해 가져왔는데 왜 자꾸 `src`에 쌓이는 거여!!!
    `src/github.com` , `pkg` 폴더를 gitignore를 했더니 뭐야 왜 못찾아 오는거야 ㅠㅠ
    음 문제는 `GOPATH`에 있었다. `fmt`기능을 사용하기 위해 단순히 'import fmt`만 하면 되었다.
    그 `fmt 어딨는겨? 바로!! `GOROOT/src`에 있다. 그런데  3rd party 라이브러리는 어디에 저장될까
    바로!!! `GOPATH/src' 다. 그래서 `go get` 명령어로 라이브러리를 가져올때 내 프로젝트 src에 깔려버리네....ㅡㅡ;;
- 해결
  - 우선 못했음
  - 왜 덕철님이  projects 폴더 밑에 `pkg, bin, src`폴더를 만들고 src에 프로젝트를 만드는 이유가 
    있어서 그래도 하려고 했더니 안됨..... 막 터짐 ㅠㅠ 내 속창아지도 터짐..
    내일 동료분께 여쭤봐야 할것 같음 ㅠㅠ



[2019-08-12 월]

- 책 진도 (93/420, 22.1%)

- 맵

  - 맵 선언 

    - 맵은 선언하는 방식은 다른 타입과 동일하다. 맵은 make를 통해 공간 할당을 해야 사용가능하다.

      ```go
      var a map[string]int // 맵 선언
      a = make(map[string]int) // make를 통해 공간을 할당해야지만 사용가능하다.
      
      var b = make(map[string]int) // make를 통해 타입을 선언하지 않고 맵 생성
      
      c := make(map[string]int) // var 생략
      
      d := map[string]int{"Math": 100, "English": 50} // 초기값으로 맵 생성
      
      f := map[string]int{
          "Math": 100,
          "English": 50, // 마지막은 ,를 붙인다.
      }
      ```

  - 맵 순회

    - 맵 순회는 range를 활용한다.

      ```go
      score := map[string]int{
          "Math": 100,
          "English": 50,
          "Korean": 80,
      }
      
      totalScore := 0
      
      // key나 valuefmf 사용하지 않을 경우 `_` 로 대입한다.
      for key, value := range score { 
          fmt.Println(key, " : ", value)
          totalScore += value
      }
      
      fmt.Println(totalScore) // 230
      ```

[2019-08-11 일]

- 책 진도 (89/420, 21.1%)

- 슬라이스

  - 슬라이스 복사

    - 슬라이스는 레퍼런스 타입이라 슬라이스를 단순 대입하게 되면 같은 레퍼런스를 가르키는
      변수가 된다.

      ```go
      slice1 := []int{1, 2, 3, 4, 5}
      var slice2 = slice1
      slice1[0] = 9
      
      printSlice(slice1) // [9, 2, 3, 4, 5]
      printSlice(slice2) // [9, 2, 3, 4, 5]
      ```

    - 값 복사를 하기 위해서는 `copy` 함수 를 사용한다.

      ```go
      slice1 := []int{1, 2, 3, 4, 5}
      slice2 := make([]int, 3)
      
      copy(slice2, slice1)
      
      slice1[0] = 9
      
      printSlice(slice1) // [9, 2, 3, 4, 5]
      printSlice(slice2) // [1, 2, 3], 길이가 3이라 복사가 길이 3밖에 되지 않음
      ```

  - 부분 슬라이스

    - 일단 예제, 예제를 보자

      ```go
      slice1 := []int{1, 2, 3, 4, 5}
      slice2 := slice1[0:2] // index 0부터 index 2 - 1 (0~1)까지 부분을 가져온다.
      
      printSlice(slice1) // [1, 2, 3, 4, 5]
      printSlice(slice2) // [1, 2]
      
      ```

    - index 앞과 뒤를 생략 할 수 있다.

      [:] 슬라이스의 모든 걸 가져온다.
      [:6] 0부터 6-1까지 index를 가져온다.
      [0:] 0부터 모든 정보를 가져9온다.

      

    - 부분 슬라이스는 레퍼런스라 부분 슬라이스를 수정하면, 원본도 바뀐다.

      ```go
      slice1 := []int{1, 2, 3, 4, 5}
      slice2 := slice1[0:2]
      
      slice2[0] = 99
      
      printSlice(slice1) // [99, 2, 3, 4, 5]
      printSlice(slice2) // [99, 2]
      ```

    - 배열도 동일하게 가능~

[2019-08-10 토]

- 책 진도 (110 / 528, 20%)

- 배열과 슬라이스

  - 일반적인 배열 선언

    ```go
    var arr [5]int // 배열 갯수를 선언 모두 초기값 0이 됨
    arr = [5]int{1, 2, 3, 4, 5}
    var arr2 = [5]int{1, 2, 3, 4, 5}
    arr3 := [5]int{1, 2, 3, 4, 5}
    ```

  - 배열의 순회 타입 첫번째 `len` 을 사용

    ```go
    arr := [5]int{1, 2, 3, 4, 5}
    for i := 0; i < len(arr); i++ {
      fmt.Println(i)
    }
    ```

  - 배열의 순회 타입 두번째 `range`를 사용

    ```go
    arr := [5]int{1, 2, 3, 4, 5}
    for i, value := range(arr) {
      fmt.Println(i, " : ", value)
    }
    0 : 1
    1 : 2
    2 : 3
    3 : 4
    4 : 5  출력
    
    // i는 인덱스 value는 배열의 값이 할당
    // 인덱스를 사용하지 않으려면 _ 를 i변수 대신 사용한다.
    
    for _, value := range(arr) {
      fmt.Println(value)
    }
    1
    2
    3
    4
    5 출력
    ```

  - 배열은 `레퍼런스 타입` 아니므로 다른 변수에 할당할때 `값 복사`만 된다.

    ```go
    arr := [2]int{1, 2}
    arr2 := arr
    
    arr[1] = 7
    
    for _, value := range(arr) {
      fmt.Println(value)
    }
    1
    7 출력
    
    for _, value := range(arr2) {
      fmt.Println(value)
    }
    1
    2 출력
    ```

- 슬라이스

  - 슬라이스는 배열과 같지만 길이가 고정되어 있지 않고, 동적으로 크기가 늘어나는 특징이 있다.
    슬라이스는 `레퍼런스 타입` 이다

  - 슬라이스 변수 선언 및 값 할당

    ```go
    var slice []int
    
    // make를 통해 공간을 확보 한다.
    // 슬라이스는 make를 통해 공간을 확보해야 값을 넣을 수 있다.
    slice = make([]int, 5)
    
    slice2 := []int{1, 2, 3, 4, 5} // 바로 값을 할당
    ```

  - 슬라이스는 `make` 를 통해 길이와 용량을 확보한다. 미리 용량을 확보하면, 길이에 벗어나는 값을 추가시 추가 용량 확보가 필요 하지 않아 처리 속도가 높다. 하지만 미리 공간을 잡고 있어야 하는 만큼 메모리 낭비가 있을 수 있다.
    용량을 미리 설정 하지 않게 되면, 추가될때마다 공간을 확보해야 하므로 처리 속도가 떨어진다.
    어쨎든 잘 생각해서 할당하자~~

  - 슬라이스 값 추가

    ```go
    // 슬라이스는 append 함수를 이용해서 값을 추가한다.
    slice := []int{1, 2, 3, 4, 5}
    slice = append(slice, 6, 7, 8, 9)
    
    // 슬라이스에 슬라이스를 추가하려면 추가할 슬라이스에 ...을 덧붙인다.
    slice1 := []int{1, 2, 3}
    slice2 := []int{4, 5, 6}
    
    slice1 = append(slice1, slice2...)
    ```

    

  -  



[2019-08-09 금] 

공부 정리

- switch case문
  - Go의 swtich case문도 다른 언어들과 크게 다를게 없음 끗~

[2019-08-08 목] 

공부 정리

- 문자열 다루기

  - 모든 문자열은 utf-8이다.
  - 문자열 길이를 재는 함수는 len인데 그 이유는
    **“한글”**을 UTF-8로 저장하면 `0xed, 0x95, 0x9c, 0xea, 0xb8, 0x80`가 되기 때문
  - 그래서 `unicode/utf8 패키지의 RuneCountInString 함수를 사용`하여 길이를 구해야 한다.

- if

  - if 는 무조건 bool이 와야한다. javascript는 특정 조건 `null, undefiend 등`이 와도 괜찮은데
    여기는 얄짤없이 에러 팍!!

  - 아래 코드처럼 if 절에 변수를 대입하여 바로 사용가능하다

    ```go
    if b, err = ioutil.ReadFile("./hello.txt"); err == nil {
    		fmt.Printf("%s", b)
    }
    ```

    단 이렇게 선언하고 if 바깥 구문에서는 `b` 변수를 사용 할 수 없다.

- for

  - Label를 사용하는거 그건 신기

[2019-08-07] 

공부 정리

- 변수 할당에 관하여

  - 변수의 선언과 할당은 기존의 내가 배워온 방식과는 사뭇 다르다

  - 변수의 선언은 `var`와 타입을 반드시 명해줘야 한다.

    ```go
    var i int
    var j // Compile error
    ```

  - 변수의 값 할당은 `타입`을 선언하지 않고 가능하다. 단 `var`는 반드시 있어야 한다.

    ```go
    var i = 10
    j = 20 // compile error
    ```

  - 변수의 선언과 값 할당을 바로 해줄 수 있다.

    ```go
    i := 10
    ```

  - 멀티로도 선언과 할당이 가능하다

    ```go
    var i, j = 10, 11
    	_ = i
    	_ = j
    ```

  - 위의 코드를 보면 `_`언더스코어로 i와 j를 할당했는데
    왜냐하면 변수를 선언하고 사용하지 않으면 에러가 나기 때문

- `부동 소수점 반올림 오차`와  `epsilon` 

  > 실수를 `==` 등호로 비교하면 안된다.
  >
  > 왜냐하면 .... 사실 봐도 모르겠음. 내일 자료를 더 찾아 보던지 
  >
  > 아니면 덕철님께 여쭤 봐야지~~



[2019-08-06] 

> 공부 정리
>
> - Go의 특성에 관하여
>
>   - 컴파일 언어이면서 강타입 언어
>
>     > 자바는 약타입 언어이다. 약타입 언어는 타입이 자동 형 변환된다.
>     >
>     > ```java
>     > String str = "hello";
>     > int number = 1;
>     > String result = str + number; // result is hello1
>     > ```
>     >
>     > 자바는 문자형과 숫자형을 `+` 연산을 하게 되면, 문자형으로 자동 형변환을 한다.
>     >
>     > 해당 코드는 문제가 되지 않지만 Go 언어에서는 `강타입`이므로 형변환이 되지 않아 에러가 난다.
>
>   - 가비지 컬렉션을 지원한다.
>
>     - 자바, 루비등의 가상머신에서 가비지 컬렉션을 지원하는 것과는 다르게 언어 차원에서 지원한다.
>     - 메모리를 직접 다루는 운영체제 프로그램을 만들기에는 부적합 하다고 한다. (절대적이진 않다.)
>     - 대신 다양한 네트워크 라이브러리가 지원되므로 인터넷 프로그래밍에 유용하다. (절대적이진 않다.)
>
> - Go 설치 및 환경 설정
>
>   > 현재 내 환경은 windows 10며 scoop로 Go를 설치한다.
>   >
>   > 하지만 일반적으로  `scoop install go`명령어로만 한다면 `go` 설치된 곳을 윈도우는 알지 못한다.
>   >
>   > ```bash
>   > > setx scoopApps "C:\Users\%username%\scoop\apps" /M
>   > > echo %scoopApps%
>   > > scoop install go 
>   > ```
>   >
>   > 위와 같은 명령어를 통해 환경변수 설정 이후 scoop을 통해 go를 설치 한다면 cmd에서 바로 go 명령어를 알아 들을 수 있다.
>   >
>   > go가 설치된 path를 알고 싶다면 `go env` 명령어를 통해 알수 있다.