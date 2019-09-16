package unit37_dynamic_function

import (
	"fmt"
	"reflect"
)

func h(_ []reflect.Value) []reflect.Value {
	fmt.Println("Hello, world!")
	return nil
}

var nilError = reflect.Zero(reflect.TypeOf((*error)(nil)).Elem())

func sum(args []reflect.Value) []reflect.Value {
	a, b := args[0], args[1]

	if a.Kind() != b.Kind() {
		fmt.Println("타입이 다릅니다.!")
		return []reflect.Value{reflect.ValueOf(nil)}
	}

	switch a.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64:
		return []reflect.Value{reflect.ValueOf(a.Int() + b.Int()), nilError}

	case reflect.Float32, reflect.Float64:
		return []reflect.Value{reflect.ValueOf(a.Float() + b.Float()), nilError}
	}

	return []reflect.Value{reflect.ValueOf(0), nilError}
}

func makeSum(fptr interface{}) {
	fn := reflect.ValueOf(fptr).Elem()

	v := reflect.MakeFunc(fn.Type(), sum)

	fn.Set(v)
}
