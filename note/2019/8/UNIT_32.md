#### UNIT 32 Interface

> `Struct`가 필드의 집합이듯 `Interface` 는 함수의 집합이다.



- 인터페이스의 선언

  > 인터페이스의 기본 선언은 아래와 같이하며, 기본값은 nil이다.

  ```go
  type Printer interface {
  }
  
  var v Printer
  
  fmt.Println(p) // nill
  ```



- 인터페이스의 사용

  > 인터페이스는 함수의 집합이다. 여기에 선언된 함수가 어떤 구조체나 사용자정의 타입에서 구현한 메서드인지는 모른다. 아니 관심도 없다. 인터페이스는 자신이 선언한 함수와 동일한 구조의 함수를 구현하고 있는 구조체나 타입이라면
  >
  > 뭐든지 땡큐다. 

  ```go
  var myInt int
  
  func(i myInt) print() {
    fmt.Println(i)
  }
  
  type Printer interface {
    print()
  }
  
  var p Printer
  
  myInt := 5
  
  p = myInt
  
  fmt.Println(p.print()) // 5
  ```


- 왜 인터페이스를 사용하는가? (뇌피셜 주의)

  > 기본적으로 관심사의 분리이다.  나는 삼각형의 넓이를 구하고 싶다. 아니면 사각형의 넓이를 구하고 싶다는 특정 도형이라는 관심사가 붙어 있다. 여기서 삼각형이든 사각형이든 상관없이 도형의 넓이를 구하고 싶다면 어떻게 해야 하는가?
  >
  > 바로 인터페이스를 사용하여 관심사를 분리 하는 것이다.

  ```go
  type Rectangle struct {
  	width  int
  	height int
  }
  
  func (rect *Rectangle) area() int {
  	return rect.width * rect.height
  }
  
  type Triangle struct {
  	width  int
  	height int
  }
  
  func (rect *Triangle) area() int {
  	return (rect.width * rect.height) / 2
  }
  
  type Shape interface {
    area() int
  }
  
  
  func printShapeArea(shape Shape) {
    fmt.Println(shape.area())
  }
  
  var shape Shape
  
  rect := Rectangle{10, 20}
  triangle := Triangle{10, 20}
  
  shape = &rect
  printShapeArea(shape) // 200
  
  shape = &triangle
  printShapeArea(shape) // 100
  ```

  
  
  
  
  각 값이나 인스턴스의 실제 타입은 상관하지 않고, 구현된 메서드로만 타입을 판단하는
  방식을 `덕 타이핑(Duck typing)`이라 한다.
  
  > "만약 어떤 새가 오리처럼 걷고, 헤엄치고, 꽥꽥 거리는 소리는 낸다면 나는 그 새를
  > 오리라 부르겠다."
  
  아래 코드는 덕 타이핑을 구현한것
  
  ```go
  type Duck struct {
  }
  
  func (d Duck) quack() string {
  	return "quack"
  }

  type Person struct{}
  
  func (p Person) quack() string {
  	return "꽥"
  }
  
  type Quacker interface {
  	quack() string
  }
  
  func inTheForest(q Quacker) {
      fmt.Println(q.quack())
  }
  
  var donald Duck
  inTheForest(donald) // quack
  
  var jo Person
  inTheForest(jo) // 꽥
  ```
  
  `inTheForest`는 그냥 오리가 운다라는 것만 관심이 있을뿐
  그게 오리인지, 사람인지는 관심이 없다. 그냥 오리처럼 울수 있다면 `quack()`
  오리라고 판단한다.
  
  하지만 실제 숲에서 오리 흉내를 낸다면 총맞기 쉽지 않을까??
  
  



