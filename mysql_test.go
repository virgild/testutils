package testutils

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestPanicRecoverCleanup(t *testing.T) {
	b, err := StartMySQLBox(&MySQLBoxConfig{})
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		recover()
		b.StopFunc()()
	}()

	panic("panic!")
}

func TestMySQLBoxWithInitialSchema(t *testing.T) {
	schemaFile, err := os.Open("testdata/schema.sql")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		schemaFile.Close()
	}()

	b, err := StartMySQLBox(&MySQLBoxConfig{
		InitialSchemaSQL: schemaFile,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(b.StopFunc())

	for n := 0; n < 100; n++ {
		query := "INSERT INTO users (id, email, created_at, updated_at) VALUES (?, ?, ?, ?)"
		now := time.Now()
		_, err := b.DB().Exec(query, fmt.Sprintf("U-%d", n), fmt.Sprintf("user%d@example.com", n), now, now)
		if err != nil {
			t.Error(err)
		}
	}

	var count uint64
	err = b.DB().Get(&count, "SELECT COUNT(*) FROM users")
	if err != nil {
		t.Error(err)
	}
	if count != 100 {
		t.Error("count does not match")
	}
}
