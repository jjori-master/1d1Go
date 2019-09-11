#### Unit36 리플렉션 사용

> 리플렉션은 실행시점에 인터페이스나 구조체 등의 타입 정보를 얻어내거나 결정하는 기능



- 변수들의 타입을 얻어 낸다.

  ```go
  var num int = 1
  
  var expectType string
  expectType = fmt.Sprint(reflect.TypeOf(num))
  Expect(expectType).Should(Equal("int"))
  
  var s string = "Hello world!!"
  expectType = fmt.Sprint(reflect.TypeOf(s))
  Expect(expectType).Should(Equal("string"))
  
  var f float32 = 1.3
  expectType = fmt.Sprint(reflect.TypeOf(f))
  Expect(expectType).Should(Equal("float32"))
  
  var data Data = Data{1, 2}
  expectType = fmt.Sprint(reflect.TypeOf(data))
  Expect(expectType).Should(Equal("unit36_reflection.Data"))
  ```

  

- 변수의 타입 및 값의 상세 정보를 알수 있다.

  ```go
  var f float64 = 1.3
  
  t := reflect.TypeOf(f)
  v := reflect.ValueOf(f)
  
  Expect(t.Name()).Should(Equal("float64"))
  Expect(fmt.Sprint(t.Size())).Should(Equal("8"))
  
  Expect(fmt.Sprint(v.Type())).Should(Equal("float64"))
  
  Expect(v.Kind() == reflect.Float64).Should(Equal(true))
  Expect(v.Kind() == reflect.Int64).Should(Equal(false))
  
  Expect(v.Float()).Should(Equal(1.3))
  ```

  