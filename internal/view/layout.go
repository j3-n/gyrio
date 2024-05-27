package view

import (
	ui "github.com/gizak/termui/v3"
	"github.com/j3-n/gyrio/internal/components"
	"github.com/j3-n/gyrio/internal/pkg/util"
)

const (
	INTERACT_KEY = "<Enter>"
)

// A layout is a Grid that can have widgets selected so that they can receive input
type Layout struct {
	*ui.Grid

	// NavLayout represents the navigation structure of the Layout in Rows and Columns.
	NavLayout [][]components.Interactable

	interactMode bool
	row, col     int
}

func (l *Layout) Draw(buf *ui.Buffer) {
	l.Grid.Draw(buf)
}

func NewLayout(grid *ui.Grid, layout [][]components.Interactable) *Layout {
	l := &Layout{
		Grid:      grid,
		NavLayout: layout,
	}

	// Initialise layout styling
	for i, r := range layout {
		for j, c := range r {
			c.SetActive(false)
			if i == 0 && j == 0 {
				c.SetHovered(true)
			} else {
				c.SetHovered(false)
			}
		}
	}

	return l
}

// Navigate right (+) or left (-) x spaces
func (l *Layout) NavX(x int) {
	if len(l.NavLayout) > 0 {
		if len(l.NavLayout[l.row]) > 0 {
			l.NavLayout[l.row][l.col].SetHovered(false)
			l.col = util.Mod(l.col+x, len(l.NavLayout[l.row]))
			l.NavLayout[l.row][l.col].SetHovered(true)
		}
	}
}

// Navigate down (+) or up (-) y spaces
func (l *Layout) NavY(y int) {
	if len(l.NavLayout) > 0 {
		if len(l.NavLayout[l.row]) > l.col {
			l.NavLayout[l.row][l.col].SetHovered(false)
		}
		l.row = util.Mod(l.row+y, len(l.NavLayout))
		l.col = min(len(l.NavLayout[l.row])-1, l.col)
		l.NavLayout[l.row][l.col].SetHovered(true)
	}
}

// Toggle interact mode
func (l *Layout) ToggleInteract() {
	l.interactMode = !l.interactMode
	l.NavLayout[l.row][l.col].SetActive(l.interactMode)
}

func (l *Layout) KeyboardEvent(e *ui.Event) {
	if e.ID == INTERACT_KEY {
		l.ToggleInteract()
	} else {
		if l.interactMode {
			// Send input to the selected element
			if len(l.NavLayout) > 0 && len(l.NavLayout) >= l.row-1 {
				c := l.NavLayout[l.row]
				if len(c) >= l.col-1 {
					c[l.col].KeyboardEvent(e)
				}
			}
		} else {
			// Navigation input
			switch e.ID {
			case "<Left>":
				l.NavX(-1)
			case "<Right>":
				l.NavX(1)
			case "<Up>":
				l.NavY(-1)
			case "<Down>":
				l.NavY(1)
			}
		}
	}
}
