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
			textView, ok := w.(*tview.TextView)
			if !ok {
				continue
			}

			if i == index {
				textView.SetBorderColor(tcell.ColorSpringGreen)
				textView.SetTitleColor(tcell.ColorSpringGreen)
				textView.SetBorderAttributes(tcell.AttrBold)
			} else {
				textView.SetBorderColor(tcell.ColorWhite)
				textView.SetTitleColor(tcell.ColorWhite)
				textView.SetBorderAttributes(tcell.AttrNone)
			}
		}
	}

	// Set up input handling for navigation and screen opening
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
			case tcell.KeyRune:
				switch event.Rune() {
					case 'q', 'Q': // q to quit app globally
						app.Stop()
						return nil

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

	updateFocus(focusIndex)
}
