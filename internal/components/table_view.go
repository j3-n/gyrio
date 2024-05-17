package components

import (
	"errors"
	"image"

	ui "github.com/gizak/termui/v3"
	"github.com/j3-n/gyrio/internal/pkg/util"
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
	// Text alignment mode to use for cell contents
	TextAlignment ui.Alignment
	// Style to use for column titles
	ColumnTitleStyle ui.Style
	// Style to use for column borders
	TableBorderStyle ui.Style

	scrollX, scrollY int
}

// NewTableView initialises and returns a new empty TableView widget.
func NewTableView() *TableView {
	return &TableView{
		Block:            *ui.NewBlock(),
		Columns:          []string{},
		Data:             [][]string{},
		TextAlignment:    ui.AlignLeft,
		ColumnTitleStyle: ui.NewStyle(ui.ColorWhite),
		TableBorderStyle: ui.NewStyle(ui.ColorWhite),

		scrollX: 0,
		scrollY: 0,
	}
}

func (t *TableView) Draw(buf *ui.Buffer) {
	t.Block.Draw(buf)

	widths, err := computeWidths(t.Columns, t.Data)
	if err != nil {
		// TODO: report error to user
		return
	}

	// Check and enforce bounds
	maxRows, maxCols := computeBounds(t.Inner.Min, t.Inner.Max.Sub(image.Pt(1, 1)), 3, widths, t.scrollX)
	// Cap Y scroll
	t.scrollY = min(len(t.Data)-maxRows, t.scrollY)
	// Crop visible columns and rows to fit container
	croppedWidths := widths[:min(maxCols, len(widths))]
	croppedColumns := t.Columns[:min(maxCols, len(t.Columns))]
	croppedRows := t.Data[t.scrollY:min(maxRows+t.scrollY, len(t.Data))]

	// Draw column titles
	drawRow(buf, t.TableBorderStyle, t.TextAlignment, t.Inner.Min, croppedColumns, croppedWidths, PositionTop)

	// Draw rows
	if len(t.Data) > 0 {
		for i, r := range croppedRows {
			var p Position
			switch i + t.scrollY {
			case len(t.Data) - 1:
				p = PositionBottom
			default:
				p = PositionMiddle
			}
			drawRow(buf, t.TableBorderStyle, t.TextAlignment, t.Inner.Min.Add(image.Pt(0, (2*i)+2)), r[:min(maxCols, len(r))], croppedWidths, p)
		}
		// Draw another border if this isn't the end of the table (it's been cropped)
		if t.scrollY+len(croppedRows) < len(t.Data) {
			drawHorizontalBorder(buf, t.TableBorderStyle, t.Inner.Min.Add(image.Pt(0, 2*maxRows+2)), croppedWidths, PositionMiddle)
			buf.SetCell(ui.NewCell(ui.DOWN_ARROW, t.ColumnTitleStyle), t.Inner.Max.Sub(image.Pt(t.Inner.Max.X/2+1, 1)))
		}
	} else {
		// Table is empty, draw a small empty row
		width := util.Sum(croppedWidths) + len(croppedWidths) - 1
		drawHorizontalBorder(buf, t.TableBorderStyle, t.Inner.Min.Add(image.Pt(0, 2)), croppedWidths, PositionMiddle)
		drawHorizontalBorder(buf, t.TableBorderStyle, t.Inner.Min.Add(image.Pt(0, 3)), []int{width}, PositionBottom)
		// Merge columns
		i := 0
		for j, w := range croppedWidths {
			i += w + 1
			var r rune
			switch j {
			case len(croppedWidths) - 1:
				r = ui.VERTICAL_LEFT
			default:
				r = ui.HORIZONTAL_UP
			}
			buf.SetCell(ui.NewCell(r, t.TableBorderStyle), t.Inner.Min.Add(image.Pt(i, 2)))
		}
	}
	// Draw column extra arrow if needed
	if len(t.Columns) > maxCols {
		buf.SetCell(ui.NewCell('►', t.ColumnTitleStyle), t.Inner.Max.Sub(image.Pt(1, t.Inner.Max.Y/2+1)))
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
func drawRow(buf *ui.Buffer, style ui.Style, align ui.Alignment, offset image.Point, row []string, widths []int, position Position) {
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
		// Calculate text beginning offset based on text alignment
		var textOffset int
		switch align {
		case ui.AlignLeft:
			textOffset = 0
		case ui.AlignRight:
			textOffset = w - len(row[j])
		default:
			textOffset = (w - len(row[j])) / 2
		}
		// Draw row text
		for k, c := range row[j] {
			buf.SetCell(ui.NewCell(c, style), offset.Add(image.Pt(i+textOffset+k, 1)))
		}
		i += w + 1
		// Draw edge of columns
		buf.SetCell(ui.NewCell(ui.VERTICAL_LINE, style), offset.Add(image.Pt(i-1, 1)))
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
	if len(data) > 0 {
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
	} else {
		// Empty table
		for _, t := range titles {
			result = append(result, len(t))
		}
	}
	return result, nil
}

// computeBounds runs a bounds check on the table based on the current horizontal scroll and row height and returns
// the maximum number of rows and columns that can fit in the space the widget has been given at one time.
// Takes into account the borders between cells.
func computeBounds(min image.Point, max image.Point, rowHeight int, widths []int, currentCol int) (int, int) {
	bounds := max.Sub(min)
	// Compute max rows
	rows := bounds.Y / rowHeight
	// Compute max cols
	i := 1
	cols := 0
	for _, col := range widths[currentCol:] {
		if i+col <= bounds.X {
			cols += 1
			i += col + 1
		} else {
			break
		}
	}
	return rows, cols
}

// ScrollUp scrolls the view up by 1 row.
func (t *TableView) ScrollUp() {
	t.scrollY = max(t.scrollY-1, 0)
}

// ScrollUp scrolls the view up by 1 row.
func (t *TableView) ScrollDown() {
	t.scrollY = min(t.scrollY+1, len(t.Data)-1)
}
