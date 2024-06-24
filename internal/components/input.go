package components

import (
	"image"
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

	observers []func()
}

// NewInput initialises and returns a new Input component.
func NewInput() *Input {
	return &Input{
		Block: *ui.NewBlock(),
	}
}

func (i *Input) KeyboardEvent(e *ui.Event) {
	if e.Type == ui.KeyboardEvent {
		if !strings.Contains(e.ID, "<") && len(e.ID) == 1 {
			i.Insert(e.ID)
		} else {
			switch e.ID {
			case "<Space>":
				i.Insert(" ")
			case "<Left>":
				i.NavLeft()
			case "<Right>":
				i.NavRight()
			case "<Backspace>":
				i.Backspace()
			case "<Enter>":
				i.NotifyObservers()
			}
		}
	}
}

func (i *Input) Insert(c string) {
	i.CurrentText = i.CurrentText[:i.cursorPos] + c + i.CurrentText[i.cursorPos:]
	i.cursorPos += 1
}

func (i *Input) NavLeft() {
	i.cursorPos = max(0, i.cursorPos-1)
}

func (i *Input) NavRight() {
	i.cursorPos = min(len(i.CurrentText), i.cursorPos+1)
}

func (i *Input) Backspace() {
	if i.cursorPos > 0 {
		i.CurrentText = i.CurrentText[:i.cursorPos-1] + i.CurrentText[i.cursorPos:]
		i.cursorPos -= 1
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

	// Render text
	body := strings.Split(util.WrapString(i.CurrentText, i.Block.Inner.Dx()-1), "\n")
	for j, str := range body[max(0, len(body)-i.Block.Inner.Dy()):] {
		buf.SetString(str, ui.NewStyle(ui.ColorWhite), i.Block.Inner.Min.Add(image.Pt(1, j)))
	}
	// Render cursor
	if i.isActive {
		row := min(i.cursorPos/(i.Block.Inner.Dx()-1), i.Block.Inner.Dy()-1)
		col := i.cursorPos % (i.Block.Inner.Dx() - 1)
		p := i.Block.Inner.Min.Add(image.Pt(col+1, row))
		r := buf.GetCell(p).Rune
		buf.SetCell(ui.NewCell(r, util.STYLE_CURSOR), p)
	}
}

func (i *Input) RegisterObserver(f func()) {
	i.observers = append(i.observers, f)
}

func (i *Input) NotifyObservers() {
	for _, f := range i.observers {
		f()
	}
}
