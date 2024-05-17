package db

import (
	"errors"

	"gorm.io/gorm"
)

type DBType int

const (
	Empty    DBType = iota // 0
	SQLite                 // 1
	Postgres               // 2
	MySQL                  // 3
)

type DB struct {
	db *gorm.DB

	dbType DBType
}

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

func (d *DB) DB() *gorm.DB {
	return d.db
}

func (d *DB) DBType() DBType {
	return d.dbType
}

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
	err := d.db.Raw(cmd).Scan(&tables).Error
	if err != nil {
		return nil, err
	}

	return tables, nil
}

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
