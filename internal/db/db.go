package db

import "gorm.io/gorm"

type DB struct {
	DB *gorm.DB
}

func (d *DB) Ping() error {
	sql, err := d.DB.DB()
	if err != nil {
		return err
	}

	err = sql.Ping()
	if err != nil {
		return err
	}

	return nil
}
