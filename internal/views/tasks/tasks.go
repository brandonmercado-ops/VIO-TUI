package tasks

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func TasksPage(app *tview.Application, returnTo func()) tview.Primitive {

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
▗▄▄▄▖▗▄▖  ▗▄▄▖▗▖ ▗▖ ▗▄▄▖
  █ ▐▌ ▐▌▐▌   ▐▌▗▞▘▐▌   
  █ ▐▛▀▜▌ ▝▀▚▖▐▛▚▖  ▝▀▚▖
  █ ▐▌ ▐▌▗▄▄▞▘▐▌ ▐▌▗▄▄▞▘
		`)

	header := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(quit, 0, 4, false). // narrow weight
		AddItem(title, 0, 6, false) // wider weight

	// Paddings inbetween both boxes and left and right of screen

	leftPadding := tview.NewBox()
	middlePadding := tview.NewBox()
	rightPadding := tview.NewBox()

	// TODO, IN PROGRESS, and DONE task sections

	todo := tview.NewBox().SetBorder(true).SetTitle("[ TODO ]")
	inProg := tview.NewBox().SetBorder(true).SetTitle("[ IN PROGRESS ]")
	done := tview.NewBox().SetBorder(true).SetTitle("[ DONE ]")

	// Main body with paddings

	mainBody := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(leftPadding, 0, 1, false).
		AddItem(todo, 0, 8, false).
		AddItem(middlePadding, 0, 1, false).
		AddItem(inProg, 0, 8, false).
		AddItem(middlePadding, 0, 1, false).
		AddItem(done, 0, 7, false).
		AddItem(rightPadding, 0, 1, false)

	// Footer below all 3 task columns
	footer := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true) // .
		// SetText("FOOTER FOOTER FOOTER")

	// Bringing header, main body, and footer together with paddings
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
