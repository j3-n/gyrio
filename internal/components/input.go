package components

import (
	"fmt"
	"strings"

	ui "github.com/gizak/termui/v3"
	"github.com/j3-n/gyrio/internal/pkg/util"
)

type Interactable interface {
	SetActive(bool)
	SetHovered(bool)
	KeyboardEvent(*ui.Event)
}

type Input struct {
	ui.Block

	CurrentText string
	CharLimit   uint

	cursorPos int
	isActive  bool
	isHovered bool
}

// NewInput initialises and returns a new Input component.
func NewInput() *Input {
	return &Input{
		Block: *ui.NewBlock(),
	}
}

func (i *Input) KeyboardEvent(e *ui.Event) {
	if e.Type == ui.KeyboardEvent && !strings.Contains(e.ID, "<") {
		fmt.Println(e.ID)
	}
}

func (i *Input) SetActive(active bool) {
	i.isActive = active
	i.UpdateBorderStyle()
}

func (i *Input) SetHovered(hovered bool) {
	i.isHovered = hovered
	i.UpdateBorderStyle()
}

func (i *Input) UpdateBorderStyle() {
	if i.isActive {
		i.Block.BorderStyle = util.STYLE_SELECTED
	} else if i.isHovered {
		i.Block.BorderStyle = util.STYLE_HOVER
	} else {
		i.Block.BorderStyle = util.STYLE_UNSELECTED
	}
}

func (i *Input) Draw(buf *ui.Buffer) {
	i.Block.Draw(buf)
}
