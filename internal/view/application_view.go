package view

import (
	"image"
	"sync"

	ui "github.com/gizak/termui/v3"
	"github.com/j3-n/gyrio/internal/components"
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
	// Stores the stack of errors to be displayed over the application view.
	errors []string

	sync.Mutex
}

// NewApplicationView creates a new view with the given set of available screen layouts.
func NewApplicationView(screens []*Layout) *ApplicationView {
	// Initialise toolbar component
	return &ApplicationView{
		screens:       screens,
		currentScreen: 0,
		errors:        []string{},
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

	v.drawErrors(buf)
}

func (v *ApplicationView) drawErrors(buf *ui.Buffer) {
	if len(v.errors) > 0 {
		// Draw most recent error
		err := v.errors[len(v.errors)-1]
		popup := components.NewErrorPopup(err, len(v.errors))
		center := v.rect.Max.Div(2)
		w := float32(v.rect.Max.X) / 1.5
		h := float32(v.rect.Max.Y) / 1.5
		popup.SetRect(center.X-int(w/2), center.Y-int(h/2), center.X+int(w/2), center.Y+int(h/2))
		buf.Fill(ui.NewCell(' ', util.STYLE_ERROR_BORDER), popup.GetRect().Inset(-1))
		popup.Title = "Error"
		popup.TitleStyle = util.STYLE_ERROR_TITLE
		popup.BorderStyle = util.STYLE_ERROR_BORDER

		popup.Draw(buf)
	}
}

// KeyboardEvent handles keyboard inputs to this view. If it is not a control input
// it will be forwarded to the active widget.
func (v *ApplicationView) KeyboardEvent(e *ui.Event) {
	if e.ID == "<Tab>" {
		// Cycle along one screen layout
		v.currentScreen = (v.currentScreen + 1) % len(v.screens)
	} else if len(v.errors) > 0 && e.ID == "<Enter>" {
		v.dismissError()
	} else {
		// Pass down to current layout
		v.screens[v.currentScreen].KeyboardEvent(e)
	}
}

// ShowError displays an error on top of the application view. The most recent error renders on top.
func (v *ApplicationView) ShowError(err string) {
	v.errors = append(v.errors, err)
}

func (v *ApplicationView) dismissError() {
	v.errors = v.errors[:len(v.errors)-1]
}
