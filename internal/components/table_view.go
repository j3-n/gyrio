package components

import (
	"errors"
	"image"
	"unicode/utf8"

	ui "github.com/gizak/termui/v3"
)

var ErrColumnCountMismatch = errors.New("number of columns in each row must be constant")

// Positions describe whether a border is on top, middle or bottom of its box
type Position int

var (
	PositionTop    Position = 0
	PositionMiddle Position = 1
	PositionBottom Position = 2
)

// TableView is a termui component for drawing tables.
type TableView struct {
	ui.Block

	// Columns is a slice of column titles
	Columns []string
	// Data is the data to populate the table with
	Data [][]string
	// Style to use for column titles
	ColumnTitleStyle ui.Style
	// Style to use for column borders
	TableBorderStyle ui.Style
}

// NewTableView initialises and returns a new empty TableView widget.
func NewTableView() *TableView {
	return &TableView{
		Block:            *ui.NewBlock(),
		Columns:          []string{},
		Data:             [][]string{},
		ColumnTitleStyle: ui.NewStyle(ui.ColorWhite),
		TableBorderStyle: ui.NewStyle(ui.ColorWhite),
	}
}

func (t *TableView) Draw(buf *ui.Buffer) {
	t.Block.Draw(buf)

	widths, err := computeWidths(t.Columns, t.Data)
	if err != nil {
		// TODO: report error to user
		return
	}

	// Draw column titles
	drawRow(buf, t.TableBorderStyle, t.Inner.Min, t.Columns, widths, PositionTop)

	// Draw rows
	for i, r := range t.Data {
		var p Position
		switch i {
		case len(t.Data) - 1:
			p = PositionBottom
		default:
			p = PositionMiddle
		}
		drawRow(buf, t.TableBorderStyle, t.Inner.Min.Add(image.Pt(0, (2*i)+2)), r, widths, p)
	}
}

// drawHorizontalBorder draws a horizontal table border at the given offset with the given column widths.
// position should be one of PositionTop, PositionMiddle, or PositionBottom depending on where the border is in the table.
func drawHorizontalBorder(buf *ui.Buffer, style ui.Style, offset image.Point, widths []int, position Position) {
	// Choose runes for first, middle and last corners
	var f, m, l rune
	switch position {
	case PositionBottom:
		f = ui.BOTTOM_LEFT
		m = ui.HORIZONTAL_UP
		l = ui.BOTTOM_RIGHT
	case PositionMiddle:
		f = ui.VERTICAL_RIGHT
		m = '┼'
		l = ui.VERTICAL_LEFT
	default:
		f = ui.TOP_LEFT
		m = ui.HORIZONTAL_DOWN
		l = ui.TOP_RIGHT
	}
	buf.SetCell(ui.NewCell(f, style), offset)

	// Draw horizontal section
	i := 1
	for j, w := range widths {
		for range w {
			buf.SetCell(ui.NewCell(ui.HORIZONTAL_LINE, style), offset.Add(image.Pt(i, 0)))
			i += 1
		}
		// Draw intermediate corner
		if j < len(widths)-1 {
			buf.SetCell(ui.NewCell(m, style), offset.Add(image.Pt(i, 0)))
			i += 1
		}
	}

	// Draw last corner
	buf.SetCell(ui.NewCell(l, style), offset.Add(image.Pt(i, 0)))
}

// drawRow draws the given row to the buffer at the given offset with the specified column widths.
// Each row draws the border above and beside itself, except one with PositionBottom which will draw its bottom border as well.
// Position indicates whether the row is the first, middle or last row of the table.
func drawRow(buf *ui.Buffer, style ui.Style, offset image.Point, row []string, widths []int, position Position) {
	// Draw border above
	switch position {
	case PositionTop:
		drawHorizontalBorder(buf, style, offset, widths, position)
	default:
		drawHorizontalBorder(buf, style, offset, widths, PositionMiddle)
	}

	// Draw row contents
	buf.SetCell(ui.NewCell(ui.VERTICAL_LINE, style), offset.Add(image.Pt(0, 1)))
	i := 1
	for j, w := range widths {
		for k := range w {
			if k < len(row[j]) {
				// Draw text
				r, _ := utf8.DecodeRune([]byte{row[j][k]})
				buf.SetCell(ui.NewCell(r, style), offset.Add(image.Pt(i, 1)))
			}
			i += 1
		}
		// Draw edge of columns
		buf.SetCell(ui.NewCell(ui.VERTICAL_LINE, style), offset.Add(image.Pt(i, 1)))
		i += 1
	}

	// Draw bottom border if needed
	if position == PositionBottom {
		drawHorizontalBorder(buf, style, offset.Add(image.Pt(0, 2)), widths, position)
	}
}

// computeWidths computes a slice of the maximum widths of each column for a set of data.
// Returns an ErrColumnCountMismatch if a row contains more columns than titles.
func computeWidths(titles []string, data [][]string) ([]int, error) {
	result := []int{}
	// Initialise columns
	for range data[0] {
		result = append(result, 0)
	}
	for _, r := range data {
		if len(r) != len(titles) {
			return []int{}, ErrColumnCountMismatch
		}
		for i, c := range r {
			result[i] = max(len(c), result[i], len(titles[i]))
		}
	}
	return result, nil
}
