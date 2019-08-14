package unit24_func

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test functions Suite")
}

var _ = Describe("사칙 연산",
	func() {

		Context("사칙 연산을 해보자", func() {
			It("더하기 연산", func() {
				result := sum(1, 1)
				Expect(result).Should(Equal(2))

				result = sum(1, 2)
				Expect(result).Should(Equal(3))

				result = sum(1000, 200)
				Expect(result).Should(Equal(1200))
			})
		})
	},
)
