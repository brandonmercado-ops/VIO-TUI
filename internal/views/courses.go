package views

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func CoursesPage(app *tview.Application, returnTo func()) tview.Primitive {
	view := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true).
		SetText("📚 [::b]Courses Page[::-]\n\n[gray]Press Esc to return")

	view.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEsc {
			returnTo()
			return nil
		}
		return event
	})

	return view
}
