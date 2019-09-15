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




- 리플렉션은 구조체 필드(Field)의 `태그 정보` 를 가져온다.

  > 구조체의 필드는 태그를 지정할 수 있다.
  >
  > ```go
  > type Person struct {
  > 	name string `tag1:"이름"tag2:"Name"`
  > 	age int `tag1:"나이"tag2:"Age"`
  > }
  > ```

  - 구조체 필드 태그 정보를 가져올 수 있다.

    ```go
    p := Person{}
    
    name, ok := reflect.TypeOf(p).FieldByName("name")
    Expect(ok).To(BeTrue()) 
    Expect(name.Tag.Get("tag1")).Should(Equal("이름"))
    Expect(name.Tag.Get("tag2")).Should(Equal("Name"))
    
    age, ok := reflect.TypeOf(p).FieldByName("age")
    Expect(ok).To(BeTrue())
    Expect(age.Tag.Get("tag1")).Should(Equal("나이"))
    Expect(age.Tag.Get("tag2")).Should(Equal("Age"))
    ```




- 리플렉션을 이용해 일반 포인터와 인터페이스 정보를 확인한다.

  ```go
  // 포인터 정보 확인
  var a *int = new(int)
  *a = 1
  
  pType := fmt.Sprint(reflect.TypeOf(a))
  
  // 포인터의 실제 정보를 가져오기 위해서는 Elem를 사용하여야 합니다.
  pValueType := fmt.Sprint(reflect.ValueOf(a).Elem())
  // pValue := refect.ValueOf(a).Int() -> panic 발생
  var pValue int64 = reflect.ValueOf(a).Elem().Int()
  
  Expect(pType).Should(Equal("*int"))
  Expect(pValueType).Should(Equal("1"))
  Expect(pValue).Should(Equal(int64(1)))
  
  var b interface{}
  b = 2
  
  bType := fmt.Sprint(reflect.TypeOf(b))
  bValueOf := fmt.Sprint(reflect.ValueOf(b))
  bValue := reflect.ValueOf(b).Int()
  
  Expect(bType).Should(Equal("int"))
  Expect(bValueOf).Should(Equal("2"))
  Expect(bValue).Should(Equal(int64(2)))
  ```

  