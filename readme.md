# railsdbconfig

Golang package to read the Rails database settings


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
  Allow_concurrency   bool
  Timeout             int
  Pool                int
  Host                string
  Socket              string
  Prepared_statements bool
  Statement_limit     int
}

```

Copyright
---------

Copyright (c) 2015 Frank Oxener. See LICENSE.txt for further details.
