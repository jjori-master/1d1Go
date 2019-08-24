#### 구조체 임베딩 사용하기



- 구조체 안에 구조체를 인자로 지정해서 사용한다. 전형적인 Has a 관계

  ```go
  type Person struct {
  	name string
  	age  int
  }
  
  func (p *Person) greeting() string {
  	return "Say Hello!!"
  }
  
  type Students struct {
  	p      Person
  	school string
  	grade  int
  }
  
  s := Students{}
  
  fmt.Println(s.p.gretting()) // Say Hello!!
  ```

- 구조체를 구조체 안에 변수로 선언하지 않고 그냥 작성한다. Is a 관계

  ```go
  type Person struct {
  	name string
  	age  int
  }
  
  func (p *Person) greeting() string {
  	return "Say Hello!!"
  }
  
  type Students struct {
  	Person // 변수가 아닌 구조체를 인자로 넣음
  	school string
  	grade  int
  }
  
  s := Students{}
  
  fmt.Println(s.gretting()) // Say Hello!!
  ```

- 만약 구조체를 임베딩한 구조체가 부모 구조체와 이름이 동일한 함수를 가지고 있다면, 오버라이딩 된다.

  ```go
  type Person struct {
  	name string
  	age  int
  }
  
  func (p *Person) greeting() string {
  	return "Say Hello!!"
  }
  
  type Students struct {
  	Person // 변수가 아닌 구조체를 인자로 넣음
  	school string
  	grade  int
  }
  
  func (p *Students) greeting() string {
  	return "BAAAM!!"
  }
  
  s := Students{}
  
  fmt.Println(s.gretting()) // BAAAM!!
  ```

  