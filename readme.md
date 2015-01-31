# dbconfig

[![Travis CI](https://secure.travis-ci.org/dovadi/dbconfig.png)](http://travis-ci.org/dovadi/dbconfig)

Golang package to read the database settings in the same way as the rails database.yml convention, see [Rails guide](http://guides.rubyonrails.org/configuring.html#configuring-a-database) and generate a connection string for the pq and mysql drivers.


Install
-------

```
go get -u github.com/dovadi/dbconfig
```


Example
-------

In the settings.json the location to the database yaml file and the application environment need to be defined

```Javascript
{
  "database_file": "/Users/dovadi/rails/blog/config/database.yml"
  "environment"  : "development"
}
```

Example of usage of Settings

```Go
package main

import (
  "fmt"

  "github.com/dovadi/dbconfig"
)

func main() {
  settings := dbconfig.Settings('settings.json')
  if err != nil {
    panic(err)
  }

  fmt.Println(settings.Database)
  fmt.Println(settings.Username)
  fmt.Println(settings.Password)
}
```

The most common settings are predefined:

```Go
type DbSettings struct {
  Development DbParameters
  Test        DbParameters
  Production  DbParameters
  Staging     DbParameters
}

type DbParameters struct {
  Adapter             string
  Encoding            string
  Database            string
  Username            string
  Password            string
  Port                string
  Allow_concurrency   string
  Timeout             string
  Pool                string
  Host                string
  Socket              string
  Prepared_statements string
  Statement_limit     string
}
```

Connection string
-----------------

Use dbconfig for generating a connection string for the pq or mysql driver

```Go
package main

import (
  "database/sql"
  "fmt"

  _ "github.com/go-sql-driver/mysql"
  _ "github.com/lib/pq"

  "github.com/dovadi/dbconfig"
)

func main() {

  connectionString := dbconfig.PostgresConnectionString("settings.json", "disable") // second parameter for sslmode
  // => "host=dbserver.org password=password user=dbuser dbname=blog_production sslmode=disable"
  db, err := sql.Open("postgres", connectionString)

  // connectionString := dbconfig.MysqlConnectionString("settings.json")
  // => "dbuser:password@tcp(dbserver.org:3309)/blog_production"
  // db, err := sql.Open("mysql", connectionString)
  
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()

  .
  .
  .
}
```

Run the tests
-------------

```
go test
```
or with the use of ginkgo

```
ginkgo -watch=true
```



Copyright
---------

Copyright (c) 2015 Frank Oxener. See LICENSE.txt for further details.
