package dbconfig

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

//DbParameters with all the rails settings of the corresponding database
type DbParameters struct {
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

//DbConfig contains all the information from the database yaml config file
type DbConfig struct {
	Development DbParameters
	Test        DbParameters
	Production  DbParameters
	Staging     DbParameters
}

//LoadYamlConfig is loading the yaml config file
func LoadYamlConfig(path string) DbConfig {
	var dbconfig = DbConfig{}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(data, &dbconfig)
	if err != nil {
		panic(err)
	}

	return dbconfig
}
