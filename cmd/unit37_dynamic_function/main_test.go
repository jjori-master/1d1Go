package unit37_dynamic_function

import (
	"reflect"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test functions Suite")
}

var _ = Describe("Unit 37 동적 함수 생성", func() {
	Context("reflect를 이용하여 함수 동적으로 생성하기", func() {
		It("예제 1번 함수 생성 ", func() {
			// 함수를 담을 변수 생성
			var hello func()

			// hello 의 주소를 넘긴 뒤 Elem으로 값 정보를 가져옴
			fn := reflect.ValueOf(&hello).Elem()

			v := reflect.MakeFunc(fn.Type(), h)

			fn.Set(v)

			hello()
		})

		It("함수를 생성하는 makeFunc 함수는 동적으로 함수를 생성 가능하다.", func() {
			var intSum func(int, int) (int64, error)
			var floatSum func(float64, float64) (float64, error)

			makeSum(&intSum)

			_, err := intSum(1, 1)
			Expect(err).To(BeNil())

			makeSum(&floatSum)
			_, err = floatSum(1.2, 1.1)
			Expect(err).To(BeNil())
		})

	})
})
