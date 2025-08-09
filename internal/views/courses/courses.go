package courses

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func CoursesPage(app *tview.Application, returnTo func()) tview.Primitive {

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
		SetText("COURSES PAGE - TEMP HEADER")

	header := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(quit, 0, 4, false). // narrow weight
		AddItem(title, 0, 6, false) // wider weight

	// School and student info

	schoolTitle := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetText(`



 ▗▄▄▖ ▗▄▄▖▗▖ ▗▖▗▖  ▗▖
▐▌   ▐▌   ▐▌ ▐▌▐▛▚▖▐▌
▐▌    ▝▀▚▖▐▌ ▐▌▐▌ ▝▜▌
▝▚▄▄▖▗▄▄▞▘▝▚▄▞▘▐▌  ▐▌
                     


		`)

	studentInfo := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetText(`
무명왕 님
대학 4학년
컴퓨터 전공
		`)

	schoolInfo := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(schoolTitle, 0, 1, false).
		AddItem(studentInfo, 0, 1, false)

	// Course info

	courseInfo := tview.NewBox().SetBorder(true).SetTitle("[ COURSES ]")

	// Paddings inbetween both boxes and left and right of screen
	leftPadding := tview.NewBox()
	middlePadding := tview.NewBox()
	rightPadding := tview.NewBox()

	// Main body

	mainBody := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(leftPadding, 0, 1, false).
		AddItem(schoolInfo, 0, 3, false).
		AddItem(middlePadding, 0, 1, false).
		AddItem(courseInfo, 0, 7, false).
		AddItem(rightPadding, 0, 1, false)

	// Footer below calendar and mini daily schedule menu
	footer := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true).
		SetText("FOOTER FOOTER FOOTER")

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
