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

// Styles for input component
var (
	STYLE_CURSOR = ui.NewStyle(ui.ColorBlack, (ui.Color)(248))
)

// Styles for toolbar
var (
	STYLE_TOOLBAR_KEY  = ui.NewStyle(ui.ColorBlack, (ui.Color)(248))
	STYLE_TOOLBAR_TEXT = ui.NewStyle(248)
)

// Styles for error popups
var (
	STYLE_ERROR_BORDER         = ui.NewStyle(ui.ColorRed)
	STYLE_ERROR_TITLE          = ui.NewStyle(ui.ColorClear, ui.ColorRed)
	STYLE_ERROR_TEXT           = ui.NewStyle(ui.ColorRed)
	STYLE_ERROR_DISMISS_BUTTON = ui.NewStyle(ui.ColorRed, (ui.Color)(248))
)
