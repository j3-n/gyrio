package view

import (
	"image"
	"sync"

	ui "github.com/gizak/termui/v3"
	"github.com/j3-n/gyrio/internal/pkg/util"
)

const (
	TOOLBAR_HEIGHT = 1
)

// ApplicationView allows one grid layout to be drawn on screen at a time.
type ApplicationView struct {
	rect image.Rectangle
	// screens is the slice of screens this view has stored
	screens []*Layout
	// currentScreen stores the index of the active screen which will be rendered
	currentScreen int

	sync.Mutex
}

// NewApplicationView creates a new view with the given set of available screen layouts.
func NewApplicationView(screens []*Layout) *ApplicationView {
	// Initialise toolbar component
	return &ApplicationView{
		screens:       screens,
		currentScreen: 0,
	}
}

func (v *ApplicationView) GetRect() image.Rectangle {
	return v.rect
}

func (v *ApplicationView) SetRect(x1 int, y1 int, x2 int, y2 int) {
	v.rect = image.Rect(x1, y1, x2, y2)
	// Pass down to screens, save
	for _, s := range v.screens {
		s.SetRect(x1, y1, x2, y2-TOOLBAR_HEIGHT)
	}
}

func (v *ApplicationView) Draw(buf *ui.Buffer) {
	v.screens[v.currentScreen].Draw(buf)
	// Draw toolbar
	e := v.screens[v.currentScreen].GetToolbarEntry()
	buf.SetString(e.Key, util.STYLE_TOOLBAR_KEY, image.Pt(v.rect.Min.X, v.rect.Max.Y-TOOLBAR_HEIGHT))
	buf.SetString(e.Text, util.STYLE_TOOLBAR_TEXT, image.Pt(v.rect.Min.X+len(e.Key)+1, v.rect.Max.Y-TOOLBAR_HEIGHT))
}

// KeyboardEvent handles keyboard inputs to this view. If it is not a control input
// it will be forwarded to the active widget.
func (v *ApplicationView) KeyboardEvent(e *ui.Event) {
	if e.ID == "<Tab>" {
		// Cycle along one screen layout
		v.currentScreen = (v.currentScreen + 1) % len(v.screens)
	} else {
		// Pass down to current layout
		v.screens[v.currentScreen].KeyboardEvent(e)
	}
}
