package gohelloworld

import (
	. "github.com/r7kamura/gospel"
	"testing"
)

func TestDescribe(t *testing.T) {
	Describe(t, "HelloWorld", func() {
		Context("call HelloWorld", func() {
			It("should be 'Hello World'", func() {
				Expect(HelloWorld()).To(Equal, "Hello World")
			})
		})
	})
}
