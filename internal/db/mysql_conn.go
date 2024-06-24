package db

import (
	"errors"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mysqlConn struct{}

func (c mysqlConn) Conn(args ...interface{}) (*DB, error) {
	if len(args) != 5 {
		return nil, errors.New("error, invalid argument length for postgres connection")
	}

	user, ok := args[0].(string)
	if !ok || user == "" {
		return nil, errors.New("error, invalid database user given")
	}

	pass, ok := args[1].(string)
	if !ok || pass == "" {
		return nil, errors.New("error, invalid database password given")
	}

	addr, ok := args[2].(string)
	if !ok || addr == "" {
		return nil, errors.New("error, invalid database address given")
	}

	port, ok := args[3].(string)
	if !ok || port == "" {
		return nil, errors.New("error, invalid database port given")
	}

	name, ok := args[4].(string)
	if !ok || name == "" {
		return nil, errors.New("error, invalid database name given")
	}

	// user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, addr, port, name,
	)

	fmt.Println(dsn)

	gdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db := &DB{db: gdb, dbType: MySQL}

	return db, nil
}
