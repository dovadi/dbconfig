package dbconfig_test

import (
	"os"

	"github.com/dovadi/dbconfig"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("dbconfig", func() {

	It("should return the database settings of the corresponding application environment", func() {
		settings := dbconfig.Settings("test-files/settings.json")
		Expect(settings.Database).Should(Equal("blog_production"))
	})

	Context("For the corresponding application environment", func() {

		It("should return the correct connection string for postgresql", func() {
			connectionString := dbconfig.PostgresConnectionString("test-files/settings.json", "disable")
			Expect(connectionString).Should(Equal("host=dbserver.org password=password user=dbuser dbname=blog_production sslmode=disable"))
		})

		It("should return the correct connection string for mysql", func() {
			connectionString := dbconfig.MysqlConnectionString("test-files/settings.json")
			Expect(connectionString).Should(Equal("dbuser:password@tcp(dbserver.org:3309)/blog_production"))
		})

	})

	Context("Application environment not set in json config file", func() {

		It("should return the database settings of the default development environment", func() {
			settings := dbconfig.Settings("test-files/settings_without_env.json")
			Expect(settings.Database).Should(Equal("blog_development"))
		})

		It("should return the database settings of the corresponding environment set as environment variable", func() {
			os.Setenv("APPLICATION_ENV", "test")
			settings := dbconfig.Settings("test-files/settings_without_env.json")
			Expect(settings.Database).Should(Equal("blog_test"))
		})

	})

})
