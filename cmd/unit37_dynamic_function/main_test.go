package unit37_dynamic_function

import (
	"math"
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

		It("int, float, string 타입별 sum 함수 실행", func() {
			var intSum func(int, int) (int64, error)
			var floatSum func(float64, float64) (float64, error)
			var stringSum func(string, string) (string, error)

			makeSum(&intSum)
			makeSum(&floatSum)
			makeSum(&stringSum)

			v, err := intSum(1, 2)
			Expect(err).To(BeNil())
			Expect(v).Should(Equal(int64(3)))

			const epsilon = 1e-4
			v2, err2 := floatSum(1.2, 2.3)
			result := math.Abs((1.2+2.3)-v2) <= epsilon
			Expect(err2).To(BeNil())
			Expect(result).To(BeTrue())

			v3, err3 := stringSum("1", "1")
			Expect(err3).To(BeNil())
			Expect(v3).Should(Equal("11"))
		})
	})
})
