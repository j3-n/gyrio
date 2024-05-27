package util

import (
	ui "github.com/gizak/termui/v3"
)

var (
	STYLE_SELECTED   = ui.NewStyle(ui.ColorWhite, ui.ColorClear, ui.ModifierBold)
	STYLE_UNSELECTED = ui.NewStyle(8) // grey
	STYLE_HOVER      = ui.NewStyle(ui.ColorWhite)
)
