package unit36_reflection

import (
	"fmt"
	"reflect"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test functions Suite")
}

var _ = Describe("Unit 36 reflection", func() {
	Context("reflection 실습", func() {
		It("각 변수의 타입 확인하기 ", func() {
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
		})

		It("리플렉션으로 변수의 값에 대한 상세 정보를 가져온다.", func() {
			var f float64 = 1.3

			t := reflect.TypeOf(f)
			v := reflect.ValueOf(f)

			Expect(t.Name()).Should(Equal("float64"))
			Expect(fmt.Sprint(t.Size())).Should(Equal("8"))

			Expect(fmt.Sprint(v.Type())).Should(Equal("float64"))

			Expect(v.Kind() == reflect.Float64).Should(Equal(true))
			Expect(v.Kind() == reflect.Int64).Should(Equal(false))

			Expect(v.Float()).Should(Equal(1.3))
		})

		It("리플렉션으로 구조체 태그 이름 가져오기", func() {
			p := Person{}

			name, ok := reflect.TypeOf(p).FieldByName("name")
			Expect(ok).To(BeTrue())
			Expect(name.Tag.Get("tag1")).Should(Equal("이름"))
			Expect(name.Tag.Get("tag2")).Should(Equal("Name"))

			age, ok := reflect.TypeOf(p).FieldByName("age")
			Expect(ok).To(BeTrue())
			Expect(age.Tag.Get("tag1")).Should(Equal("나이"))
			Expect(age.Tag.Get("tag2")).Should(Equal("Age"))
		})
	})
})
