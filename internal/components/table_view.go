package components

import (
	"errors"
	"image"

	ui "github.com/gizak/termui/v3"
)

var ErrColumnCountMismatch = errors.New("number of columns in each row must be constant")

// TableView is a termui component for drawing SQL tables.
type TableView struct {
	ui.Block

	// Columns is a slice of column titles
	Columns []string
	// Data is the data to populate the table with
	Data [][]string
	// Style to use for column titles
	ColumnTitleStyle ui.Style
	// Style to use for column borders
	ColumnBorderStyle ui.Style
}

// NewTableView initialises and returns a new empty TableView widget.
func NewTableView() *TableView {
	return &TableView{
		Block:             *ui.NewBlock(),
		Columns:           []string{},
		Data:              [][]string{},
		ColumnTitleStyle:  ui.NewStyle(ui.ColorWhite),
		ColumnBorderStyle: ui.NewStyle(ui.ColorWhite),
	}
}

func (t *TableView) Draw(buf *ui.Buffer) {
	t.Block.Draw(buf)

	// Draw column titles
	i := 1
	p := t.Inner.Min
	for j, col := range t.Columns {
		// Draw column border edges
		switch j {
		case 0:
			buf.SetCell(ui.NewCell(ui.TOP_LEFT, t.ColumnBorderStyle), p)
			buf.SetCell(ui.NewCell(ui.VERTICAL_LINE, t.ColumnBorderStyle), p.Add(image.Pt(0, 1)))
			buf.SetCell(ui.NewCell(ui.BOTTOM_LEFT, t.ColumnBorderStyle), p.Add(image.Pt(0, 2)))
		default:
			buf.SetCell(ui.NewCell(ui.HORIZONTAL_DOWN, t.ColumnBorderStyle), p.Add(image.Pt(i-1, 0)))
			buf.SetCell(ui.NewCell(ui.VERTICAL_LINE, t.ColumnBorderStyle), p.Add(image.Pt(i-1, 1)))
			buf.SetCell(ui.NewCell(ui.HORIZONTAL_UP, t.ColumnBorderStyle), p.Add(image.Pt(i-1, 2)))
		}
		// Draw titles
		for _, c := range col {
			// Draw text
			cell := ui.NewCell(c, t.ColumnTitleStyle)
			buf.SetCell(cell, t.Inner.Min.Add(image.Pt(i, 1)))
			// Draw top and bottom borders
			buf.SetCell(ui.NewCell(ui.HORIZONTAL_LINE, t.ColumnBorderStyle), t.Inner.Min.Add(image.Pt(i, 0)))
			buf.SetCell(ui.NewCell(ui.HORIZONTAL_LINE, t.ColumnBorderStyle), t.Inner.Min.Add(image.Pt(i, 2)))
			i += 1
		}
		i += 1
	}
	// Draw last column border
	buf.SetCell(ui.NewCell(ui.TOP_RIGHT, t.ColumnBorderStyle), p.Add(image.Pt(i-1, 0)))
	buf.SetCell(ui.NewCell(ui.VERTICAL_LINE, t.ColumnBorderStyle), p.Add(image.Pt(i-1, 1)))
	buf.SetCell(ui.NewCell(ui.BOTTOM_RIGHT, t.ColumnBorderStyle), p.Add(image.Pt(i-1, 2)))
}

// computeWidths computes a slice of the maximum widths of each column for a set of data
func computeWidths(data [][]string) ([]int, error) {
	result := []int{}
	// Initialise columns
	for range data[0] {
		result = append(result, 0)
	}
	for _, r := range data {
		if len(r) != len(result) {
			return []int{}, ErrColumnCountMismatch
		}
		for i, c := range r {
			result[i] = max(len(c), result[i])
		}
	}
	return result, nil
}
