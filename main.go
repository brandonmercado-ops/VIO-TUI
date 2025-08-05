package main

import (
	"VIO/pages"
	"VIO/ui"

	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {

	// Create new TUI Application
	app := tview.NewApplication()

	// Grab widgets
	widgets, flex := ui.BuildMainWidgets()

	openScreen := func(index int) {
		var screen tview.Primitive

		switch index {
		case 0:
			screen = pages.CalendarPage(app, func() {
				app.SetRoot(flex, true).SetFocus(widgets[index])
			})
		case 1:
			screen = pages.CoursesPage(app, func() {
				app.SetRoot(flex, true).SetFocus(widgets[index])
			})
		case 2:
			screen = pages.TasksPage(app, func() {
				app.SetRoot(flex, true).SetFocus(widgets[index])
			})
		case 3:
			screen = pages.SchedulePage(app, func() {
				app.SetRoot(flex, true).SetFocus(widgets[index])
			})
		case 4:
			screen = pages.AssignmentsPage(app, func() {
				app.SetRoot(flex, true).SetFocus(widgets[index])
			})

		default:
			tv := tview.NewTextView().
				SetTextAlign(tview.AlignCenter).
				SetDynamicColors(true).
				SetText(fmt.Sprintf("🌟 You selected [::b]Box %d[::-]!\n\n[gray]Press Esc to return", index+1))

			tv.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
				if event.Key() == tcell.KeyEsc {
					app.SetRoot(flex, true).SetFocus(widgets[index])
					return nil
				}
				return event
			})

			screen = tv
		}

		app.SetRoot(screen, true).SetFocus(screen)
	}

	// -----------------------------------------------------------------------------------------------
	//				HANDLING WIDGET HIGHLIGHT AND NAVIGATION
	// -----------------------------------------------------------------------------------------------

	// Focused on Calendar by default (upon entering application)
	focusIndex := 0

	// Anonymous function to change both focus and color of chosen widget
	updateFocus := func(index int) {

		// Change colors for highlighted widget and reset color when out of focus
		for i, w := range widgets {
			if box, ok := w.(*tview.Box); ok {
				if i == index {
					box.SetBorderColor(tcell.ColorSpringGreen)
					box.SetTitleColor(tcell.ColorSpringGreen)
					box.SetBorderAttributes(tcell.AttrBold)
				} else {
					box.SetBorderColor(tcell.ColorWhite)
					box.SetTitleColor(tcell.ColorWhite)
					box.SetBorderAttributes(tcell.AttrNone)
				}
			}
		}
	}

	// Focusing on other widgets indicated by their corresponding numbers
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyRune:
			switch event.Rune() {
			case '1', '2', '3', '4', '5':
				focusIndex = int(event.Rune() - '1')
				app.SetFocus(widgets[focusIndex])
				updateFocus(focusIndex)
			}
		case tcell.KeyEnter:
			openScreen(focusIndex)
		}

		return event
	})
	// -----------------------------------------------------------------------------------------------
	//				START-OF-PROGRAM ERROR HANDLING
	// -----------------------------------------------------------------------------------------------

	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}
}
