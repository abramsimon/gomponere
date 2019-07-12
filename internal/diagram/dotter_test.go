package diagram_test

import (
	"../diagram"
	"../model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MakeDot", func() {
	var (
		err error
		d   model.Diagram
		dot string
	)

	JustBeforeEach(func() {
		dot, err = diagram.MakeDot(d)
	})

	Context("with empty diagram", func() {
		BeforeEach(func() {
			d = model.Diagram{}
		})

		It("does not error", func() {
			Expect(err).To(BeNil())
		})
		It("has empty dot", func() {
			Expect(dot).To(ContainSubstring("digraph"))
		})
	})
})
