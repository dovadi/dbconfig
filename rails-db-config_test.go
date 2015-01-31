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

	It("should return the connection string of the corresponding application environment for postgresql", func() {
		connectionString := dbconfig.PostgresConnectionString("db-settings-test.json", "disable")
		Expect(connectionString).Should(Equal("host=dbserver.org password=password user=dbuser dbname=blog_production sslmode=disable"))
	})

})
