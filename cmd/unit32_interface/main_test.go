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

		It("오리는 quack이라는 소리로 웁니다", func() {
			var donald Duck

			r := vocalCord(donald)

			Expect(r).Should(Equal("quack"))
		})

		It("사람은 꽥이라고 오리 소리를 흉내 냅니다", func() {
			var jo Person

			r := vocalCord(jo)

			Expect(r).Should(Equal("꽥"))
		})

		It("donal type은 Quacker 인터페이스 타입입니다..", func() {
			var donald Duck

			_, ok := interface{}(donald).(Quacker)
			Expect(ok).Should(Equal(true))

		})

		It("빈인터페이스는 무슨 타입이든 받아 줍니다.", func() {
			r1 := formatString(1)
			Expect(r1).Should(Equal("1"))

			r2 := formatString(2.5)
			Expect(r2).Should(Equal("2.5"))

			r3 := formatString("Hello, world")
			Expect(r3).Should(Equal("Hello, world"))

			p1 := Person{"Jo", 10}
			r4 := formatString(p1)
			Expect(r4).Should(Equal("Jo 10"))

			p2 := Person{"So", 20}
			r5 := formatString(&p2)
			Expect(r5).Should(Equal("So 20 1"))
		})
	})
})
