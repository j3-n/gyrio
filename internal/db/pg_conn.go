package db

import (
	"errors"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type pgConn struct{}

func (c pgConn) Conn(args ...interface{}) (*DB, error) {
	if len(args) != 5 {
		return nil, errors.New("error, invalid argument length for postgres connection")
	}

	addr, ok := args[0].(string)
	if !ok || addr == "" {
		return nil, errors.New("error, invalid database address given")
	}

	user, ok := args[1].(string)
	if !ok || user == "" {
		return nil, errors.New("error, invalid database user given")
	}

	pass, ok := args[2].(string)
	if !ok || pass == "" {
		return nil, errors.New("error, invalid database password given")
	}

	name, ok := args[3].(string)
	if !ok || name == "" {
		return nil, errors.New("error, invalid database name given")
	}

	port, ok := args[4].(string)
	if !ok || port == "" {
		return nil, errors.New("error, invalid database port given")
	}

	dsn := fmt.Sprintf(`
		host=%s 
		user=%s 
		password=%s 
		dbname=%s 
		port=%s 
		sslmode=disable 
		TimeZone=Europe/London`,
		addr,
		user,
		pass,
		name,
		port,
	)

	gdb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db := &DB{db: gdb, dbType: Postgres}

	return db, nil
}
