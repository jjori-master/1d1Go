package unit30_struct_func

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test functions Suite")
}

var _ = Describe("Unit 30 구조체에 메서드 연결하기", func() {
	Context("구조체에 메서드 연결하기 실습", func() {
		It("구조체는 메서드를 연결 할 수 있다.", func() {

			rect := Rectangle{10, 20}

			area := rect.area()

			Expect(area).Should(Equal(area))
		})

		It("포인터로 구조체 리시버 변수를 받은 경우 값이 변경되어 있다.", func() {
			rect1 := Rectangle{10, 20}

			rect1.scaleA(10)

			Expect(rect1.width).Should(Equal(100))
			Expect(rect1.height).Should(Equal(200))
		})

		It("일반 구조체 리시버 변수를 받은 경우 값이 변경되지 않는다.", func() {

			rect2 := Rectangle{10, 20}

			rect2.scaleB(10)

			Expect(rect2.width).Should(Equal(10))
			Expect(rect2.height).Should(Equal(20))
		})
	})
})
