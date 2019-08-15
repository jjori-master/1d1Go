package factorial

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test functions Suite")
}

var _ = Describe("팩토리얼 테스트",
	func() {
		Context("팩토리얼", func() {
			It("팩토리얼", func() {

				r := factorial(5)
				Expect(r).Should(Equal(uint64(120)))
			})
		})
	},
)
