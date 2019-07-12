package diagram_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDiagram(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Diagram Suite")
}
