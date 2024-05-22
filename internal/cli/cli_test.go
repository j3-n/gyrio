package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	cli := New()
	assert.NotNil(t, cli)
}
