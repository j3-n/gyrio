package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	ui "github.com/gizak/termui/v3"
	"github.com/j3-n/gyrio/internal/components"
	"github.com/j3-n/gyrio/internal/controller"
	"github.com/j3-n/gyrio/internal/view"
)

func main() {
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
		ui.NewCol(1.0/4, table),
		ui.NewCol(1.0/4, table),
		ui.NewCol(1.0/4, table),
	),
	)

	grid2 := ui.NewGrid()
	grid2.SetRect(0, 0, w, h)
	grid2.Set(ui.NewRow(1.0, ui.NewCol(1.0, table)))

	// Create a quit signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Create and start controller
	c := &controller.Controller{
		View: view.NewSingleScreenView([]*ui.Grid{
			grid,
			grid2,
		}),
		Toolbar: *components.NewToolbar(),
	}
	c.Start()
}
