package entity

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
	"gorm.io/gorm"
)

type Section struct {
	gorm.Model
	Sec_Name   string
	Sec_Salary int
	Leave    []Leave `gorm:"foreignkey:Section_ID"`
}

type L_Type struct {
	gorm.Model
	Type_Name      string
	Type_Condition string
	Type_NTime     int
	Leave         []Leave `gorm:"foreignkey:L_Type_ID"`
}

type Leave struct {
	gorm.Model
	Doc_Reason string `json:"Doc_Reason" valid:"required~กรุณากรอกเหตุผลการลา / รายละเอียด"`
	Doc_DateS  time.Time `valid:"Past~วันที่เริ่มลาต้องไม่เป็นวันที่ผ่านมาแล้ว"`
	Doc_DateE  time.Time `valid:"Past~วันที่สิ้นสุดการลาต้องไม่เป็นวันที่ผ่านมาแล้ว"`
	Doc_Cont   string `json:"Doc_Cont" valid:"required~กรุณากรอกเบอร์ติดต่อ"`

	Section_ID  *uint `valid:"-"`
	Section     Section `valid:"-"`
	L_Type_ID   *uint `valid:"-"`
	L_Type      L_Type `valid:"-"`
	Employee_ID *uint `valid:"-"`
	Employee    Employee `valid:"-"`
}

func TestReasonNotBlank(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	tomorrow := time.Now().Add(24 * time.Hour);

	t.Run("check Reason can not blank", func(t *testing.T) {
		leave := Leave{
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
		leave := Leave{
			Doc_Reason: "im sick",
			Doc_DateS:  time.Now(),
			Doc_DateE: tomorrow,
			Doc_Cont: "",
		}
		ok, err := govalidator.ValidateStruct(leave)
		g.Expect(ok).NotTo(gomega.BeTrue())
		g.Expect(err).ToNot(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("กรุณากรอกกรอกช่องทางการติดต่อ"))
	})
}
func TestLeaveEndTimeNotPass(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	yesterday := time.Now().AddDate(0, 0, -1);

	t.Run("check End date cannot be past", func(t *testing.T) {
		leave := Leave{
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
