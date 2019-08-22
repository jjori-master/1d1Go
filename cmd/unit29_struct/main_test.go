package unit29_struct

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test functions Suite")
}

var _ = Describe("Unit 28 struct", func() {
	Context("struct 실습", func() {
		It("구조체 선언 ", func() {
			var rect0 Rectangle
			rect0 = Rectangle{10, 20}
			Expect(rect0.width).Should(Equal(10))
			Expect(rect0.height).Should(Equal(20))

			var rect1 *Rectangle   // 구조체 포인터 선언
			rect1 = new(Rectangle) // 구조체 메모리 할당

			Expect(rect1.height).Should(Equal(0))

			rect2 := new(Rectangle)

			Expect(rect2.height).Should(Equal(0))
		})

		It("구조체 생성시 값 초기화", func() {
			var rect1 Rectangle = Rectangle{10, 20}

			Expect(rect1.width).Should(Equal(10))
			Expect(rect1.height).Should(Equal(20))
		})

		It("구조체 생성자 패턴으로 생성하기", func() {
			rect1 := NewRectangle(10, 20)
			Expect(rect1.width).Should(Equal(10))
			Expect(rect1.height).Should(Equal(20))

			rect2 := NewRectangle(30, 10)
			Expect(rect2.width).Should(Equal(30))
			Expect(rect2.height).Should(Equal(10))

			Expect(rect1.width).Should(Equal(10))
			Expect(rect1.height).Should(Equal(20))
		})

		It("구조체 인자를 주소로 넘겨줘서 값이 변경됨을 확인", func() {
			rect1 := Rectangle{10, 20}

			rectangleScaleA(&rect1, 10)

			Expect(rect1.width).Should(Equal(100))
			Expect(rect1.height).Should(Equal(200))
		})

		It("구조체 인자를 값으로 넘겨줘서 값이 변경이 안됨을 확인", func() {
			rect1 := Rectangle{10, 20}

			rectangleScaleB(rect1, 10)

			Expect(rect1.width).Should(Equal(10))
			Expect(rect1.height).Should(Equal(20))
		})
	})
})
