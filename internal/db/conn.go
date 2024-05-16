package db

import "gorm.io/gorm"

type Connector interface {
	Connect(args any) error
	Close() error
	DB() (*gorm.DB, error)
}
