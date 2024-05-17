package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDBPing(t *testing.T) {
	db, err := SQLiteConn.Conn(":memory:")
	assert.NoError(t, err)

	err = db.Ping()
	assert.NoError(t, err)
}
