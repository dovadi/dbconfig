package dbconfig_test

import (
	"os"

	"github.com/dovadi/dbconfig"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Reading the Rails database yml file with default settings", func() {

	var dbConf dbconfig.DbConfig

	BeforeEach(func() {
		os.Setenv("POSTGRESQL_PASSWORD", "password")
		dbConf = dbconfig.LoadYamlConfig("test-files/database.yml")
	})

	It("should return the database name of the development environment", func() {
		Expect(dbConf["development"]["database"]).Should(Equal("blog_development"))
	})

	It("should return the database name of the production environment", func() {
		Expect(dbConf["production"]["statement_limit"]).Should(Equal("200"))
	})

	It("should return the get corresponding environment variable if defined in erb tags", func() {
		Expect(dbConf["test"]["password"]).Should(Equal("password"))
	})

})
