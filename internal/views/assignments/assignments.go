package assignments

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func AssignmentsPage(app *tview.Application, returnTo func()) tview.Primitive {

	// Header

	quitPadding := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true).
		SetText(`


		`)

	quitText := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter).
		SetText("[white::b][ ESC ] To RETURN TO MAIN")

	quit := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(quitPadding, 2, 1, false).
		AddItem(quitText, 0, 2, false)

	title := tview.NewTextView().
		SetTextAlign(tview.AlignLeft).
		SetDynamicColors(true).
		SetText(`
 ▗▄▖  ▗▄▄▖ ▗▄▄▖▗▄▄▄▖ ▗▄▄▖▗▖  ▗▖▗▖  ▗▖▗▄▄▄▖▗▖  ▗▖▗▄▄▄▖▗▄▄▖
▐▌ ▐▌▐▌   ▐▌     █  ▐▌   ▐▛▚▖▐▌▐▛▚▞▜▌▐▌   ▐▛▚▖▐▌  █ ▐▌   
▐▛▀▜▌ ▝▀▚▖ ▝▀▚▖  █  ▐▌▝▜▌▐▌ ▝▜▌▐▌  ▐▌▐▛▀▀▘▐▌ ▝▜▌  █  ▝▀▚▖
▐▌ ▐▌▗▄▄▞▘▗▄▄▞▘▗▄█▄▖▝▚▄▞▘▐▌  ▐▌▐▌  ▐▌▐▙▄▄▖▐▌  ▐▌  █ ▗▄▄▞▘
`)

	header := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(quit, 0, 3, false).
		AddItem(title, 0, 6, false)

	// Paddings

	leftPadding := tview.NewBox()
	rightPadding := tview.NewBox()

	// Main body

	assignmentsBody := tview.NewBox()

	mainBody := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(leftPadding, 0, 1, false).
		AddItem(assignmentsBody, 0, 6, false).
		AddItem(rightPadding, 0, 1, false)

	// Footer

	footer := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true).
		SetText("[ + ] ADD   [ - ] REMOVE   [ E ] EDIT")

	page := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(header, 0, 1, false).
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
