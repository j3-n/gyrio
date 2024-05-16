package db

import (
	"errors"

	"gorm.io/gorm"
)

type SqliteArgs struct {
	File string
}

type SqliteConn struct {
	db *gorm.DB
}

func (c *SqliteConn) Connect(args SqliteArgs) error {
	return nil
}

func (c *SqliteConn) Close() error {
	return nil
}

func (c *SqliteConn) DB() (*gorm.DB, error) {
	if c.db == nil {
		return nil, errors.New("error, nil db conn")
	}

	return c.db, nil
}
