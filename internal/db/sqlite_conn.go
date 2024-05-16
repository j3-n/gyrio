package db

import (
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SqliteArgs struct {
	file string
}

type SqliteConn struct {
	db *gorm.DB
}

func (c *SqliteConn) Connect(args SqliteArgs) error {
	db, err := gorm.Open(sqlite.Open(args.file), &gorm.Config{})
	if err != nil {
		return err
	}

	c.db = db

	return nil
}

func (c *SqliteConn) Close() error {
	if c.db == nil {
		return errors.New("error, nil db conn")
	}

	db, err := c.db.DB()
	if err != nil {
		return err
	}

	err = db.Close()
	if err != nil {
		return err
	}

	return nil
}

func (c *SqliteConn) DB() (*gorm.DB, error) {
	if c.db == nil {
		return nil, errors.New("error, nil db conn")
	}

	return c.db, nil
}
