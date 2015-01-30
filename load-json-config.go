package railsdbconfig

import (
	"encoding/json"
	"io/ioutil"
)

//JSONConfig is type for the configuration read from config.json
type JSONConfig map[string]interface{}

//RailsDir return the directory to the Rails app
func (c JSONConfig) RailsDir() string {
	return c["rails-dir"].(string)
}

//LoadJSONConfig is loading the json config file
func LoadJSONConfig(path string) (JSONConfig, error) {
	jsonConfig := JSONConfig{}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return jsonConfig, err
	}

	err = json.Unmarshal(data, &jsonConfig)
	return jsonConfig, err
}
