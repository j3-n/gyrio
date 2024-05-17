package db

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyConn(t *testing.T) {
	c := emptyConn{}

	assert.True(t, reflect.TypeOf(c).Implements(reflect.TypeOf((*Connector)(nil)).Elem()))
}

func TestEmptyConnConn(t *testing.T) {
	c := emptyConn{}

	_, err := c.Conn()
	assert.NoError(t, err)
}
