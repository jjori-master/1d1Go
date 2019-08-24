package unit31_struct_embedding

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test functions Suite")
}

var _ = Describe("Unit 31 struct embedding", func() {
	Context("struct embedding", func() {
		It("구조체 안의 구조체 Has a 관계 표현 ", func() {
			s := Students{}
			Expect(s.p.greeting()).Should(Equal("Say Hello!!"))
		})

		It("구조체 안의 구조체를 Embedding 함 Is a 관계 형성", func() {
			m := Man{}
			Expect(m.greeting()).Should(Equal("Say Hello!!"))
		})

		It("Embedding한 구조체 함수와 동일한 함수가 있다면 오버라이딩 된다.", func() {
			b := Bitch{}
			Expect(b.greeting()).Should(Equal("It's none your business"))
		})
	})
})
