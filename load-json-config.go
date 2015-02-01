package dbconfig

import (
	"encoding/json"
	"io/ioutil"
)

//JSONConfig is defining the configuration (read from json config file).
type JSONConfig struct {
	Environment   string
	Database_file string
}

/*
LoadJSONConfig is loading the json config file defining
location to database.yml and the application environment.
*/
func LoadJSONConfig(path string) JSONConfig {
	jsonConfig := JSONConfig{}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &jsonConfig)
	if err != nil {
		panic(err)
	}

	return jsonConfig
}
