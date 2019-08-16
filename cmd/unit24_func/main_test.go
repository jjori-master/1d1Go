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

var _ = Describe("Unit 24 Go func",
	func() {

		Context("뭐! Go func도 변수처럼 쓸수 있다go??!!", func() {
			It("sum 함수 hello 변수에 대입", func() {

				var hello func(a int, b int) int = sum
				r := hello(1, 2)

				Expect(r).Should(Equal(3))

				world := sum
				r = world(13, 490)

				Expect(r).Should(Equal(503))

			})

			It("slice에도 함수가 들어갈 수 있어~~", func() {
				f := []func(int, int) int{sum, diff}

				r := f[0](20, 11)
				Expect(r).Should(Equal(31))

				r = f[1](10, 1)
				Expect(r).Should(Equal(9))
			})

			It("map에도 함수가 value로 들어 갈 수 있어~", func() {
				f := map[string]func(int, int) int{
					"sum":  sum,
					"diff": diff,
				}

				r := f["sum"](3, 5)
				Expect(r).Should(Equal(8))

				r = f["diff"](3, 5)
				Expect(r).Should(Equal(-2))
			})

			It("익명함수 사용", func() {
				r := func(a int, b int) int {
					return a + b
				}(1, 2)
				Expect(r).Should(Equal(3))

			})
		})

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
