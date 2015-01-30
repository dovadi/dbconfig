package railsdbconfig

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

//DbParameters with all the rails settings of the corresponding database
type DbParameters struct {
	Adapter             string
	Encoding            string
	Database            string
	Username            string
	Password            string
	Allow_concurrency   bool //We need to use underscore to follow Rails naming conventions
	Timeout             int
	Pool                int
	Host                string
	Socket              string
	Prepared_statements bool
	Statement_limit     int
}

//DbConfig contains all the information from the database yaml config file
type DbConfig struct {
	Development DbParameters
	Test        DbParameters
	Production  DbParameters
	Staging     DbParameters
}

//LoadYamlConfig is loading the yaml config file
func LoadYamlConfig(path string) (DbConfig, error) {
	var dbconfig = DbConfig{}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = yaml.Unmarshal(data, &dbconfig)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return dbconfig, err
}
