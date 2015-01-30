package railsdbconfig

import "log"

/*
Settings returns the database settings from the database.yml file from a given rails app
Rails directory is configured in the config.json file
*/
func Settings() (DbConfig, error) {
	jsonConf, err := LoadJSONConfig("config.json")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	dbConfig, err := LoadYamlConfig(jsonConf.RailsDir() + "/config/database.yml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return dbConfig, err
}
