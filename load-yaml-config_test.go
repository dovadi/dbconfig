package railsdbconfig_test

import (
	"github.com/dovadi/railsdbconfig"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Reading the Rails database yml file with default settings", func() {

	var (
		dbConf railsdbconfig.DbConfig
		err    error
	)

	BeforeEach(func() {
		dbConf, err = railsdbconfig.LoadYamlConfig("database.yml")
	})

	It("should load without an error", func() {
		Expect(err).Should(BeNil())
	})

	It("should return the database name of the development environment", func() {
		Expect(dbConf.Development.Database).Should(Equal("blog_development"))
	})

	It("should return the database name of the production environment", func() {
		Expect(dbConf.Production.Statement_limit).Should(Equal(200))
	})

})
