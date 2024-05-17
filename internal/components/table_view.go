package components

import (
	"image"

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

	// Draw column titles
	i := 1
	for _, col := range t.Columns {
		for _, c := range col {
			cell := ui.NewCell(c, ui.NewStyle(ui.ColorWhite))
			p := image.Pt(t.Inner.Min.X+i, t.Inner.Min.Y+1)
			buf.SetCell(cell, p)
			i += 1
		}
		i += 1
	}
}
