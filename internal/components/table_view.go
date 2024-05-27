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

const (
	PositionTop Position = iota
	PositionMiddle
	PositionBottom
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

	scrollX, scrollY    int
	isHovered, isActive bool
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
	maxRows, maxCols := computeBounds(t.Inner.Min, t.Inner.Max.Sub(image.Pt(1, 2)), 3, widths, t.scrollX)
	maxRows = min(maxRows, len(t.Data))
	maxCols = min(maxCols, len(t.Columns))
	// Cap Y scroll
	t.scrollY = min(len(t.Data)-maxRows, t.scrollY)
	// Crop visible columns and rows to fit container
	croppedWidths := widths[t.scrollX:min(maxCols+t.scrollX, len(widths))]
	croppedColumns := t.Columns[t.scrollX:min(maxCols+t.scrollX, len(t.Columns))]
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
			drawRow(buf, t.TableBorderStyle, t.TextAlignment, t.Inner.Min.Add(image.Pt(0, (2*i)+2)), r[t.scrollX:min(maxCols+t.scrollX, len(r))], croppedWidths, p)
		}
		// Draw another border if this isn't the end of the table (it's been cropped)
		if t.scrollY+len(croppedRows) < len(t.Data) {
			drawHorizontalBorder(buf, t.TableBorderStyle, t.Inner.Min.Add(image.Pt(0, 2*maxRows+2)), croppedWidths, PositionMiddle)
		}
		// Draw the same for vertical
		if t.scrollX+len(croppedColumns) < len(t.Columns) {
			width := util.Sum(croppedWidths) + len(croppedWidths)
			buf.SetCell(ui.NewCell(ui.HORIZONTAL_DOWN, t.TableBorderStyle), t.Inner.Min.Add(image.Pt(width, 0)))
			for i := range len(croppedRows) {
				buf.SetCell(ui.NewCell(util.CROSS, t.TableBorderStyle), t.Inner.Min.Add(image.Pt(width, i*2+2)))
			}
			var r rune
			if t.scrollY+len(croppedRows) < len(t.Data) {
				r = util.CROSS
			} else {
				r = ui.HORIZONTAL_UP
			}
			buf.SetCell(ui.NewCell(r, t.TableBorderStyle), t.Inner.Min.Add(image.Pt(width, len(croppedRows)*2+2)))
		}
	} else {
		drawEmptyRow(buf, t.TableBorderStyle, croppedWidths, t.Inner.Min)
	}
	drawScrollArrows(buf, t.ColumnTitleStyle, image.Pt(t.scrollX, t.scrollY), image.Pt(maxCols, maxRows), image.Pt(len(t.Columns), len(t.Data)), t.Inner)
}

// drawEmptyRow draws an empty row with no borders spanning the width of all the given columns.
func drawEmptyRow(buf *ui.Buffer, style ui.Style, widths []int, origin image.Point) {
	// Table is empty, draw a small empty row
	width := util.Sum(widths) + len(widths) - 1
	drawHorizontalBorder(buf, style, origin.Add(image.Pt(0, 2)), widths, PositionMiddle)
	drawHorizontalBorder(buf, style, origin.Add(image.Pt(0, 3)), []int{width}, PositionBottom)
	// Merge columns
	i := 0
	for j, w := range widths {
		i += w + 1
		var r rune
		switch j {
		case len(widths) - 1:
			r = ui.VERTICAL_LEFT
		default:
			r = ui.HORIZONTAL_UP
		}
		buf.SetCell(ui.NewCell(r, style), origin.Add(image.Pt(i, 2)))
	}
}

// drawScrollArrows draws the required scrolling arrows onto the buffer.
// scroll is the current X and Y scroll, max is the max number of rows (Y) and cols (X) as calculated by compute Bounds,
// data is the size of the rows and columns of the data itself. bounds are the inner bounds of the element.
func drawScrollArrows(buf *ui.Buffer, style ui.Style, scroll image.Point, max image.Point, data image.Point, bounds image.Rectangle) {
	if scroll.X > 0 {
		// Left scroll arrow
		buf.SetCell(ui.NewCell(util.LEFT_ARROW, style), image.Pt(1, bounds.Dy()/2))
	}
	if scroll.Y > 0 {
		// Top scroll arrow
		buf.SetCell(ui.NewCell(ui.UP_ARROW, style), image.Pt(bounds.Dx()/2, 1))
	}
	if scroll.X+max.X < data.X {
		// Right scroll arrow
		buf.SetCell(ui.NewCell(util.RIGHT_ARROW, style), bounds.Max.Sub(image.Pt(1, bounds.Dy()/2+1)))
	}
	if scroll.Y+max.Y < data.Y {
		// Bottom scroll arrow
		buf.SetCell(ui.NewCell(ui.DOWN_ARROW, style), bounds.Max.Sub(image.Pt(bounds.Dx()/2+1, 1)))
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
		m = util.CROSS
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
	t.scrollY = max(0, min(t.scrollY+1, len(t.Data)-1))
}

// ScrollLeft scrolls the view left by 1 row.
func (t *TableView) ScrollLeft() {
	t.scrollX = max(0, t.scrollX-1)
}

// ScrollRight scrolls the view right by 1 column.
func (t *TableView) ScrollRight() {
	t.scrollX = max(0, min(t.scrollX+1, len(t.Columns)-1))
}

func (t *TableView) KeyboardEvent(e *ui.Event) {
	switch e.ID {
	case "<Left>":
		t.ScrollLeft()
	case "<Right>":
		t.ScrollRight()
	case "<Up>":
		t.ScrollUp()
	case "<Down>":
		t.ScrollDown()
	}
}

// TODO: refactor these out somehow
func (t *TableView) SetActive(active bool) {
	t.isActive = active
	t.UpdateBorderStyle()
}

func (t *TableView) SetHovered(hovered bool) {
	t.isHovered = hovered
	t.UpdateBorderStyle()
}

func (t *TableView) UpdateBorderStyle() {
	if t.isActive {
		t.Block.BorderStyle = util.STYLE_SELECTED
	} else if t.isHovered {
		t.Block.BorderStyle = util.STYLE_HOVER
	} else {
		t.Block.BorderStyle = util.STYLE_UNSELECTED
	}
}
