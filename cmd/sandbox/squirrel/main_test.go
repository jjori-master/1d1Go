package squirrel

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	sq "github.com/Masterminds/squirrel"
)

func TestGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test functions Suite")
}

var _ = Describe("squirrel 청솔모 프레임워크? 테스트", func() {
	Context("squirrel 을 사용하여 여러 쿼리 문법 생성", func() {
		It("기본 select 쿼리 문장 생성", func() {
			users := sq.Select("*").From("users").Where(sq.Eq{"deleted_at": nil})

			sql, args, error := users.ToSql()

			Expect(error).ShouldNot(HaveOccurred())

			const exceptedSql = "SELECT * FROM users WHERE deleted_at IS NULL"
			Expect(sql).Should(Equal(exceptedSql))

			fmt.Println("args is ", args)
		})
	})
})
