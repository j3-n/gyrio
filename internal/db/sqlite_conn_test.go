package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestSqliteConnConnect(t *testing.T) {
	c := SqliteConn{}
	err := c.Connect(SqliteArgs{})
	assert.NoError(t, err)
}

func TestSqliteConnClose(t *testing.T) {
	c := SqliteConn{}
	err := c.Close()
	assert.NoError(t, err)
}

func TestSqliteConnDB(t *testing.T) {
	c := SqliteConn{}
	_, err := c.DB()
	assert.Error(t, err)

	c.db = &gorm.DB{}
	_, err = c.DB()
	assert.NoError(t, err)
}
