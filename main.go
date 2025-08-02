package main

import (
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
	w_calendar := tview.NewBox().SetBorder(true).SetTitle("Left (1/2 x width of Top)")
	w_calendar_fixed_size := 0
	w_calendar_proportion := 2
	w_calendar_focus := true // App to be focues on calendar widget on startup

	// School Courses App
	w_courses := tview.NewBox().SetBorder(true).SetTitle("Top")
	w_courses_fixed_size := 0
	w_courses_proportion := 1
	w_courses_focus := false

	// Todo List App
	w_todo := tview.NewBox().SetBorder(true).SetTitle("Top")
	w_todo_fixed_size := 0
	w_todo_proportion := 1
	w_todo_focus := false

	// Schedule App
	w_schedule := tview.NewBox().SetBorder(true).SetTitle("Top")
	w_schedule_fixed_size := 0
	w_schedule_proportion := 1
	w_schedule_focus := false

	// Assignments List App
	w_assignments := tview.NewBox().SetBorder(true).SetTitle("Top")
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

	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}
}
