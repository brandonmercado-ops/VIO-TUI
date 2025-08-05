package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// SetupGlobalKeys installs global key handling and returns the focus updater.
func SetupGlobalKeys(app *tview.Application, widgets []tview.Primitive, openScreen func(int)) func(int) {
	focusIndex := 0

	// updateFocus visually highlights a focused widget
	updateFocus := func(index int) {
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

	// Handle number key and enter key
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

	return updateFocus
}
