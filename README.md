# 1d1Go

> 하루 30분 [가장 빨리 만나는 Go(이재홍)](http://pyrasis.com/go.html) 책을 읽는다.
>
> 책을 읽은 후 10분 정도 가볍게 정리를 한다.
>
> 예제가 있으면 우선 풀어 본다.



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