package view

import ui "github.com/gizak/termui/v3"

// SingleScreenView allows one grid layout to be drawn on screen at a time.
type SingleScreenView struct {
	// screens is the slice of screens this view has stored
	screens []*ui.Grid
	// currentScreen stores the index of the active screen which will be rendered
	currentScreen uint
}
