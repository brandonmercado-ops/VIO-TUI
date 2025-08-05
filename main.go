package main

import (
	"VIO/pages"
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {

	// Create new TUI Application
	app := tview.NewApplication()
	// -----------------------------------------------------------------------------------------------
	//				WIDGET VARIABLES AND SETTINGS
	// -----------------------------------------------------------------------------------------------

	// Create widgets for each individual page view
	// !! widget var names are prefixed with "w_"

	// Calendar App
	w_calendar := tview.NewBox().SetBorder(true).SetTitle("[ 1 ]")
	w_calendar_fixed_size := 0
	w_calendar_proportion := 2
	w_calendar_focus := true // App to be focues on calendar widget on startup

	// School Courses App
	w_courses := tview.NewBox().SetBorder(true).SetTitle("[ 2 ]")
	w_courses_fixed_size := 0
	w_courses_proportion := 1
	w_courses_focus := false

	// Todo List App
	w_todo := tview.NewBox().SetBorder(true).SetTitle("[ 3 ]")
	w_todo_fixed_size := 0
	w_todo_proportion := 1
	w_todo_focus := false

	// Schedule App
	w_schedule := tview.NewBox().SetBorder(true).SetTitle("[ 4 ]")
	w_schedule_fixed_size := 0
	w_schedule_proportion := 1
	w_schedule_focus := false

	// Assignments List App
	w_assignments := tview.NewBox().SetBorder(true).SetTitle("[ 5 ]")
	w_assignments_fixed_size := 0
	w_assignments_proportion := 1
	w_assignments_focus := false
	// -----------------------------------------------------------------------------------------------
	//				    MAIN FUNCTIONAL WIDGETS
	// -----------------------------------------------------------------------------------------------

	// Arrange all windows accordingly
	mainBody := tview.NewFlex().SetDirection(tview.FlexRow).

		// Adding padding above to make space for Title
		AddItem(tview.NewBox(), 1, 3, false).
		AddItem(
			tview.NewFlex().SetDirection(tview.FlexColumn).
				AddItem(w_calendar, w_calendar_fixed_size, w_calendar_proportion, w_calendar_focus).
				AddItem(w_courses, w_courses_fixed_size, w_courses_proportion, w_courses_focus).
				AddItem(w_todo, w_todo_fixed_size, w_todo_proportion, w_todo_focus), 0, 1, false).
		AddItem(
			tview.NewFlex().SetDirection(tview.FlexColumn).
				AddItem(w_schedule, w_schedule_fixed_size, w_schedule_proportion, w_schedule_focus).
				AddItem(w_assignments, w_assignments_fixed_size, w_assignments_proportion, w_assignments_focus), 0, 1, false)
	// -----------------------------------------------------------------------------------------------
	//				    HEADER TITLE SECTION
	// -----------------------------------------------------------------------------------------------

	vioTitle := `
	 __     __   _     ____    
	\ \   / /  | |   / __ \ 
	 \ \_/ /   | |  | |__| |
	  \___/    |_|   \____/ 
	`

	titleBanner := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetText(vioTitle).
		SetDynamicColors(true)
		// SetTextColor(tcell.ColorLightCyan)

	// -----------------------------------------------------------------------------------------------
	//				   COMBINING HEADER AND MAIN BODY
	// -----------------------------------------------------------------------------------------------

	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(titleBanner, 5, 0, false).
		AddItem(mainBody, 0, 1, false)

	// -----------------------------------------------------------------------------------------------
	//				SWITCH SCREENS UPON WIDGET-CLICK
	// -----------------------------------------------------------------------------------------------

	// Slice of widgets for indexed access
	widgets := []tview.Primitive{
		w_calendar,    // index 0 = key '1'
		w_courses,     // index 1 = key '2'
		w_todo,        // index 2 = key '3'
		w_schedule,    // index 3 = key '4'
		w_assignments, // index 4 = key '5'
	}

	openScreen := func(index int) {
		var screen tview.Primitive

		switch index {
		case 0:
			screen = pages.CalendarPage(app, func() {
				app.SetRoot(flex, true).SetFocus(widgets[index])
			})
		case 1:
			screen = pages.CoursesPage(app, func() {
				app.SetRoot(flex, true).SetFocus(widgets[index])
			})
		case 2:
			screen = pages.TasksPage(app, func() {
				app.SetRoot(flex, true).SetFocus(widgets[index])
			})
		case 3:
			screen = pages.SchedulePage(app, func() {
				app.SetRoot(flex, true).SetFocus(widgets[index])
			})
		case 4:
			screen = pages.AssignmentsPage(app, func() {
				app.SetRoot(flex, true).SetFocus(widgets[index])
			})

		default:
			tv := tview.NewTextView().
				SetTextAlign(tview.AlignCenter).
				SetDynamicColors(true).
				SetText(fmt.Sprintf("🌟 You selected [::b]Box %d[::-]!\n\n[gray]Press Esc to return", index+1))

			tv.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
				if event.Key() == tcell.KeyEsc {
					app.SetRoot(flex, true).SetFocus(widgets[index])
					return nil
				}
				return event
			})

			screen = tv
		}

		app.SetRoot(screen, true).SetFocus(screen)
	}

	// -----------------------------------------------------------------------------------------------
	//				HANDLING WIDGET HIGHLIGHT AND NAVIGATION
	// -----------------------------------------------------------------------------------------------

	// Focused on Calendar by default (upon entering application)
	focusIndex := 0

	// Anonymous function to change both focus and color of chosen widget
	updateFocus := func(index int) {

		// Change colors for highlighted widget and reset color when out of focus
		for i, w := range widgets {
			if box, ok := w.(*tview.Box); ok {
				if i == index {
					box.SetBorderColor(tcell.ColorSpringGreen)
					box.SetTitleColor(tcell.ColorSpringGreen)
					box.SetBorderAttributes(tcell.AttrBold)
				} else {
					box.SetBorderColor(tcell.ColorWhite)
					box.SetTitleColor(tcell.ColorWhite)
					box.SetBorderAttributes(tcell.AttrNone)
				}
			}
		}
	}

	// Focusing on other widgets indicated by their corresponding numbers
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyRune:
			switch event.Rune() {
			case '1', '2', '3', '4', '5':
				focusIndex = int(event.Rune() - '1')
				app.SetFocus(widgets[focusIndex])
				updateFocus(focusIndex)
			}
		case tcell.KeyEnter:
			openScreen(focusIndex)
		}

		return event
	})
	// -----------------------------------------------------------------------------------------------
	//				START-OF-PROGRAM ERROR HANDLING
	// -----------------------------------------------------------------------------------------------

	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}
}
