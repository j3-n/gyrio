package components

import (
	ui "github.com/gizak/termui/v3"
)

// TableView is a termui component for drawing SQL tables.
type TableView struct {
	ui.Block

	// Columns is a slice of column titles
	Columns []string
}

// NewTableView initialises and returns a new TableView widget.
func NewTableView() *TableView {
	return &TableView{
		Block: *ui.NewBlock(),
	}
}

func (t *TableView) Draw(buf *ui.Buffer) {
	t.Block.Draw(buf)
}
