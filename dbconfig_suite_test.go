package dbconfig_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRailsdbconfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Railsdbconfig Suite")
}
