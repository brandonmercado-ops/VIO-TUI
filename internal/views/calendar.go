package views

import (
	// "fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func CalendarPage(app *tview.Application, returnTo func()) tview.Primitive {

	// Header above (later to denote month of the year)
	header := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true).
		SetText("📅 [::b]Calendar View[::-]\n\n[gray]Press Esc to return")

	// Box on left-middle side of screen that shows days of the month
	calendarBox := tview.NewBox().SetBorder(true).SetTitle("[ CALENDAR ]")

	// Mini daily schedule menu that lists all meetings for the day
	dailyScheduleMini := tview.NewBox().SetBorder(true).SetTitle("[ TODAY'S SCHEDULE ]")

	// Paddings inbetween both boxes and left and right of screen
	leftPadding := tview.NewBox()
	middlePadding := tview.NewBox()
	rightPadding := tview.NewBox()

	// Assembling the calendar box and mini daily schedule with all paddings
	mainBody := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(leftPadding, 0, 1, false).
		AddItem(calendarBox, 0, 5, false).
		AddItem(middlePadding, 0, 1, false).
		AddItem(dailyScheduleMini, 0, 2, false).
		AddItem(rightPadding, 0, 1, false)

	// Footer below calendar and mini daily schedule menu
	footer := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true).
		SetText("FOOTER FOOTER FOOTER")

	// Bringing header, main body, and footer together
	page := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(header, 0, 2, false).
		AddItem(mainBody, 0, 5, false).
		AddItem(footer, 0, 1, false)

	// Listening for escape key to quit back to main page
	page.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEsc {
			returnTo()
			return nil
		}
		return event
	})

	return page
}
