package db

import (
	"reflect"
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
				return
			}

			assert.NoError(t, err)
			assert.Len(t, tables, test.Tables)
		})
	}
}

func TestDBList(t *testing.T) {
	type Mock struct {
		ID    uint   `gorm:"primarykey"`
		Value string `gorm:"type:text"`
	}

	testData := []struct {
		Name      string
		Connector Connector
		Table     string
		Address   []interface{}
		Data      []map[string]any
		Before    func(db *DB) error
		Fails     bool
	}{
		{
			Name:      "sqlite db, no entries",
			Connector: SQLiteConn,
			Table:     "mocks",
			Address:   []interface{}{":memory:"},
			Data:      []map[string]any{},
			Before: func(db *DB) error {
				db.DB().Migrator().CreateTable(&Mock{})

				return nil
			},
			Fails: true,
		},
		{
			Name:      "sqlite db, one entry",
			Connector: SQLiteConn,
			Table:     "mocks",
			Address:   []interface{}{":memory:"},
			Data: []map[string]any{
				{
					"ID":    1,
					"Value": "hello",
				},
			},
			Before: func(db *DB) error {
				db.DB().Migrator().CreateTable(&Mock{})

				db.DB().Table("mocks").Create(&Mock{Value: "hello"})

				return nil
			},
			Fails: true,
		},
		{
			Name:      "sqlite db, many entry",
			Connector: SQLiteConn,
			Table:     "mocks",
			Address:   []interface{}{":memory:"},
			Data: []map[string]any{
				{
					"ID":    1,
					"Value": "hello",
				},
				{
					"ID":    2,
					"Value": "hello",
				},
			},
			Before: func(db *DB) error {
				db.DB().Migrator().CreateTable(&Mock{})

				db.DB().Table("mocks").Create(&Mock{Value: "hello"})
				db.DB().Table("mocks").Create(&Mock{Value: "hello"})

				return nil
			},
			Fails: true,
		},
	}

	for _, test := range testData {
		t.Run(test.Name, func(t *testing.T) {
			db, err := test.Connector.Conn(test.Address...)
			assert.NoError(t, err)
			err = test.Before(db)
			assert.NoError(t, err)

			data, err := db.List("table")
			if test.Fails {
				assert.Error(t, err, "failed to connect")
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, test.Data, data)
			assert.True(t, reflect.DeepEqual(test.Data, data))
		})
	}
}

func TestDBClose(t *testing.T) {
	db, err := SQLiteConn.Conn(":memory:")
	assert.NoError(t, err)
	err = db.Close()
	assert.NoError(t, err)
}
