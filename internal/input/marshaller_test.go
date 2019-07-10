package input_test

import (
	"../input"
	"../model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Unmarshall", func() {
	var (
		err     error
		data    []byte
		diagram model.Diagram
	)

	BeforeEach(func() {
	})

	JustBeforeEach(func() {
		diagram, err = input.Unmarshal(data)
	})

	Context("with no data", func() {
		BeforeEach(func() {
			data = nil
		})

		It("does error", func() {
			Expect(err).ToNot(BeNil())
		})
	})

	Context("with areas", func() {
		BeforeEach(func() {
			data = []byte(`areas:
  area-1:
    name: The First Area
  area-2:
    name: The Second Area
    parent: area-1
`)
		})

		It("does not error", func() {
			Expect(err).To(BeNil())
		})
		It("can unmarshall", func() {
			Expect(diagram.Areas).To(HaveLen(2))
			Expect(diagram.Areas["area-1"].Name).To(Equal("The First Area"))
			Expect(diagram.Areas["area-2"].Name).To(Equal("The Second Area"))
			Expect(diagram.Areas["area-2"].ParentKey).To(Equal("area-1"))
		})
	})
})
