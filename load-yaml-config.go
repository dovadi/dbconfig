package dbconfig

import (
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/fatih/structs"
	"gopkg.in/yaml.v2"
)

var config = make(map[string]string)

//dbParameters with all the rails settings of the corresponding database.
type dbParameters struct {
	Adapter             string
	Encoding            string
	Database            string
	Username            string
	Password            string
	Port                string
	Allow_concurrency   string //We need to use underscore to follow Rails naming conventions
	Timeout             string
	Pool                string
	Host                string
	Socket              string
	Prepared_statements string
	Statement_limit     string
}

type dbYamlConfig struct {
	Development dbParameters
	Test        dbParameters
	Production  dbParameters
	Staging     dbParameters
}

//DbConfig contains all the information from the database yaml config file.
type DbConfig map[string]map[string]string

func parseErb(value string) string {
	if len(value) > 0 {
		re := regexp.MustCompile("<%=\\s+ENV\\['(.+)']\\s+%>")
		match := re.FindStringSubmatch(value)
		if len(match) == 2 {
			value = os.Getenv(match[1])
		}
	}
	return value
}

func assignToMultiMap(config map[string]string, converted map[string]interface{}) {
	for k, v := range converted {
		config[strings.ToLower(k)] = parseErb(v.(string))
	}
}

//LoadYamlConfig is loading the yaml config file.
func LoadYamlConfig(path string) DbConfig {
	var dbyamlconfig = dbYamlConfig{}
	var dbconfig = DbConfig{}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(data, &dbyamlconfig)
	if err != nil {
		panic(err)
	}

	convertedDevelopment := structs.Map(dbyamlconfig.Development)
	convertedTest := structs.Map(dbyamlconfig.Test)
	convertedStaging := structs.Map(dbyamlconfig.Staging)
	convertedProduction := structs.Map(dbyamlconfig.Production)

	config = make(map[string]string)
	assignToMultiMap(config, convertedDevelopment)
	dbconfig["development"] = config

	config = make(map[string]string)
	assignToMultiMap(config, convertedProduction)
	dbconfig["production"] = config

	config = make(map[string]string)
	assignToMultiMap(config, convertedTest)
	dbconfig["test"] = config

	config = make(map[string]string)
	assignToMultiMap(config, convertedStaging)
	dbconfig["staging"] = config

	return dbconfig
}
