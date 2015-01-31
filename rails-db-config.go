package dbconfig

import (
	"os"
	"strings"
)

/*
Settings returns the database settings from the database.yml file from a given application
and the corresponding application enviroment.
The location to the database.yml file and the enviroment is configured in the settings json configuration file
*/
func Settings(path string) DbParameters {
	var environment string

	jsonConf := LoadJSONConfig(path)
	dbConfig := LoadYamlConfig(jsonConf.Database_file)

	if len(jsonConf.Environment) == 0 {
		environment = os.Getenv("APPLICATION_ENV")
	} else {
		environment = jsonConf.Environment
	}

	switch environment {
	case "staging":
		return dbConfig.Staging
	case "test":
		return dbConfig.Test
	case "production":
		return dbConfig.Production
	default:
		return dbConfig.Development
	}

}

/*
PostgresConnectionString returns the connection string to open a sql session used with the database/sql package
For example host=dbserver.org password=password user=dbuser dbname=blog_production sslmode=disable
First parameter is the path to the database settings configuration (json) file
Second paramater defines the sslmode
*/
func PostgresConnectionString(path string, sslmode string) string {
	settings := Settings(path)

	connection := []string{
		"host=", settings.Host, " ",
		"password=", settings.Password, " ",
		"user=", settings.Username, " ",
		"dbname=", settings.Database, " ",
		"sslmode=", sslmode}

	return strings.Join(connection, "")
}
