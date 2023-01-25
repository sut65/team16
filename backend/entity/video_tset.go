package entity

// import (
// 	"testing"

// 	"github.com/asaskevich/govalidator"
// 	. "github.com/onsi/gomega"
// 	"gorm.io/gorm"
// )

// type Video struct {
// 	gorm.Model
// 	Name string `valid:"required~Name cannot be blank"`
// 	Url  string `gorm:"uniqueIndex" valid:"url~Url incorrect"`
// }

// func TestUserNameNotBlank(t *testing.T) {
// 	g := NewGomegaWithT(t)

// 	t.Run("check name not blank", func(t *testing.T) {
// 		user := Video{
// 			Name: "",
// 			Url:  "https://github.com/asaskevich/govalidator",
// 		}

// 		ok, err := govalidator.ValidateStruct(user)

// 		g.Expect(ok).NotTo(BeTrue())
// 		g.Expect(err).ToNot(BeNil())
// 		g.Expect(err.Error()).To(Equal("Name cannot be blank"))
// 	})

// 	t.Run("check url is valid", func(t *testing.T) {
// 		user := Video{
// 			Name: "keng",
// 			Url:  "htppabc//ทดสอบ",
// 		}
// 		ok, err := govalidator.ValidateStruct(user)
// 		g.Expect(ok).NotTo(BeTrue())
// 		g.Expect(err).ToNot(BeNil())
// 		g.Expect(err.Error()).To(Equal("Url incorrect"))
// 	})

// 	t.Run("check correct", func(t *testing.T) {
// 		user := Video{
// 			Name: "keng",
// 			Url:  "https://github.com/asaskevich/govalidator",
// 		}
// 		ok, err := govalidator.ValidateStruct(user)
// 		g.Expect(ok).To(BeTrue())
// 		g.Expect(err).To(BeNil())
// 		//g.Expect(err.Error()).To(Equal("Url incorrect"))
// 	})
// }

// func NewGomegaWithT(t *testing.T) {
// 	panic("unimplemented")
// }