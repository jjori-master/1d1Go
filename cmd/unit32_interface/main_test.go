package unit32_interface

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test functions Suite")
}

var _ = Describe("Unit 32 interface", func() {
	Context("Interface 실습", func() {
		It("Interface 선언, 기본 Interface는 nil 이다 ", func() {
			var h hello
			Expect(h).To(BeNil())
		})

		It("가로 6, 세로 3을 가지고 있는 직사각형의 넓이는 18이다. 18..", func() {
			rect := Rectangle{6, 3}

			var ac AreaCalculator
			ac = &rect

			Expect(ac.area()).Should(Equal(18))
		})

		It("가로 6, 세로 3을 가지고 있는 삼각형의 넓이는 9다.", func() {
			triangle := Triangle{6, 3}

			var ac AreaCalculator
			ac = &triangle

			Expect(ac.area()).Should(Equal(9))
		})

		It("인터페이스 선언 및 초기화", func() {
			rect := Rectangle{10, 20}
			triangle := Triangle{10, 20}

			ac1 := AreaCalculator(&rect)
			ac2 := AreaCalculator(&triangle)

			Expect(ac1.area()).Should(Equal(200))
			Expect(ac2.area()).Should(Equal(100))
		})
	})
})
