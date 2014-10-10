package go_gospel

import (
	. "github.com/r7kamura/gospel"
	"testing"
)

func TestPersonSpec(t *testing.T) {
	Describe(t, "person class", func() {
		Context("create instance by specifying the parameters to yamada", func() {
			sut := &Person{"yamada", "taro"}
			It("get lastname", func() {
				Expect(sut.LastName).To(Equal, "yamada")
			})
			It("get firstname", func() {
				Expect(sut.FirstName).To(Equal, "taro")
			})
		})
		Context("create instance by specifying the parameters to Sato", func() {
			sut := &Person{"sato", "ichiro"}
			It("get last name", func() {
				Expect(sut.LastName).To(Equal, "sato")
			})
			It("get first name", func() {
				Expect(sut.FirstName).To(Equal, "ichiro")
			})
		})
	})
}
