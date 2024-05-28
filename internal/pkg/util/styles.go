package util

import (
	ui "github.com/gizak/termui/v3"
)

// Styles for widgets
var (
	STYLE_SELECTED   = ui.NewStyle(46)
	STYLE_UNSELECTED = ui.NewStyle(8)
	STYLE_HOVER      = ui.NewStyle(ui.ColorWhite)
)

var (
	STYLE_CURSOR = ui.NewStyle(ui.ColorBlack, (ui.Color)(248))
)
