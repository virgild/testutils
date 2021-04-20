# testutils

testutils contains packages that helps in testing Go programs.

## MySQLBox

![MySQLBox logo](https://github.com/virgild/testutils/blob/main/static/logo.png?raw=true)

MySQLBox creates a ready to use MySQL server running in a Docker container that can be
used in Go tests. The `Start()` function returns a `MySQLBox` that has a container running MySQL server. 
It has a `Stop()` function that stops the container when called. The `DB()` function returns a connected 
`sql.DB` object that can be used to send queries to MySQL. 

```go
package mytests

import (
	"testing"

	"github.com/virgild/testutils/mysqlbox"
)

func TestMyCode(t *testing.T) {
	// Start MySQL container
	box, err := mysqlbox.Start(&mysqlbox.Config{})
	if err != nil {
	    t.Fatal(err)
	}

	// Register the stop func to stop the container after the test.
	t.Cleanup(func() {
	    err := b.Stop()
	    if err != nil {
	        t.Fatal(err)
	    }
	})
	
	// Use the sql.DB object to query the database.
	box.DB().QueryRow("SELECT NOW()")
	
	var now time.Time
	err = row.Scan(&row)
	if err != nil {
	    t.Error(err)
	}
	
	if now.IsZero() {
	    t.Error("now is zero")
	}
}
```

### Other Features

* Initial script
* Clean tables

### With initial schema

From file:

```go
schemaFile, err := os.Open("testdata/schema.sql")
if err != nil {
    t.Fatal(err)
}
defer schemaFile.Close()

box, err = mysqlbox.Start(&Config{
    InitialSchema: InitialSchemaFromReader(schemaFile),
}
if err != nil {
    t.Fatal(err)
}

t.Cleanup(b.StopFunc())
```