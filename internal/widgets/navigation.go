package widgets

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// SetupNavigation installs global input handlers for navigating between widgets.
// - `app` is the tview.Application
// - `widgets` are the focusable boxes
// - `openScreen` is the callback to open the focused screen on Enter
func HandleNavigation(app *tview.Application, widgets []tview.Primitive, openScreen func(index int)) {
	// Focused on Calendar by default (upon entering application)
	focusIndex := 0

	// Change border colors and styles when focus changes
	updateFocus := func(index int) {
		for i, w := range widgets {
			if box, ok := w.(*tview.Box); ok {
				if i == index {
					box.SetBorderColor(tcell.ColorSpringGreen).
						SetTitleColor(tcell.ColorSpringGreen).
						SetBorderAttributes(tcell.AttrBold)
				} else {
					box.SetBorderColor(tcell.ColorWhite).
						SetTitleColor(tcell.ColorWhite).
						SetBorderAttributes(tcell.AttrNone)
				}
			}
		}
	}

	// Set up input handling for navigation and screen opening
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

	// Set initial focus highlight
	// updateFocus(focusIndex)
}
