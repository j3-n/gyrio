package db

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPgConn(t *testing.T) {
	c := pgConn{}

	assert.True(t, reflect.TypeOf(c).Implements(reflect.TypeOf((*Connector)(nil)).Elem()))
}

func TestPgConnConn(t *testing.T) {
	c := pgConn{}

	_, err := c.Conn()
	assert.Error(t, err)

	_, err = c.Conn("", "", "", "", "")
	assert.Error(t, err)

	_, err = c.Conn("addr", "", "", "", "")
	assert.Error(t, err)

	_, err = c.Conn("addr", "user", "", "", "")
	assert.Error(t, err)

	_, err = c.Conn("addr", "user", "pass", "", "")
	assert.Error(t, err)

	_, err = c.Conn("addr", "user", "pass", "name", "")
	assert.Error(t, err)

	_, err = c.Conn("addr", "user", "pass", "name", "12345")
	assert.Error(t, err)

	// TODO: test real conn
}
