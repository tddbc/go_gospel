package go_gospel

import (
	. "github.com/r7kamura/gospel"
	"testing"
)

func TestPersonSpec(t *testing.T) {
	Describe(t, "person class", func() {
		Context("create instance by specifying the parameters to yamada", func() {
			sut := &Person{"yamada"}
			It("get lastname", func() {
				Expect(sut.LastName).To(Equal, "yamada")
			})
		})
		Context("create instance by specifying the parameters to Sato", func() {
			sut := &Person{"Sato"}
			It("get lastname", func() {
				Expect(sut.LastName).To(Equal, "Sato")
			})
		})
	})
}
