package view

import ui "github.com/gizak/termui/v3"

// SingleScreenView allows one grid layout to be drawn on screen at a time.
type SingleScreenView struct {
	// screens is the slice of screens this view has stored
	screens []*ui.Grid
	// currentScreen stores the index of the active screen which will be rendered
	currentScreen int
}

// NewSingleScreenView creates a new view with the given set of available screens.
func NewSingleScreenView(screens []*ui.Grid) *SingleScreenView {
	return &SingleScreenView{
		screens:       screens,
		currentScreen: 0,
	}
}

// Render returns the currently active screen.
func (v *SingleScreenView) Render() *ui.Grid {
	return v.screens[v.currentScreen]
}

// Resize resizes all screens
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
