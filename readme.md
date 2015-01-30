# railsdbconfig

Golang package to read the Rails database settings


Example
-------

In the config.json the path to the rails directory needs to be defined

```Javascript
{
  "rails-dir": "/Users/dovadi/rails/blog/"
}
```

(The config.json should be in the root path of golang program)

```Go
package main

import (
  "fmt"

  "github.com/dovadi/railsdbconfig"
)

func main() {
  dbSettings, err := railsdbconfig.Settings()
  if err != nil {
    panic(err)
  }

  fmt.Println(dbSettings.Development.Database)
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
