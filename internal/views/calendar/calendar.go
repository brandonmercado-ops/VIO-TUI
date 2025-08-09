package calendar

import (
	// "fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"time"
)

//-----------------------------------------------------------------------------------
//                            CALENDAR MONTH HEADERS
//-----------------------------------------------------------------------------------

// https://patorjk.com/software/taag/#p=display&v=2&f=Bulbhead&t=
// ----Bulbhead----font---- (old)
// DiamFont ( NEW CURRENT FONT )

var MonthHeaders = map[string]string{
	"January": `
   ▗▖ ▗▄▖ ▗▖  ▗▖▗▖ ▗▖ ▗▄▖ ▗▄▄▖▗▖  ▗▖
   ▐▌▐▌ ▐▌▐▛▚▖▐▌▐▌ ▐▌▐▌ ▐▌▐▌ ▐▌▝▚▞▘ 
   ▐▌▐▛▀▜▌▐▌ ▝▜▌▐▌ ▐▌▐▛▀▜▌▐▛▀▚▖ ▐▌  
▗▄▄▞▘▐▌ ▐▌▐▌  ▐▌▝▚▄▞▘▐▌ ▐▌▐▌ ▐▌ ▐▌  
                                    `,
	"February": `
▗▄▄▄▖▗▄▄▄▖▗▄▄▖ ▗▄▄▖ ▗▖ ▗▖ ▗▄▖ ▗▄▄▖▗▖  ▗▖
▐▌   ▐▌   ▐▌ ▐▌▐▌ ▐▌▐▌ ▐▌▐▌ ▐▌▐▌ ▐▌▝▚▞▘ 
▐▛▀▀▘▐▛▀▀▘▐▛▀▚▖▐▛▀▚▖▐▌ ▐▌▐▛▀▜▌▐▛▀▚▖ ▐▌  
▐▌   ▐▙▄▄▖▐▙▄▞▘▐▌ ▐▌▝▚▄▞▘▐▌ ▐▌▐▌ ▐▌ ▐▌  
                                        `,
	"March": `
▗▖  ▗▖ ▗▄▖ ▗▄▄▖  ▗▄▄▖▗▖ ▗▖
▐▛▚▞▜▌▐▌ ▐▌▐▌ ▐▌▐▌   ▐▌ ▐▌
▐▌  ▐▌▐▛▀▜▌▐▛▀▚▖▐▌   ▐▛▀▜▌
▐▌  ▐▌▐▌ ▐▌▐▌ ▐▌▝▚▄▄▖▐▌ ▐▌
                          `,
	"April": `
 ▗▄▖ ▗▄▄▖ ▗▄▄▖ ▗▄▄▄▖▗▖   
▐▌ ▐▌▐▌ ▐▌▐▌ ▐▌  █  ▐▌   
▐▛▀▜▌▐▛▀▘ ▐▛▀▚▖  █  ▐▌   
▐▌ ▐▌▐▌   ▐▌ ▐▌▗▄█▄▖▐▙▄▄▖
                         `,
	"May": `
▗▖  ▗▖ ▗▄▖▗▖  ▗▖
▐▛▚▞▜▌▐▌ ▐▌▝▚▞▘ 
▐▌  ▐▌▐▛▀▜▌ ▐▌  
▐▌  ▐▌▐▌ ▐▌ ▐▌  
                `,
	"June": `
   ▗▖▗▖ ▗▖▗▖  ▗▖▗▄▄▄▖
   ▐▌▐▌ ▐▌▐▛▚▖▐▌▐▌   
   ▐▌▐▌ ▐▌▐▌ ▝▜▌▐▛▀▀▘
▗▄▄▞▘▝▚▄▞▘▐▌  ▐▌▐▙▄▄▖
                     `,
	"July": `
   ▗▖▗▖ ▗▖▗▖ ▗▖  ▗▖
   ▐▌▐▌ ▐▌▐▌  ▝▚▞▘ 
   ▐▌▐▌ ▐▌▐▌   ▐▌  
▗▄▄▞▘▝▚▄▞▘▐▙▄▄▖▐▌  
                   `,
	"August": `
 ▗▄▖ ▗▖ ▗▖ ▗▄▄▖▗▖ ▗▖ ▗▄▄▖▗▄▄▄▖
▐▌ ▐▌▐▌ ▐▌▐▌   ▐▌ ▐▌▐▌     █  
▐▛▀▜▌▐▌ ▐▌▐▌▝▜▌▐▌ ▐▌ ▝▀▚▖  █  
▐▌ ▐▌▝▚▄▞▘▝▚▄▞▘▝▚▄▞▘▗▄▄▞▘  █  
                              `,
	"September": `
 ▗▄▄▖▗▄▄▄▖▗▄▄▖▗▄▄▄▖▗▄▄▄▖▗▖  ▗▖▗▄▄▖ ▗▄▄▄▖▗▄▄▖ 
▐▌   ▐▌   ▐▌ ▐▌ █  ▐▌   ▐▛▚▞▜▌▐▌ ▐▌▐▌   ▐▌ ▐▌
 ▝▀▚▖▐▛▀▀▘▐▛▀▘  █  ▐▛▀▀▘▐▌  ▐▌▐▛▀▚▖▐▛▀▀▘▐▛▀▚▖
▗▄▄▞▘▐▙▄▄▖▐▌    █  ▐▙▄▄▖▐▌  ▐▌▐▙▄▞▘▐▙▄▄▖▐▌ ▐▌
                                             `,
	"October": `
 ▗▄▖  ▗▄▄▖▗▄▄▄▖▗▄▖ ▗▄▄▖ ▗▄▄▄▖▗▄▄▖ 
▐▌ ▐▌▐▌     █ ▐▌ ▐▌▐▌ ▐▌▐▌   ▐▌ ▐▌
▐▌ ▐▌▐▌     █ ▐▌ ▐▌▐▛▀▚▖▐▛▀▀▘▐▛▀▚▖
▝▚▄▞▘▝▚▄▄▖  █ ▝▚▄▞▘▐▙▄▞▘▐▙▄▄▖▐▌ ▐▌
                                  `,
	"November": `
▗▖  ▗▖ ▗▄▖ ▗▖  ▗▖▗▄▄▄▖▗▖  ▗▖▗▄▄▖ ▗▄▄▄▖▗▄▄▖ 
▐▛▚▖▐▌▐▌ ▐▌▐▌  ▐▌▐▌   ▐▛▚▞▜▌▐▌ ▐▌▐▌   ▐▌ ▐▌
▐▌ ▝▜▌▐▌ ▐▌▐▌  ▐▌▐▛▀▀▘▐▌  ▐▌▐▛▀▚▖▐▛▀▀▘▐▛▀▚▖
▐▌  ▐▌▝▚▄▞▘ ▝▚▞▘ ▐▙▄▄▖▐▌  ▐▌▐▙▄▞▘▐▙▄▄▖▐▌ ▐▌
                                           `,
	"December": `
▗▄▄▄  ▗▄▄▄▖ ▗▄▄▖▗▄▄▄▖▗▖  ▗▖▗▄▄▖ ▗▄▄▄▖▗▄▄▖ 
▐▌  █ ▐▌   ▐▌   ▐▌   ▐▛▚▞▜▌▐▌ ▐▌▐▌   ▐▌ ▐▌
▐▌  █ ▐▛▀▀▘▐▌   ▐▛▀▀▘▐▌  ▐▌▐▛▀▚▖▐▛▀▀▘▐▛▀▚▖
▐▙▄▄▀ ▐▙▄▄▖▝▚▄▄▖▐▙▄▄▖▐▌  ▐▌▐▙▄▞▘▐▙▄▄▖▐▌ ▐▌
                                          `,
}

func CalendarPage(app *tview.Application, returnTo func()) tview.Primitive {

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
		SetDynamicColors(true)

	currMonth := time.Now().Month().String()
	title.SetText(MonthHeaders[currMonth])

	header := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(quit, 0, 4, false). // narrow weight
		AddItem(title, 0, 6, false) // wider weight

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
		SetDynamicColors(true)

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
