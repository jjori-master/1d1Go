package unit27_panic

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test functions Suite")
}

var _ = Describe("Unit 27 panic", func() {
	Context("panic 실습", func() {
		It("panic 발생 함수 실행 ", func() {
			ExamplePanic()
		})

		It("panic array 범위 초과", func() {
			defer func() {
				if r := recover(); r != nil {
					Expect(r).Should(Equal("panic!!!"))
				}
			}()
			panic1()
		})
	})
})
