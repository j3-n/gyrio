package view

import (
	"fmt"

	ui "github.com/gizak/termui/v3"
	"github.com/j3-n/gyrio/internal/components"
)

// SingleScreenView allows one grid layout to be drawn on screen at a time.
type SingleScreenView struct {
	// screens is the slice of screens this view has stored
	screens []*ui.Grid
	// currentScreen stores the index of the active screen which will be rendered
	currentScreen int
	// toolbar stores the toolbar component of the view
	pageCounter components.ToolbarEntry
}

// NewSingleScreenView creates a new view with the given set of available screens.
func NewSingleScreenView(screens []*ui.Grid) *SingleScreenView {
	// Initialise toolbar component
	pageCounter := components.ToolbarEntry{
		Key:  "TAB",
		Text: fmt.Sprintf("(%d/%d) ➤", 1, len(screens)),
	}
	return &SingleScreenView{
		screens:       screens,
		currentScreen: 0,
		pageCounter:   pageCounter,
	}
}

// GetToolbar returns the toolbar for the current screen and view overall.
func (v *SingleScreenView) GetToolbar() []components.ToolbarEntry {
	// Update the tooltip text for page count
	v.pageCounter.Text = fmt.Sprintf("(%d/%d) ➤", v.currentScreen+1, len(v.screens))
	return []components.ToolbarEntry{v.pageCounter}
}

// Render returns the currently active screen.
func (v *SingleScreenView) Render() *ui.Grid {
	return v.screens[v.currentScreen]
}

// Resize resizes all screens.
func (v *SingleScreenView) Resize(w int, h int) {
	for _, screen := range v.screens {
		screen.SetRect(0, 0, w, h)
	}
}

// KeyboardEvent handles keyboard inputs to this view. If it is not a control input
// it will be forwarded to the active widget.
func (v *SingleScreenView) KeyboardEvent(e *ui.Event) {
	switch e.ID {
	case "<Tab>":
		// Cycle along one screen layout
		v.currentScreen = (v.currentScreen + 1) % len(v.screens)
	default:
		return
	}
}
