package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDBPing(t *testing.T) {
	db, err := SQLiteConn.Conn(":memory:")
	assert.NoError(t, err)

	err = db.Ping()
	assert.NoError(t, err)
}

func TestDBDB(t *testing.T) {
	db, err := SQLiteConn.Conn(":memory:")
	assert.NoError(t, err)
	assert.Equal(t, db.db, db.DB())
}

func TestDBDBType(t *testing.T) {
	db, err := SQLiteConn.Conn(":memory:")
	assert.NoError(t, err)
	assert.Equal(t, db.dbType, db.DBType())
	assert.Equal(t, SQLite, db.DBType())
}

func TestDBTables(t *testing.T) {
	testData := []struct {
		Name      string
		Connector Connector
		Address   []interface{}
		Tables    int
		Fails     bool
	}{
		{
			Name:      "sqlite db, no tables",
			Connector: SQLiteConn,
			Address:   []interface{}{":memory:"},
			Tables:    0,
			Fails:     false,
		},
		{
			Name:      "empty conn, no tables",
			Connector: EmptyConn,
			Address:   []interface{}{},
			Tables:    0,
			Fails:     true,
		},
	}

	for _, test := range testData {
		t.Run(test.Name, func(t *testing.T) {
			db, err := test.Connector.Conn(test.Address...)
			assert.NoError(t, err)

			tables, err := db.Tables()
			if test.Fails {
				assert.Error(t, err, "failed to connect")
			} else {
				assert.NoError(t, err)
			}
			assert.Len(t, tables, test.Tables)
		})
	}
}

func TestDBClose(t *testing.T) {
	db, err := SQLiteConn.Conn(":memory:")
	assert.NoError(t, err)
	err = db.Close()
	assert.NoError(t, err)
}
