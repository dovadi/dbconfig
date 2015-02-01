/*
Package dbconfig reads database settings from a database.yml file (following the rails database.yml convention)
and generates a connection string for the github.com/lib/pq and github.com/go-sql-driver/mysql drivers.
*/
package dbconfig

import (
	"os"
	"strings"
)

/*
Settings returns the database settings from the database.yml file from a given application
and the corresponding application enviroment. The location to the database.yml file and
the enviroment is configured in the settings json configuration file.
If environment is NOT configured, you can set the environment variable APPLICATION_ENV (on os level).
If this is also not defined "development" is the default.
*/
func Settings(path string) DbParameters {
	var environment string

	jsonConf := LoadJSONConfig(path)
	dbConfig := LoadYamlConfig(jsonConf.Database_file)

	if len(jsonConf.Environment) == 0 {
		if len(os.Getenv("APPLICATION_ENV")) > 0 {
			environment = os.Getenv("APPLICATION_ENV")
		} else {
			environment = "development"
		}
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
PostgresConnectionString returns the connection string to open a sql session
used by the github.com/lib/pq package like for example:
"host=dbserver.org password=password user=dbuser dbname=blog_production sslmode=disable"
The first parameter is the path to the database settings configuration (json) file
and the second paramater defines the sslmode.
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

/*
MysqlConnectionString returns the connection string to open a sql session used by the github.com/go-sql-driver/mysql/
For example: "username:password@tcp(localhost:3306)/dbname:
Pass the path to the database settings configuration (json) file.
*/
func MysqlConnectionString(path string) string {
	var port string
	settings := Settings(path)

	if len(settings.Port) == 0 {
		port = "3306"
	} else {
		port = settings.Port
	}

	connection := []string{
		settings.Username, ":",
		settings.Password, "@tcp(",
		settings.Host, ":",
		port, ")/",
		settings.Database, ""}

	return strings.Join(connection, "")
}
