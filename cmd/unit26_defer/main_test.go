package unit26_defer

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test functions Suite")
}

var _ = Describe("Unit 26 defer", func() {
	Context("defer 실습", func() {
		It("hello 함수 종료 world 함수가 실행 ", func() {

			f := sayHello()
			greeting := f()

			Expect(greeting).Should(Equal("hello world!!"))
		})
	})
})
