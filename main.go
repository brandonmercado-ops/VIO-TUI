package main

import (
	"VIO/internal/widgets"

	"github.com/rivo/tview"
)

func main() {

	// Create new TUI Application
	app := tview.NewApplication()

	// Grab widgets
	Widgets, flex := widgets.BuildMainWidgets()

	// Choose screen to display (subject to change with widget-click-selection
	openScreen := widgets.ScreenRouter(app, Widgets, flex)

	// Handle navigation around widgets on main page
	widgets.HandleNavigation(app, Widgets, openScreen)

	// Start program
	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}
}
