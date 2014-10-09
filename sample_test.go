package go_gospel

import (
	. "github.com/r7kamura/gospel"
	"testing"
)

func TestSpec(t *testing.T) {
	Describe(t, "Hello World を取得するクラス", func() {
		Context("Hello Worldを取得する", func() {
			It("関数を呼び出しHello World!をが取得できる", func() {
				sut := &HelloWorld{}
				Expect(sut.say()).To(Equal, "Hello World!")
			})
		})
	})
}
