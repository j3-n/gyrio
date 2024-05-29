package components

import (
	"image"
	"strings"

	ui "github.com/gizak/termui/v3"
	"github.com/j3-n/gyrio/internal/pkg/util"
)

type ErrorPopup struct {
	ui.Block

	// The error message to display in the popup
	ErrorMsg string
	// The total number of errors currently being displayed
	NumErrors int
}

func NewErrorPopup(msg string, num int) *ErrorPopup {
	return &ErrorPopup{
		Block:     *ui.NewBlock(),
		ErrorMsg:  msg,
		NumErrors: num,
	}
}

func (e *ErrorPopup) Draw(buf *ui.Buffer) {
	e.Block.Draw(buf)

	// Draw error text and wrap it
	for i, line := range strings.Split(util.WrapString(e.ErrorMsg, e.Block.Inner.Dx()-1), "\n") {
		if i > e.Block.Inner.Dy()-2 {
			break
		}
		buf.SetString(line, util.STYLE_ERROR_TEXT, e.Block.Inner.Min.Add(image.Pt(0, i)))
	}

	buf.SetString("<OK>", util.STYLE_ERROR_DISMISS_BUTTON, e.Block.Inner.Max.Sub(image.Pt(e.Block.Inner.Dx()/2+2, 1)))
}
