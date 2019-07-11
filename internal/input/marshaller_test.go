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
		It("can unmarshal", func() {
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
		It("can unmarshal", func() {
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
	Context("with levels", func() {
		BeforeEach(func() {
			data = []byte(levels)
		})

		It("does not error", func() {
			Expect(err).To(BeNil())
		})
		It("can unmarshal", func() {
			Expect(diagram.Levels).To(HaveLen(2))
			Expect(diagram.Levels["level-1"]).To(MatchFields(IgnoreExtras, Fields{
				"Name": Equal("The First Level"),
			}))
			Expect(diagram.Levels["level-2"]).To(MatchFields(IgnoreExtras, Fields{
				"Name": Equal("The Second Level"),
			}))
		})
	})
	Context("with teams", func() {
		BeforeEach(func() {
			data = []byte(teams)
		})

		It("does not error", func() {
			Expect(err).To(BeNil())
		})
		It("can unmarshal", func() {
			Expect(diagram.Teams).To(HaveLen(2))
			Expect(diagram.Teams["team-1"]).To(MatchFields(IgnoreExtras, Fields{
				"Name": Equal("The First Team"),
			}))
			Expect(diagram.Teams["team-2"]).To(MatchFields(IgnoreExtras, Fields{
				"Name": Equal("The Second Team"),
				"TeamContact": MatchFields(IgnoreExtras, Fields{
					"Email": Equal("team@team2.com"),
				}),
				"LeadContact": MatchFields(IgnoreExtras, Fields{
					"Name":  Equal("The Lead"),
					"Email": Equal("lead@team2.com"),
				}),
				"Display": MatchFields(IgnoreExtras, Fields{
					"BackgroundColor": Equal("coral"),
					"ForegroundColor": Equal("black"),
				}),
			}))
		})
	})
	Context("with types", func() {
		BeforeEach(func() {
			data = []byte(types)
		})

		It("does not error", func() {
			Expect(err).To(BeNil())
		})
		It("can unmarshal", func() {
			Expect(diagram.Types).To(HaveLen(2))
			Expect(diagram.Types["type-1"]).To(MatchFields(IgnoreExtras, Fields{
				"Name": Equal("The First Type"),
			}))
			Expect(diagram.Types["type-2"]).To(MatchFields(IgnoreExtras, Fields{
				"Name":        Equal("The Second Type"),
				"Description": Equal("The type after the first type"),
				"Shape":       Equal("the-shape"),
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
		It("can unmarshal", func() {
			Expect(diagram).To(MatchFields(IgnoreExtras, Fields{
				"Areas":      HaveLen(2),
				"Components": HaveLen(2),
				"Levels":     HaveLen(2),
				"Teams":      HaveLen(2),
				"Types":      HaveLen(2),
			}))
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
levels:
  level-1:
    name: The First Level
  level-2:
    name: The Second Level
`
const teams string = `
teams:
  team-1:
    name: The First Team
  team-2:
    name: The Second Team
    team-contact:
      email: team@team2.com
    lead-contact:
      name: The Lead
      email: lead@team2.com
    display:
        background-color: coral
        foreground-color: black
`
const types string = `
types:
  type-1:
    name: The First Type
  type-2:
    name: The Second Type
    description: The type after the first type
    shape: the-shape
`
const duplicates string = `
areas:
  area-1:
    name: The First Area
  area-1:
    name: The Second Area
    parent: area-1
`
