package view

import (
	"testing"

	ui "github.com/gizak/termui/v3"
	"github.com/j3-n/gyrio/internal/components"
	"github.com/stretchr/testify/assert"
)

func TestNavLayout(t *testing.T) {
	layout := Layout{
		NavLayout: [][]components.Interactable{
			{}, // 0
			{components.NewInput(), components.NewInput()}, // 2
			{components.NewInput()},                        // 1
			{components.NewInput(), components.NewInput(), components.NewInput(), components.NewInput()}, // 4
		},
	}

	assert.Equal(t, 0, layout.row, "Test that navigation begins on row 0")
	assert.Equal(t, 0, layout.col, "Test that navigation begins on col 0")
	layout.NavX(1)
	assert.Equal(t, 0, layout.col)
	layout.NavX(-1)
	assert.Equal(t, 0, layout.col)
	layout.NavY(-1)
	assert.Equal(t, 3, layout.row, "Test that vertical wrapping works correctly")
	layout.NavY(2)
	assert.Equal(t, 1, layout.row)
	layout.NavX(1)
	assert.Equal(t, 1, layout.col)
	layout.NavX(1)
	assert.Equal(t, 0, layout.col)
	layout.NavY(2)
	layout.NavX(3)
	assert.Equal(t, 3, layout.col)
	layout.NavY(-1)
	assert.Equal(t, 0, layout.col)
}

func TestLayout(t *testing.T) {
	grid := ui.NewGrid()
	input := components.NewInput()
	grid.Set(ui.NewRow(1.0, ui.NewCol(1.0, input)))

	layout := Layout{
		Grid:      grid,
		NavLayout: [][]components.Interactable{},
	}
	// Test that this does not create a runtime error since layout has no nav components
	layout.KeyboardEvent(&ui.Event{})
}
