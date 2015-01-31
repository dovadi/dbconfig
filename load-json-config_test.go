package dbconfig_test

import (
	"github.com/dovadi/dbconfig"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Reading a json config file", func() {

	var jsonConf dbconfig.JSONConfig

	BeforeEach(func() {
		jsonConf = dbconfig.LoadJSONConfig("db-settings.json")
	})

	It("should return the directory to the given rails app", func() {
		Expect(jsonConf.Database_file).Should(Equal("/Users/dovadi/rails/blog/config/database.yml"))
	})

	It("should return the rails environment", func() {
		Expect(jsonConf.Environment).Should(Equal("development"))
	})

})
