package components

import (
	"fmt"
	"image"

	ui "github.com/gizak/termui/v3"
)

// Toolbar is a component taking up one line at the bottom of its container to show interaction options.
type Toolbar struct {
	ui.Block
	// Elements is a list of elements to draw in the toolbar, represented by a slice of toolbar entries.
	Elements []ToolbarEntry
}

// NewToolbar initialises a new toolbar component and returns it.
func NewToolbar() *Toolbar {
	toolbar := &Toolbar{
		Block:    *ui.NewBlock(),
		Elements: []ToolbarEntry{},
	}
	toolbar.Border = false
	return toolbar
}

func (t *Toolbar) Draw(buf *ui.Buffer) {
	t.Block.Draw(buf)

	// Draw each element
	i := 1
	for _, e := range t.Elements {
		// Draw key prompt
		keyPrompt := fmt.Sprintf("[%s]", e.Key)
		promptStyle := ui.NewStyle(ui.ColorBlack, ui.ColorWhite)
		for _, c := range keyPrompt {
			buf.SetCell(ui.NewCell(c, promptStyle), image.Pt(i, t.Inner.Max.Y-1))
			i += 1
		}
		i += 1
		// Draw the accompanying text
		textStyle := ui.NewStyle(ui.ColorWhite)
		for _, c := range e.Text {
			buf.SetCell(ui.NewCell(c, textStyle), image.Pt(i, t.Inner.Max.Y-1))
			i += 1
		}
	}
}
