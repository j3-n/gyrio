package db

import "gorm.io/gorm"

type emptyConn struct{}

func (c emptyConn) Conn(args ...interface{}) (*DB, error) {
	db := &DB{db: &gorm.DB{}, dbType: Empty}
	return db, nil
}
