package main

import (
	"log"
	"os"

	ui "github.com/gizak/termui/v3"
	"github.com/j3-n/gyrio/internal/cli"
	"github.com/j3-n/gyrio/internal/components"
)

func main() {
	app := cli.New()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

	if err := ui.Init(); err != nil {
		log.Fatalf("Error initialising termui: %v", err)
	}
	defer ui.Close()

	table := components.NewTableView()
	table.Title = "Customers"
	table.TextAlignment = ui.AlignLeft
	table.Columns = []string{"CustomerID", "Name", "Email", "Phone", "Address"}
	table.Data = [][]string{
		{"1", "Jacob Padley", "test@mail.com", "123456789", "1 cool ln."},
		{"2", "Joseph Beck", "silly@billy.ac.uk", "123456789", "3 feet ave."},
		{"3", "Bosh Jorges", "josh@lockheedmartin.com", "123456789", "4 toes way"},
	}

	grid := ui.NewGrid()
	w, h := ui.TerminalDimensions()
	grid.SetRect(0, 0, w, h)

	grid.Set(ui.NewRow(1.0/2,
		ui.NewCol(1.0/4, table),
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
		case "<Down>":
			table.ScrollDown()
			ui.Render(grid)
		case "<Up>":
			table.ScrollUp()
			ui.Render(grid)
		case "<Left>":
			table.ScrollLeft()
			ui.Render(grid)
		case "<Right>":
			table.ScrollRight()
			ui.Render(grid)
		}
	}
}
