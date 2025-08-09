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
// Bulbhead font

var MonthHeaders = map[string]string{
	"January": `
  ____   __    _  _  __  __    __    ____  _  _ 
 (_  _) /__\  ( \( )(  )(  )  /__\  (  _ \( \/ )
.-_)(  /(__)\  )  (  )(__)(  /(__)\  )   / \  / 
\____)(__)(__)(_)\_)(______)(__)(__)(_)\_) (__) `,
	"February": `
 ____  ____  ____  ____  __  __    __    ____  _  _ 
( ___)( ___)(  _ \(  _ \(  )(  )  /__\  (  _ \( \/ )
 )__)  )__)  ) _ < )   / )(__)(  /(__)\  )   / \  / 
(__)  (____)(____/(_)\_)(______)(__)(__)(_)\_) (__) `,
	"March": `
 __  __    __    ____   ___  _   _ 
(  \/  )  /__\  (  _ \ / __)( )_( )
 )    (  /(__)\  )   /( (__  ) _ ( 
(_/\/\_)(__)(__)(_)\_) \___)(_) (_)`,
	"April": `
   __    ____  ____  ____  __   
  /__\  (  _ \(  _ \(_  _)(  )  
 /(__)\  )___/ )   / _)(_  )(__ 
(__)(__)(__)  (_)\_)(____)(____)`,
	"May": `
 __  __    __   _  _ 
(  \/  )  /__\ ( \/ )
 )    (  /(__)\ \  / 
(_/\/\_)(__)(__)(__) `,
	"June": `
  ____  __  __  _  _  ____ 
 (_  _)(  )(  )( \( )( ___)
.-_)(   )(__)(  )  (  )__) 
\____) (______)(_)\_)(____)`,
	"July": `
  ____  __  __  __   _  _ 
 (_  _)(  )(  )(  ) ( \/ )
.-_)(   )(__)(  )(__ \  / 
\____) (______)(____)(__) 
`,
	"August": `
   __    __  __  ___  __  __  ___  ____ 
  /__\  (  )(  )/ __)(  )(  )/ __)(_  _)
 /(__)\  )(__)(( (_-. )(__)( \__ \  )(  
(__)(__)(______)\___/(______)(___/ (__) `,
	"September": `
 ___  ____  ____  ____  ____  __  __  ____  ____  ____ 
/ __)( ___)(  _ \(_  _)( ___)(  \/  )(  _ \( ___)(  _ \
\__ \ )__)  )___/  )(   )__)  )    (  ) _ < )__)  )   /
(___/(____)(__)   (__) (____)(_/\/\_)(____/(____)(_)\_)`,
	"October": `
 _____  ___  ____  _____  ____  ____  ____ 
(  _  )/ __)(_  _)(  _  )(  _ \( ___)(  _ \
 )(_)(( (__   )(   )(_)(  ) _ < )__)  )   /
(_____)\___) (__) (_____)(____/(____)(_)\_)
`,
	"November": `
 _  _  _____  _  _  ____  __  __  ____  ____  ____ 
( \( )(  _  )( \/ )( ___)(  \/  )(  _ \( ___)(  _ \
 )  (  )(_)(  \  /  )__)  )    (  ) _ < )__)  )   /
(_)\_)(_____)  \/  (____)(_/\/\_)(____/(____)(_)\_)`,
	"December": `
 ____  ____  ___  ____  __  __  ____  ____  ____ 
(  _ \( ___)/ __)( ___)(  \/  )(  _ \( ___)(  _ \
 )(_) ))__)( (__  )__)  )    (  ) _ < )__)  )   /
(____/(____)\___)(____)(_/\/\_)(____/(____)(_)\_)`,
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
