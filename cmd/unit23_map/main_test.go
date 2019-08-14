package unit23_map

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test functions Suite")
}

var _ = Describe("Map 사용",
	func() {
		Context("unit22_map 선언 및 데이터 넢기", func() {
			It("점수는 수학 100점, 영어 50점, 국어 80점이다.", func() {
				score := fetchScore()

				Expect(score["Math"]).Should(Equal(100))
				Expect(score["English"]).Should(Equal(50))
				Expect(score["Korean"]).Should(Equal(80))
			})

			It("국,영,수 점수의 합은 230점이다", func() {
				score := fetchScore()
				totalScore := calcTotalScore(score)
				Expect(totalScore).Should(Equal(230))
			})

			It("부끄러운 영어 점수는 삭제", func() {
				score := fetchScore()
				removeScore(score, "English")

				totalScore := calcTotalScore(score)
				Expect(totalScore).Should(Equal(180))
			})

			It("test", func() {

				score := map[string]int{
					"Math":    100,
					"English": 50,
					"Korean":  80,
				}

				totalScore := 0

				for key, value := range score {
					fmt.Println(key, " : ", value)
					totalScore += value
				}

				fmt.Println(totalScore) // 230
			})
		})
	},
)
