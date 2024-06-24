package db

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	c := New(Empty)
	assert.True(t, reflect.TypeOf(c).Implements(reflect.TypeOf((*Connector)(nil)).Elem()))
	assert.Equal(t, emptyConn{}, c)

	c = New(SQLite)
	assert.True(t, reflect.TypeOf(c).Implements(reflect.TypeOf((*Connector)(nil)).Elem()))
	assert.Equal(t, sqliteConn{}, c)

	c = New(Postgres)
	assert.True(t, reflect.TypeOf(c).Implements(reflect.TypeOf((*Connector)(nil)).Elem()))
	assert.Equal(t, pgConn{}, c)

	c = New(MySQL)
	assert.True(t, reflect.TypeOf(c).Implements(reflect.TypeOf((*Connector)(nil)).Elem()))
	assert.Equal(t, mysqlConn{}, c)

	c = New(Maria)
	assert.True(t, reflect.TypeOf(c).Implements(reflect.TypeOf((*Connector)(nil)).Elem()))
	assert.Equal(t, mariaConn{}, c)
}
