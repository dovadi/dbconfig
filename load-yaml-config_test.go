package dbconfig_test

import (
	"github.com/dovadi/dbconfig"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Reading the Rails database yml file with default settings", func() {

	var dbConf dbconfig.DbConfig

	BeforeEach(func() {
		dbConf = dbconfig.LoadYamlConfig("database.yml")
	})

	It("should return the database name of the development environment", func() {
		Expect(dbConf.Development.Database).Should(Equal("blog_development"))
	})

	It("should return the database name of the production environment", func() {
		Expect(dbConf.Production.Statement_limit).Should(Equal("200"))
	})

})
