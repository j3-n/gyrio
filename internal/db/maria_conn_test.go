package db

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMariaConn(t *testing.T) {
	c := mariaConn{}

	assert.True(t, reflect.TypeOf(c).Implements(reflect.TypeOf((*Connector)(nil)).Elem()))
}

func TestMariaConnConn(t *testing.T) {
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
			Data:  []interface{}{"user", "", "", "", ""},
			Fails: true,
		},
		{
			Name:  "user, pass test",
			Data:  []interface{}{"user", "pass", "", "", ""},
			Fails: true,
		},
		{
			Name:  "user, pass, addr test",
			Data:  []interface{}{"user", "pass", "addr", "", ""},
			Fails: true,
		},
		{
			Name:  "user, pass, addr, port test",
			Data:  []interface{}{"user", "pass", "addr", "12345", ""},
			Fails: true,
		},
		{
			Name:  "user, pass, addr, port, name test",
			Data:  []interface{}{"user", "pass", "addr", "12345", "name"},
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

	c := mariaConn{}

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
