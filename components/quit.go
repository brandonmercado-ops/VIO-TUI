package components

import (
	// "github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// NewQuitModal returns a centered modal popup with Quit/Cancel buttons.
func NewQuitModal(app *tview.Application, returnTo func()) tview.Primitive {
	modal := tview.NewModal().
		SetText("Are you sure you want to quit?").
		AddButtons([]string{"Cancel", "Quit"}).
		SetDoneFunc(func(buttonIndex int, label string) {
			if label == "Quit" {
				app.Stop()
			} else {
				returnTo()
			}
		})
	return modal
}
