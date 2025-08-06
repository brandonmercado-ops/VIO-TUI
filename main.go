package main

import (
	"VIO/ui"

	"github.com/rivo/tview"
)

func main() {

	// Create new TUI Application
	app := tview.NewApplication()

	// Grab widgets
	widgets, flex := ui.BuildMainWidgets()

	// Choose screen to display (subject to change with widget-click-selection
	openScreen := ui.ScreenRouter(app, widgets, flex)

	// Handle navigation around widgets on main page
	ui.HandleNavigation(app, widgets, openScreen)

	// Start program
	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}
}
