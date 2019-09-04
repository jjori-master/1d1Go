#### UNIT 30 구조체에 메서드 연결하기

- Go언어에는 클래스가 없는 대신 구조체에 함수를 연결 할 수 있다.
  참고로 클래스, 구조체, 열거형에 연결된 함수를 메서드라 부른다.

  메소드를 선언할때 `리시버`변수라 불리는 구조체 변수를 선언하는데
  `포인터`형과 `일반`형으로 구분한다.

  - 포인터 리시버 변수를 사용한 일반적인 구조체 메서드 연결 

    ```go
    type Rectangle struct {
    	width  int
    	height int
    }
    
    func (rect *Rectangle) area() int {
    	return rect.width * rect.height
    }
    
    rect := Rectangle{10, 20}
    area := rect.area()
    
    fmt.Println(area) // 200
    ```

  - 포인터 리시버 변수를 사용하고 값을 변경하면 선언된 구조체 값이 변경된다.

    ```go
    type Rectangle struct {
    	width  int
    	height int
    }
    
    func (rect *Rectangle) scaleA(factor int) {
    	rect.width = rect.width * factor
    	rect.height = rect.height * factor
    }
    
    rect := Rectangle{10, 20}
    rect.scaleA(10)
    
    fmt.Println(rect) // {100, 200}
    ```

  - 값(밸류) 리시버 변수를 사용하고 값을 변경하면 선언된 구조체는 값이 변경되지 않는다.
    왜냐하면 그냥 값만 복사되었기 때문이다.

    ```go
    type Rectangle struct {
    	width  int
    	height int
    }
    
    func (rect Rectangle) scaleA(factor int) {
    	rect.width = rect.width * factor
    	rect.height = rect.height * factor
    }
    
    rect := Rectangle{10, 20}
    rect.scaleA(10)
    
    fmt.Println(rect) // {10, 20}
    ```

    

