# testutils

## MySQLBox

Creates a ready to use MySQL server running in a Docker container that can be
used for testing.

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
	t.Cleanup(box.StopFunc())
}
```

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