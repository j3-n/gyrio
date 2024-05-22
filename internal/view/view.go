package view

import (
	ui "github.com/gizak/termui/v3"
)

// A View defines a layout that can be displayed in the terminal.
type View interface {
	ui.Drawable

	// KeyboardEvent sends a keyboard event to the the view
	KeyboardEvent(*ui.Event)
}
