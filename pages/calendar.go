package pages

import (
	// "fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func CalendarPage(app *tview.Application, returnTo func()) tview.Primitive {
	header := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true).
		SetText("📅 [::b]Calendar View[::-]\n\n[gray]Press Esc to return")

	calendarBox := tview.NewBox().SetBorder(true).SetTitle("[ CALENDAR ]")

	dailyScheduleMini := tview.NewBox().SetBorder(true).SetTitle("[ TODAY'S SCHEDULE ]")

	leftPadding := tview.NewBox()
	middlePadding := tview.NewBox()
	rightPadding := tview.NewBox()

	mainBody := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(leftPadding, 0, 1, false).
		AddItem(calendarBox, 0, 5, false).
		AddItem(middlePadding, 0, 1, false).
		AddItem(dailyScheduleMini, 0, 2, false).
		AddItem(rightPadding, 0, 1, false)

	footer := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true).
		SetText("FOOTER FOOTER FOOTER")

	page := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(header, 0, 2, false).
		AddItem(mainBody, 0, 5, false).
		AddItem(footer, 0, 1, false)

	page.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEsc {
			returnTo()
			return nil
		}
		return event
	})

	return page
}
