package db

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

// Stores the type of database connection.
// These are used for specific database queries.
type DBType int

const (
	Empty    DBType = iota // 0
	SQLite                 // 1
	Postgres               // 2
	MySQL                  // 3
)

// Stores the gorm database connection as well as the type of the database.
type DB struct {
	db     *gorm.DB
	dbType DBType
}

// Pings the database and errors if it fails
func (d *DB) Ping() error {
	sql, err := d.db.DB()
	if err != nil {
		return err
	}

	err = sql.Ping()
	if err != nil {
		return err
	}

	return nil
}

// Gives a pointer to the currently held gorm instance of the DB.
// Use this if you would like more specific queries.
func (d *DB) DB() *gorm.DB {
	return d.db
}

// Gives the type of the DB instance.
func (d *DB) DBType() DBType {
	return d.dbType
}

// Returns a string array of all table names from the connected database.
// Will use different queries based on the Database type, like PG, SQLite, etc.
func (d *DB) Tables() ([]string, error) {
	cmd := func() string {
		if d.dbType == SQLite {
			return "SELECT name FROM sqlite_master WHERE type='table'"
		} else if d.dbType == Postgres {
			return "SELECT tablename FROM pg_catalog.pg_tables WHERE schemaname='public'"
		} else if d.dbType == MySQL {
			return ""
		}

		return ""
	}()

	if cmd == "" {
		return nil, errors.New("error, database of incorrect type")
	}

	var tables []string
	res := d.db.Raw(cmd).Scan(&tables)
	err := res.Error
	if err != nil {
		return nil, err
	}

	return tables, nil
}

// Lists all of the entries from a given table to a an array of map[string]any.
// Also returns the error of the query.
// Use the args if you would like a more specific query, for example (..., "id = ?", 1).
func (d *DB) List(tbl string, args ...interface{}) ([]map[string]any, error) {
	var data []map[string]any
	res := d.db.Table(tbl).Find(&data, args...)
	err := res.Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Read the first item that fits the given criteria into a map[string]any.
// Also returns the error of the query.
// Use the args to make a more specific query, for example (..., "id = ?", 1).
func (d *DB) Read(tbl string, args ...interface{}) (map[string]any, error) {
	var data map[string]any
	res := d.db.Table(tbl).Find(&data, args...)
	log.Println(data)
	err := res.Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Create an item into the given table with the map[string]any.
// Returns any errors that occur during saving.
func (d *DB) Create(tbl string, obj map[string]any) error {
	res := d.db.Table(tbl).Create(&obj)
	err := res.Error
	if err != nil {
		return err
	}

	return nil
}

// Update an already existing item in the given table with the map[string]any.
// Returns any errors that occur during updating.
// The obj must have the primary key in order to update.
func (d *DB) Update(tbl string, obj map[string]any, args ...interface{}) error {
	res := d.db.Table(tbl).Save(&obj)
	err := res.Error
	if err != nil {
		return err
	}

	return nil
}

// Delete items from the database using the map[string]any or with a given table and query.
// Returns an error that occurs during deletion.
// Use the args to make a more specific query, for example (..., "id = ?", 1).
func (d *DB) Delete(tbl string, obj map[string]any, args ...interface{}) error {
	res := d.db.Table(tbl).Delete(&obj, args...)
	err := res.Error
	if err != nil {
		return err
	}

	return nil
}

// Check if an item exists within a given table, either using the map[string]any of obj or args.
// Returns a boolean, if it exists, and any errors that occur.
// Can use args, with something like (..., "id = ?", 1).
func (d *DB) Contains(tbl string, obj map[string]any, args ...interface{}) (bool, error) {
	var count int64
	res := d.db.Table(tbl).Find(&obj, args...).Count(&count)
	err := res.Error
	if err != nil {
		return false, err
	}

	return count == 0, err
}

// Closes a connection to the database instance.
// Should often defer the closing of the database after creating an instance of it, if in main.
func (d *DB) Close() error {
	sql, err := d.db.DB()
	if err != nil {
		return err
	}

	err = sql.Close()
	if err != nil {
		return err
	}

	return nil
}
