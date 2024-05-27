package view

import (
	ui "github.com/gizak/termui/v3"
	"github.com/j3-n/gyrio/internal/components"
)

// A View defines a layout that can be displayed in the terminal.
type View interface {
	ui.Drawable

	// KeyboardEvent sends a keyboard event to the the view
	KeyboardEvent(*ui.Event)
}

// StateView changes depending on what View is currently being used
var StateView View

// Views that can be used by the application
var (
	AppView      View
	MenuView     View
	SettingsView View
)

// Init must be called before any views are used
func Init() {
	initAppView()
	StateView = AppView
}

func initAppView() {
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

	input := components.NewInput()
	input.Title = "Input Box"
	input.Border = true
	grid3 := ui.NewGrid()
	grid3.SetRect(0, 0, w, h)
	grid3.Set(
		ui.NewRow(1.0/1.5,
			ui.NewCol(1.0, table),
		),
		ui.NewRow(1.0/3,
			ui.NewCol(1.0, input),
		),
	)

	AppView = NewApplicationView([]*Layout{
		{
			Grid: grid3,
			NavLayout: [][]components.Interactable{
				{table},
				{input},
			},
		},
	})
}
