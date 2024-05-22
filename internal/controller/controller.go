package controller

import (
	ui "github.com/gizak/termui/v3"
	"github.com/j3-n/gyrio/internal/components"
	"github.com/j3-n/gyrio/internal/view"
)

// Controller handles communication between the user view and the database.
// Also handles rendering the final UI layout and initial direction of user input.
type Controller struct {
	// View is the view rendered by this controller
	View view.View
	// Toolbar is the toolbar rendered under the view
	Toolbar components.Toolbar
}

// Start begins the polling loop for the Controller.
func (c *Controller) Start() {
	// Set initial view size
	w, h := ui.TerminalDimensions()
	c.View.SetRect(0, 0, w, h)

	c.Render()

	events := ui.PollEvents()
	for e := range events {
		// Handle UI event
		if e.Type == ui.KeyboardEvent {
			switch e.ID {
			case "<C-c>", "q":
				// Quit
				return
			default:
				c.View.KeyboardEvent(&e)
			}
			// Redraw view
			c.Render()
		} else if e.Type == ui.ResizeEvent {
			// Resize view
			payload := e.Payload.(ui.Resize)
			c.View.SetRect(0, 0, payload.Width, payload.Height)
			c.Render()
		}
	}
}

// Render handles the rendering of the complete UI to the screen through TermUI.
func (c *Controller) Render() {
	ui.Clear()
	ui.Render(c.View)
}
