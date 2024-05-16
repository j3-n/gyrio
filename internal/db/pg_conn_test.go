package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestPgConnConnect(t *testing.T) {
	c := PgConn{}
	err := c.Connect(PgArgs{})
	assert.NoError(t, err)
}

func TestPgConnClose(t *testing.T) {
	c := PgConn{}
	err := c.Close()
	assert.NoError(t, err)
}

func TestPgConnDB(t *testing.T) {
	c := PgConn{}
	_, err := c.DB()
	assert.Error(t, err)

	c.db = &gorm.DB{}
	_, err = c.DB()
	assert.NoError(t, err)
}
