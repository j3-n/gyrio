package db

import (
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type sqliteConn struct{}

func (c sqliteConn) Conn(args ...interface{}) (*DB, error) {
	if len(args) != 1 {
		return nil, errors.New("error, please give database file")
	}

	file, ok := args[0].(string)
	if !ok || file == "" {
		return nil, errors.New("error, invalid database file given")
	}

	gdb, err := gorm.Open(sqlite.Open(file), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db := &DB{DB: gdb}

	return db, nil
}
