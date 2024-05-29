package cli

import (
	"testing"

	"github.com/j3-n/gyrio/internal/db"
	"github.com/stretchr/testify/assert"
)

func TestFetch(t *testing.T) {
	conn := db.New(db.SQLite)
	db, err := conn.Conn(":memory:")
	assert.NoError(t, err)

	_, err = fetch(db)
	assert.NoError(t, err)
}

func TestFormat(t *testing.T) {
	data := []map[string]any{
		{"id": "1"}, {"id": "2"},
	}

	keys, tbl := format(data)
	assert.Equal(t, []string{"id"}, keys)
	assert.Equal(t, [][]string{{"1"}, {"2"}}, tbl)
}
