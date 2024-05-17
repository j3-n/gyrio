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
	testData := []struct {
		Name  string
		Data  []interface{}
		Fails bool
	}{
		{
			Name:  "empty args",
			Data:  []interface{}{},
			Fails: true,
		},
		{
			Name:  "empty file",
			Data:  []interface{}{""},
			Fails: true,
		},
		{
			Name:  "wrong type",
			Data:  []interface{}{1},
			Fails: true,
		},
		{
			Name:  "valid in memory db",
			Data:  []interface{}{":memory:"},
			Fails: false,
		},
	}

	c := sqliteConn{}

	for _, test := range testData {
		t.Run(test.Name, func(t *testing.T) {
			_, err := c.Conn(test.Data...)
			if test.Fails {
				assert.Error(t, err, "failed to connect")
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
