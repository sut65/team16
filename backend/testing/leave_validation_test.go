package test

import (
	"testing"
	"time"
	"github.com/Team16/farm_mart/entity"
	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
)

func TestReasonNotBlank(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	tomorrow := time.Now().Add(24 * time.Hour);

	t.Run("check Reason can not blank", func(t *testing.T) {
		leave := entity.Leave{
			Doc_Reason: "",
			Doc_DateS:  time.Now(),
			Doc_DateE: tomorrow,
			Doc_Cont: "0930963238",
		}

		ok, err := govalidator.ValidateStruct(leave)
		g.Expect(ok).NotTo(gomega.BeTrue())
		g.Expect(err).ToNot(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("กรุณากรอกเหตุผลการลา / รายละเอียด"))
	})
}
func TestLContact(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	tomorrow := time.Now().Add(24 * time.Hour);

	t.Run("check Contact can not blank", func(t *testing.T) {
		leave := entity.Leave {
			Doc_Reason: "im sick",
			Doc_DateS:  time.Now(),
			Doc_DateE: tomorrow,
			Doc_Cont: "",
		}
		ok, err := govalidator.ValidateStruct(leave)
		g.Expect(ok).NotTo(gomega.BeTrue())
		g.Expect(err).ToNot(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("กรุณากรอกช่องทางการติดต่อ"))
	})
}
func TestLeaveEndTimeNotPass(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	yesterday := time.Now().AddDate(0, 0, -1);

	t.Run("check End date cannot be past", func(t *testing.T) {
		leave := entity.Leave{
			Doc_Reason: "im sick",
			Doc_DateS:  time.Now(),
			Doc_DateE: yesterday,
			Doc_Cont: "0930963238",
		}
		ok, err := govalidator.ValidateStruct(leave)
		g.Expect(ok).NotTo(gomega.BeTrue())
		g.Expect(err).ToNot(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("วันสิ้นสุดการลาต้องไม่เป็นอดีต"))
	})
}
