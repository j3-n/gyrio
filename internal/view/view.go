package view

import (
	ui "github.com/gizak/termui/v3"
)

// A View defines a layout that can be displayed in the terminal.
type View interface {
	// Render returns a Grid layout to be rendered to the screen.
	Render() ui.Grid
}
