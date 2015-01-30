package railsdbconfig_test

import (
	"github.com/dovadi/railsdbconfig"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Reading a json config file", func() {

	var (
		jsonConf railsdbconfig.JSONConfig
		err      error
	)

	BeforeEach(func() {
		jsonConf, err = railsdbconfig.LoadJSONConfig("config.json")
	})

	It("should load without an error", func() {
		Expect(err).Should(BeNil())
	})

	It("should return the directory to the given rails app", func() {
		Expect(jsonConf.RailsDir()).Should(Equal("/Users/dovadi/rails/blog/"))
	})

})
