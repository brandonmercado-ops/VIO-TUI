package main

import (
	"VIO/internal/store"
	"VIO/internal/widgets"

	"github.com/rivo/tview"
)

func main() {
	// Load local JSON data
	data, err := store.LoadAppData()
	if err != nil {
		panic(err)
	}

	// Create new TUI Application
	app := tview.NewApplication()

	// Grab dashboard widgets
	mainWidgets, flex := widgets.BuildMainWidgets(data)

	// Choose screen to display (subject to change with widget-click-selection
	openScreen := widgets.ScreenRouter(app, mainWidgets, flex, data)

	// Handle navigation around widgets on main page
	widgets.HandleNavigation(app, mainWidgets, openScreen)

	// Start program
	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}
}
