package db

import (
	"errors"

	"gorm.io/gorm"
)

type PgArgs struct {
}

type PgConn struct {
	db *gorm.DB
}

func (c *PgConn) Connect(args PgArgs) error {
	return nil
}

func (c *PgConn) Close() error {
	return nil
}

func (c *PgConn) DB() (*gorm.DB, error) {
	if c.db == nil {
		return nil, errors.New("error, nil db connection")
	}

	return nil, nil
}
