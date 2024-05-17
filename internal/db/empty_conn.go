package db

import "gorm.io/gorm"

type emptyConn struct{}

func (c emptyConn) Conn(args ...interface{}) (*DB, error) {
	db := &DB{DB: &gorm.DB{}}
	return db, nil
}
