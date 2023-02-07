package entity

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

// ตรวจสอบค่าว่างของชื่อแล้วต้องเจอ Error
func TestCommentNotBeBank(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Comment cannot be blank", func(t *testing.T) {
		sep := Separation{
			Date_Out: time.Now(),
			Amount:   2,
			Status:   "",
		}

		ok, err := govalidator.ValidateStruct(sep)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("ความคิดเห็นต้องไม่เป็นค่าว่าง"))
	})

}
// ตรวจสอบค่าว่างของชื่อแล้วต้องเจอ Error
func TestDate_NowNotBeBank(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Date Now cannot be blank", func(t *testing.T) {
		sep := Separation{
			Date_Out: time.Now(),
			Amount:   2,
			Status:   "",
		}

		ok, err := govalidator.ValidateStruct(sep)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("วันที่ต้องไม่เป็นค่าว่าง"))
	})

}

func TestDate_NowCannotBePast(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Date Now cannot be past", func(t *testing.T) {
		sep := Separation{
			Date_Out: time.Now(),
			Amount:   2,
			Status:   "",
		}

		ok, err := govalidator.ValidateStruct(sep)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("วันที่ต้องไม่เป็นอดีต"))
	})

}
func TestDate_NowCannotBeFuture(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Date Now cannot be future", func(t *testing.T) {
		sep := Separation{
			Date_Out: time.Now(),
			Amount:   2,
			Status:   "",
		}

		ok, err := govalidator.ValidateStruct(sep)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("วันที่ต้องไม่เป็นอนาคต"))
	})

}







