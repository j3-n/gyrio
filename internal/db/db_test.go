package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Mock struct {
	ID    int64  `gorm:"primarykey"`
	Value string `gorm:"type:text"`
}

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
			Data:      nil,
			Before: func(db *DB) error {
				if err := db.DB().Migrator().CreateTable(&Mock{}); err != nil {
					return err
				}

				return nil
			},
			Fails: false,
		},
		{
			Name:      "sqlite db, one entry",
			Connector: SQLiteConn,
			Table:     "mocks",
			Address:   []interface{}{":memory:"},
			Data:      []map[string]any{{"id": int64(1), "value": "hello"}},
			Before: func(db *DB) error {
				if err := db.DB().Migrator().CreateTable(&Mock{}); err != nil {
					return err
				}

				if err := db.DB().Table("mocks").Create(&Mock{Value: "hello"}).Error; err != nil {
					return err
				}

				return nil
			},
			Fails: false,
		},
		{
			Name:      "sqlite db, many entry",
			Connector: SQLiteConn,
			Table:     "mocks",
			Address:   []interface{}{":memory:"},
			Data:      []map[string]any{{"id": int64(1), "value": "hello"}, {"id": int64(2), "value": "hello"}},
			Before: func(db *DB) error {
				if err := db.DB().Migrator().CreateTable(&Mock{}); err != nil {
					return err
				}

				if err := db.DB().Table("mocks").Create(&Mock{Value: "hello"}).Error; err != nil {
					return err
				}

				if err := db.DB().Table("mocks").Create(&Mock{Value: "hello"}).Error; err != nil {
					return err
				}

				return nil
			},
			Fails: false,
		},
	}

	for _, test := range testData {
		t.Run(test.Name, func(t *testing.T) {
			db, err := test.Connector.Conn(test.Address...)
			assert.NoError(t, err)
			err = test.Before(db)
			assert.NoError(t, err)

			data, err := db.List(test.Table)
			if test.Fails {
				assert.Error(t, err, "failed to execute query")
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, test.Data, data)
		})
	}
}

func TestDBRead(t *testing.T) {
	testData := []struct {
		Name      string
		Connector Connector
		Table     string
		Address   []interface{}
		Data      map[string]any
		Before    func(db *DB) error
		Fails     bool
	}{
		{
			Name:      "sqlite db, no entries",
			Connector: SQLiteConn,
			Table:     "table",
			Address:   []interface{}{":memory:"},
			Data:      nil,
			Before:    func(db *DB) error { return nil },
			Fails:     true,
		},
		{
			Name:      "sqlite db, no entries",
			Connector: SQLiteConn,
			Table:     "mocks",
			Address:   []interface{}{":memory:"},
			Data:      nil,
			Before: func(db *DB) error {
				if err := db.DB().Migrator().CreateTable(&Mock{}); err != nil {
					return err
				}

				return nil
			},
			Fails: false,
		},
		{
			Name:      "sqlite db, one read with one entry",
			Connector: SQLiteConn,
			Table:     "mocks",
			Address:   []interface{}{":memory:"},
			Data:      map[string]any{"id": int64(1), "value": "hello"},
			Before: func(db *DB) error {
				if err := db.DB().Migrator().CreateTable(&Mock{}); err != nil {
					return err
				}

				if err := db.DB().Table("mocks").Create(&Mock{Value: "hello"}).Error; err != nil {
					return err
				}

				return nil
			},
			Fails: false,
		},
		{
			Name:      "sqlite db, one read with many entry",
			Connector: SQLiteConn,
			Table:     "mocks",
			Address:   []interface{}{":memory:"},
			Data:      map[string]any{"id": int64(1), "value": "hello"},
			Before: func(db *DB) error {
				if err := db.DB().Migrator().CreateTable(&Mock{}); err != nil {
					return err
				}

				if err := db.DB().Table("mocks").Create(&Mock{Value: "hello"}).Error; err != nil {
					return err
				}

				if err := db.DB().Table("mocks").Create(&Mock{Value: "hello"}).Error; err != nil {
					return err
				}

				return nil
			},
			Fails: false,
		},
	}

	for _, test := range testData {
		t.Run(test.Name, func(t *testing.T) {
			db, err := test.Connector.Conn(test.Address...)
			assert.NoError(t, err)
			err = test.Before(db)
			assert.NoError(t, err)

			data, err := db.Read(test.Table, "id = ?", 1)
			if test.Fails {
				assert.Error(t, err, "failed to execute query")
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, test.Data, data)
		})
	}
}

func TestDBCreate(t *testing.T) {
	before := func(db *DB) error {
		err := db.DB().Migrator().CreateTable(&Mock{})
		if err != nil {
			return err
		}

		return nil
	}

	testData := []struct {
		Name      string
		Connector Connector
		Table     string
		Address   []interface{}
		Data      map[string]any
		Before    func(db *DB) error
		Fails     bool
	}{
		{
			Name:      "sqlite db, no entries",
			Connector: SQLiteConn,
			Table:     "mocks",
			Address:   []interface{}{":memory:"},
			Data:      map[string]any{},
			Before:    before,
			Fails:     false,
		},
		{
			Name:      "sqlite db, invalid entries",
			Connector: SQLiteConn,
			Table:     "mocks",
			Address:   []interface{}{":memory:"},
			Data:      map[string]any{"key": "key"},
			Before:    before,
			Fails:     true,
		},
		{
			Name:      "sqlite db, one valid entry",
			Connector: SQLiteConn,
			Table:     "mocks",
			Address:   []interface{}{":memory:"},
			Data:      map[string]any{"value": "value"},
			Before:    before,
			Fails:     false,
		},
		{
			Name:      "sqlite db, one valid entry, with id",
			Connector: SQLiteConn,
			Table:     "mocks",
			Address:   []interface{}{":memory:"},
			Data:      map[string]any{"id": 1, "value": "value"},
			Before:    before,
			Fails:     false,
		},
	}

	for _, test := range testData {
		t.Run(test.Name, func(t *testing.T) {
			db, err := test.Connector.Conn(test.Address...)
			assert.NoError(t, err)
			err = test.Before(db)
			assert.NoError(t, err)

			err = db.Create(test.Table, test.Data)
			if test.Fails {
				assert.Error(t, err, "failed to execute query")
				return
			}

			assert.NoError(t, err)
		})
	}
}

func TestDBClose(t *testing.T) {
	db, err := SQLiteConn.Conn(":memory:")
	assert.NoError(t, err)
	err = db.Close()
	assert.NoError(t, err)
}
