package unit28_pointer

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test functions Suite")
}

var _ = Describe("Unit 28 pointer", func() {
	Context("pointer 실습", func() {
		It("pointer는 기본값이 nil 이다.", func() {
			var numPtr *int
			Expect(numPtr).To(BeNil())

			var numNormal int
			Expect(numNormal).Should(Equal(0))
		})

		It("pointer는 new를 통해 메모리를 할당한다.", func() {
			var numPtr *int = new(int)

			Expect(numPtr).NotTo(BeNil())
		})

		It("ponter 변수에 값을 대입하거나 가져오라면 역참조(dereference)를 사용", func() {
			var numPtr *int = new(int)

			*numPtr = 1

			Expect(*numPtr).Should(Equal(1))
		})

		It("일반 변수에 참조(레퍼런스)를 사용하면 포인터형 변수에 대입 할 수 있다.", func() {
			var num int = 1

			var numPtr *int = &num

			Expect(&num).Should(Equal(numPtr))
		})
	})
})
