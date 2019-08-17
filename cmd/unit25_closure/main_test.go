package unit25_closure

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test functions Suite")
}

var _ = Describe("Unit 25 closure", func() {
	Context("closure  실습", func() {
		It("일반 적인 closure", func() {
			f := calc()

			Expect(f(1)).Should(Equal(8))
			Expect(f(2)).Should(Equal(11))
			Expect(f(3)).Should(Equal(14))
		})

		It("closure 실습2", func() {
			f := sayHelloTo()
			Expect(f("종균")).Should(Equal("Hello 종균"))
		})
	})
})
