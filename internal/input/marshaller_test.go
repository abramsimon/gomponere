package input_test

import (
	"../input"
	"../model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
)

var _ = Describe("Unmarshal", func() {
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
			data = []byte(areas)
		})

		It("does not error", func() {
			Expect(err).To(BeNil())
		})
		It("can unmarshall", func() {
			Expect(diagram.Areas).To(HaveLen(2))
			Expect(diagram.Areas["area-1"]).To(MatchFields(IgnoreExtras, Fields{
				"Name": Equal("The First Area"),
			}))
			Expect(diagram.Areas["area-2"]).To(MatchFields(IgnoreExtras, Fields{
				"Name":      Equal("The Second Area"),
				"ParentKey": Equal("area-1"),
			}))
		})
	})
	Context("with components", func() {
		BeforeEach(func() {
			data = []byte(components)
		})

		It("does not error", func() {
			Expect(err).To(BeNil())
		})
		It("can unmarshall", func() {
			Expect(diagram.Components).To(HaveLen(2))
			Expect(diagram.Components["component-1"]).To(MatchFields(IgnoreExtras, Fields{
				"Name": Equal("The First Component"),
			}))
			Expect(diagram.Components["component-2"]).To(MatchFields(IgnoreExtras, Fields{
				"Name":           Equal("The Second Component"),
				"LevelKey":       Equal("level-2"),
				"TypeKey":        Equal("type-2"),
				"TeamKey":        Equal("team-2"),
				"AreaKey":        Equal("area-2"),
				"DependencyKeys": ConsistOf("dep-2", "dep-3"),
			}))
		})
	})

	Context("with duplicates", func() {
		BeforeEach(func() {
			data = []byte(duplicates)
		})

		It("does error", func() {
			By("enforcing strict mode")
			Expect(err).ToNot(BeNil())
		})
	})

	Context("with everything", func() {
		BeforeEach(func() {
			data = []byte(areas + components + levels + teams + types)
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

const areas string = `
areas:
  area-1:
    name: The First Area
  area-2:
    name: The Second Area
    parent: area-1
`
const components string = `
components:
  component-1:
    name: The First Component
    level: level-1
    type: type-1
    team: team-1
    area: area-1
    dependencies:
      - dep-1
      - dep-2
  component-2:
    name: The Second Component
    level: level-2
    type: type-2
    team: team-2
    area: area-2
    dependencies:
      - dep-2
      - dep-3
`
const levels string = `
`
const teams string = `
`
const types string = `
`
const duplicates string = `
areas:
  area-1:
    name: The First Area
  area-1:
    name: The Second Area
    parent: area-1
`
