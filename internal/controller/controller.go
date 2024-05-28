package controller

import (
	ui "github.com/gizak/termui/v3"
	"github.com/j3-n/gyrio/internal/view"
)

// Start begins the polling loop for the Controller.
func Start() {
	// Set initial view size
	w, h := ui.TerminalDimensions()
	view.StateView.SetRect(0, 0, w, h)

	Render()

	events := ui.PollEvents()
	for e := range events {
		// Handle UI event
		if e.Type == ui.KeyboardEvent {
			switch e.ID {
			case "<C-c>":
				// Quit
				return
			default:
				view.StateView.KeyboardEvent(&e)
			}
			// Redraw view
			Render()
		} else if e.Type == ui.ResizeEvent {
			// Resize view
			payload := e.Payload.(ui.Resize)
			view.StateView.SetRect(0, 0, payload.Width, payload.Height)
			Render()
		}
	}
}

// Render handles the rendering of the complete UI to the screen through TermUI.
func Render() {
	ui.Clear()
	ui.Render(view.StateView)
}
