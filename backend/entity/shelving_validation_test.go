package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestAmountSH(t *testing.T) {
	g := NewGomegaWithT(t)
	fixture := []int{
		0, 21,
	}

	for _, sh := range fixture {
		sh := Shelving{
			Amount: sh,
		}

		ok, err := govalidator.ValidateStruct(sh)

		// ok ต้องไม่เป็น true แปลว่าต้องจับ error ได้
		g.Expect(ok).ToNot(BeTrue())

		// err ต้องไม่เป็น nil แปลว่าต้องจับ error ได้
		g.Expect(err).ToNot(BeNil())

		// err.Error() ต้องมี message แสดงออกมา
		g.Expect(err.Error()).To(Equal("A number of goods must be in the range 1-20;Amount must be in the range 1-1000;Name cannot be blank;Price must be in the range 1-1000"))
	}
}
