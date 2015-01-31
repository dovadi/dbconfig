package dbconfig_test

import (
	"github.com/dovadi/dbconfig"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("dbconfig", func() {

	It("should return the database settings of the corresponding application environment", func() {
		settings := dbconfig.Settings("db-settings-test.json")
		Expect(settings.Database).Should(Equal("blog_production"))
	})

	Context("For the corresponding application environment", func() {

		It("should return the correct connection string for postgresql", func() {
			connectionString := dbconfig.PostgresConnectionString("db-settings-test.json", "disable")
			Expect(connectionString).Should(Equal("host=dbserver.org password=password user=dbuser dbname=blog_production sslmode=disable"))
		})

		It("should return the correct connection string for mysql", func() {
			connectionString := dbconfig.MysqlConnectionString("db-settings-test.json")
			Expect(connectionString).Should(Equal("dbuser:password@tcp(dbserver.org:3309)/blog_production"))
		})

	})

})
