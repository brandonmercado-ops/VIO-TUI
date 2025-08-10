package schedule

import (
	"VIO/internal/asciiart"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"strconv"
	"time"
)

func SchedulePage(app *tview.Application, returnTo func()) tview.Primitive {

	// Header

	quitPadding := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true).
		SetText(`


		`)

	quitText := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignRight).
		SetText("[white::b][ ESC ] To RETURN TO MAIN")

	quit := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(quitPadding, 2, 1, false).
		AddItem(quitText, 0, 2, false)

	title := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)

	currMonth := time.Now().Month().String()
	currDayInt := time.Now().Day()      // returns integer
	currDay := strconv.Itoa(currDayInt) // converting integer to string for header map `internal/asciiart/days.go`

	monthAscii := asciiart.GetMonthHeader(currMonth)
	dayAscii := asciiart.GetDayHeader(currDay)
	dateAscii := monthAscii + dayAscii

	title.SetText(dateAscii)

	header := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(quit, 0, 2, false). // narrow weight
		AddItem(title, 0, 6, false) // wider weight
	// Main body

	// Paddings inbetween both boxes and left and right of screen

	leftPadding := tview.NewBox()
	rightPadding := tview.NewBox()

	// Schedule Box

	schedule := tview.NewBox()

	mainBody := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(leftPadding, 0, 1, false).
		AddItem(schedule, 0, 5, false).
		AddItem(rightPadding, 0, 1, false)

	// Footer

	footer := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true).
		SetText("[ + ] ADD   [ - ] REMOVE   [ E ] EDIT")

	page := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(header, 0, 3, false).
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
