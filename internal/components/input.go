package components

import (
	"fmt"
	"strings"

	ui "github.com/gizak/termui/v3"
)

type Input struct {
	ui.Block

	DefaultText string
	CurrentText string
	CharLimit   uint
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

func (i *Input) Draw(buf *ui.Buffer) {
	i.Block.Draw(buf)
}
