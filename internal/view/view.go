package view

import (
	ui "github.com/gizak/termui/v3"
	"github.com/j3-n/gyrio/internal/components"
)

// A View defines a layout that can be displayed in the terminal.
type View interface {
	// Render returns a Grid layout to be rendered to the screen.
	Render() *ui.Grid
	// KeyboardEvent sends a keyboard event to the the view
	KeyboardEvent(e *ui.Event)
	// Resize should be called with the terminal dimensions whenever it is resized.
	Resize(w int, height int)
	// ToolbarCallback should return the toolbar for the view.
	GetToolbar() []components.ToolbarEntry
}
