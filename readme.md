# dbconfig

[![Travis CI](https://secure.travis-ci.org/dovadi/dbconfig.png)](http://travis-ci.org/dovadi/dbconfig)

[![GoDoc](https://godoc.org/github.com/dovadi/dbconfig?status.svg)](https://godoc.org/github.com/dovadi/dbconfig)

Golang package to read the database settings following the rails database.yml convention, see [Rails guide](http://guides.rubyonrails.org/configuring.html#configuring-a-database) including embedded erb tags with environment variables and generate a connection string for the github.com/lib/pq and github.com/go-sql-driver/mysql drivers.


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

If no environment is defined in the json config file, you can define the environment variable APPLICATION_ENV (on operating system level), otherwise the default environment is "development"

Example of Settings

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

  fmt.Println(settings["database"])
  fmt.Println(settings["username"])
  fmt.Println(settings["password"])
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


Basic erb parsing
-----------------

A database.yml can contain erb tags with environment variables

```ruby
test:
  adapter: mysql2
  encoding: utf8
  database: blog_test
  pool: 5
  username: root
  password: <%= ENV['POSTGRESQL_PASSWORD'] %>
  socket: /tmp/mysql.sock
```

In this case the environment variable POSTGRESQL_PASSWORD is read in order to define password.

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
