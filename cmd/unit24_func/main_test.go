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

			It("더하기와 빼기 연산 두개를 동시에 해보자", func() {
				sum, diff := sumNDiff(1, 2)
				Expect(sum).Should(Equal(3))
				Expect(diff).Should(Equal(-1))

				sum, diff = sumNDiff(7, 5)
				Expect(sum).Should(Equal(12))
				Expect(diff).Should(Equal(2))

				sum, _ = sumNDiff(12, 88)
				Expect(sum).Should(Equal(100))
			})

			It("가변인자 테스트 들어오는 모든 값을 더하기", func() {
				sum := sumAll(1, 2, 3)
				Expect(sum).Should(Equal(6))

				sum = sumAll(1, 2, 4, 8, 16)
				Expect(sum).Should(Equal(31))

				n := []int{1, 2, 3, 4, 5}
				sum = sumAll(n...)

				Expect(sum).Should(Equal(15))

			})
		})
	},
)
