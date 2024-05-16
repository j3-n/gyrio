package main

import (
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/j3-n/gyrio/internal/components"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("Error initialising termui: %v", err)
	}
	defer ui.Close()

	table := components.NewTableView()
	table.Title = "Customers"

	grid := ui.NewGrid()
	w, h := ui.TerminalDimensions()
	grid.SetRect(0, 0, w, h)

	grid.Set(ui.NewRow(1.0,
		ui.NewCol(1.0, table),
	),
	)

	ui.Render(grid)

	events := ui.PollEvents()
	for {
		e := <-events
		switch e.ID {
		case "q", "<C-c>":
			return
		case "<Resize>":
			resize := e.Payload.(ui.Resize)
			grid.SetRect(0, 0, resize.Width, resize.Height)
			ui.Clear()
			ui.Render(grid)
		}
	}
}
