package db

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSqliteConn(t *testing.T) {
	c := sqliteConn{}

	assert.True(t, reflect.TypeOf(c).Implements(reflect.TypeOf((*Connector)(nil)).Elem()))
}

func TestSqliteConnConn(t *testing.T) {
	c := sqliteConn{}

	_, err := c.Conn(":memory:")
	assert.NoError(t, err)

	_, err = c.Conn()
	assert.Error(t, err)

	_, err = c.Conn("")
	assert.Error(t, err)
}
