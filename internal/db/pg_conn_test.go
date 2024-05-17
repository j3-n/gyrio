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
			Name:  `full args but all ""`,
			Data:  []interface{}{"", "", "", "", ""},
			Fails: true,
		},
		{
			Name:  "only addr test",
			Data:  []interface{}{"addr", "", "", "", ""},
			Fails: true,
		},
		{
			Name:  "addr, user test",
			Data:  []interface{}{"addr", "user", "", "", ""},
			Fails: true,
		},
		{
			Name:  "addr, user, password test",
			Data:  []interface{}{"addr", "user", "pass", "", ""},
			Fails: true,
		},
		{
			Name:  "addr, user, password, name test",
			Data:  []interface{}{"addr", "user", "pass", "name", ""},
			Fails: true,
		},
		{
			Name:  "addr, user, password, name, port test",
			Data:  []interface{}{"addr", "user", "pass", "name", "12345"},
			Fails: true,
		},
		{
			Name:  "wrong type test",
			Data:  []interface{}{1, "", "", "", ""},
			Fails: true,
		},
		{
			Name:  "wrong type test",
			Data:  []interface{}{1, 1, 1, 1, 1},
			Fails: true,
		},
	}

	c := pgConn{}

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

	// TODO: test real conn
}
