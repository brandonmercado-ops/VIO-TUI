package widgets

import (
	"VIO/internal/views"
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// ScreenRouter returns a function that opens the correct screen based on index.
// Pass in the app, widgets array, and layout to allow switching back.
func ScreenRouter(app *tview.Application, widgets []tview.Primitive, layout *tview.Flex) func(index int) {
	return func(index int) {
		var screen tview.Primitive

		switch index {
		case 0:
			screen = views.CalendarPage(app, func() {
				app.SetRoot(layout, true).SetFocus(widgets[index])
			})
		case 1:
			screen = views.CoursesPage(app, func() {
				app.SetRoot(layout, true).SetFocus(widgets[index])
			})
		case 2:
			screen = views.TasksPage(app, func() {
				app.SetRoot(layout, true).SetFocus(widgets[index])
			})
		case 3:
			screen = views.SchedulePage(app, func() {
				app.SetRoot(layout, true).SetFocus(widgets[index])
			})
		case 4:
			screen = views.AssignmentsPage(app, func() {
				app.SetRoot(layout, true).SetFocus(widgets[index])
			})
		default:
			tv := tview.NewTextView().
				SetTextAlign(tview.AlignCenter).
				SetDynamicColors(true).
				SetText(fmt.Sprintf("🌟 You selected [::b]Box %d[::-]!\n\n[gray]Press Esc to return", index+1))

			tv.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
				if event.Key() == tcell.KeyEsc {
					app.SetRoot(layout, true).SetFocus(widgets[index])
					return nil
				}
				return event
			})
			screen = tv
		}

		app.SetRoot(screen, true).SetFocus(screen)
	}
}
