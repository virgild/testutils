package testutils

import (
	"database/sql"
	"errors"
	"fmt"
	"testing"

	"github.com/jmoiron/sqlx"
)

type DataHelper struct {
	DB *sqlx.DB
}

func (d *DataHelper) RowWithID(t *testing.T, table string, id string) map[string]interface{} {
	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = ?`, table)
	row := d.DB.QueryRowx(query, id)
	if row.Err() != nil {
		t.Fatal(row.Err())
		return nil
	}

	data := map[string]interface{}{}
	err := row.MapScan(data)
	if err != nil {
		t.Fatal(err)
		return nil
	}

	return data
}

func (d *DataHelper) RowsWhere(t *testing.T, table string, column string, value interface{}) []map[string]interface{} {
	query := fmt.Sprintf(`SELECT * FROM %s WHERE %s = ?`, table, column)
	rows, err := d.DB.Queryx(query, value)
	if errors.Is(err, sql.ErrNoRows) {
		return []map[string]interface{}{}
	}
	if err != nil {
		t.Fatal(err)
	}

	var dataRows []map[string]interface{}
	for {
		if !rows.Next() {
			if rows.Err() != nil {
				t.Fatal(rows.Err())
				return nil
			}

			break
		}

		data := map[string]interface{}{}
		err := rows.MapScan(data)
		if err != nil {
			t.Fatal(err)
			return nil
		}

		dataRows = append(dataRows, data)
	}

	return dataRows
}
