package go_gospel

import (
	. "github.com/r7kamura/gospel"
	"testing"
)

func TestPersonSpec(t *testing.T) {
	Describe(t, "person class", func() {
		Context("create person", func() {
			It("get property", func() {
				sut := &Person{"yamada"}
				Expect(sut.LastName).To(Equal, "yamada")
			})
		})
	})
}
