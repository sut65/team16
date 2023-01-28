package goapi_test

import (
	"testing"
	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
	"gorm.io/gorm"
)

type Videos struct {
	gorm.Model
	Name string  `valid:"required~Name cannot be blank"`
	Url string `gorm:"uniqueIndex" valid:"url"`
}

func TestVideoValidate(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	t.Run("check name not blank and check url", func(t *testing.T) {
		vdo := Videos{
			Name:"Me",
			Url:"https://example.com",
		}
	ok, err := govalidator.ValidateStruct(vdo)

	g.Expect(ok).To(gomega.BeTrue())
	g.Expect(err).To(gomega.BeNil())
	})
}
func TestNameNotBlank(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	t.Run("check name not blank", func(t *testing.T) {
		vdo := Videos{
			Name: "",
			Url:"https://example.com",
		}
	ok, err := govalidator.ValidateStruct(vdo)

	g.Expect(ok).NotTo(gomega.BeTrue())

	g.Expect(err).ToNot(gomega.BeNil())

	g.Expect(err.Error()).To(gomega.Equal("Name cannot be blank"))
	})
}
func TestCheckUrl(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	t.Run("check url",func(t *testing.T) {
		vdo := Videos{
			Name: "Me",
			Url:"example",
		}
		ok, err := govalidator.ValidateStruct(vdo)

	g.Expect(ok).NotTo(gomega.BeTrue())

	g.Expect(err).ToNot(gomega.BeNil())

	g.Expect(err.Error()).To(gomega.Equal("Url: example does not validate as url"))
	})
}

