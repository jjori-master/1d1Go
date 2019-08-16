package unit22_slice

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test functions Suite")
}

var _ = Describe("Unit 22 slice 테스트", func() {

	Context("slice 를 테스트 하자", func() {
		It("nil slice append", func() {
			var arr []int
			arr = append(arr, []int{1, 2, 3, 4, 5}...)

			Expect(arr).Should(Equal([]int{1, 2, 3, 4, 5}))
		})
	})
})
