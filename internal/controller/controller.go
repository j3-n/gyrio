package controller

import (
	ui "github.com/gizak/termui/v3"
	"github.com/j3-n/gyrio/internal/view"
)

// Controller handles communication between the user view and the database.
// Also handles rendering the final UI layout and initial direction of user input.
type Controller struct {
	// View is the view rendered by this controller
	View view.View
}

// Start begins the polling loop for the Controller.
func (c *Controller) Start() {
	ui.Render(c.View.Render())

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
			ui.Render(c.View.Render())
		}
	}
}
