package dbconfig

import (
	"encoding/json"
	"io/ioutil"
)

//JSONConfig is type for the configuration read from config.json
type JSONConfig struct {
	Environment   string
	Database_file string
}

//LoadJSONConfig is loading the json config file
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
