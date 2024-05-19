package view

import ui "github.com/gizak/termui/v3"

// SingleScreenView allows one grid layout to be drawn on screen at a time.
type SingleScreenView struct {
	// screens is the slice of screens this view has stored
	screens []*ui.Grid
	// currentScreen stores the index of the active screen which will be rendered
	currentScreen uint
}

// Render returns the currently active screen.
func (v *SingleScreenView) Render() *ui.Grid {
	return v.screens[v.currentScreen]
}

// KeyboardEvent handles keyboard inputs to this view. If it is not a control input
// it will be forwarded to the active widget.
func (v *SingleScreenView) KeyboardEvent(e *ui.Event) {

}
