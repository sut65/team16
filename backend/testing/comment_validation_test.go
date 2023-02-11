package test

import (
	"testing"
	"time"
	"github.com/Team16/farm_mart/entity"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

// ตรวจสอบค่าว่างของชื่อแล้วต้องเจอ Error
func TestCommentNotBeBank(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Comment cannot be blank", func(t *testing.T) {
		com := entity.Comment{
			Comments: "",
			Date_Now: time.Now(),
			Bought_now:   2,
		}

		ok, err := govalidator.ValidateStruct(com)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("ความคิดเห็นต้องไม่เป็นค่าว่าง"))
	})

}

func TestDate_NowCannotBePast(t *testing.T) {
	g := NewGomegaWithT(t)
	// today := time.Now();
	// tomorrow := today.Add(24 * time.Hour);
	yesterday := time.Now().AddDate(0, 0, -1);

	t.Run("Date Now cannot be past", func(t *testing.T) {
		com := entity.Comment{
			Comments: "test",
			Date_Now: yesterday,
			Bought_now:   2,
		}

		ok, err := govalidator.ValidateStruct(com)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("วันที่ต้องไม่เป็นอดีต"))
	})

}

func TestDate_NowCannotBeFuture(t *testing.T) {
	g := NewGomegaWithT(t)
	today := time.Now();
	tomorrow := today.Add(24 * time.Hour);
	// yesterday := time.Now().AddDate(0, 0, -1);

	t.Run("Date Now cannot be future", func(t *testing.T) {
		com := entity.Comment{
			Comments: "test",
			Date_Now: tomorrow,
			Bought_now:   2,
		}

		ok, err := govalidator.ValidateStruct(com)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("วันที่ต้องไม่เป็นอนาคต"))
	})

}

func TestBought_NowCannotBeNegative(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Date now cannot be negative number", func(t *testing.T) {
		com := entity.Comment{
			Comments: "test",
			Date_Now: time.Now(),
			Bought_now:   -2,
		}

		ok, err := govalidator.ValidateStruct(com)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("กรุณากรอกจำนวนเต็มบวกเท่านั้น"))
	})

}







