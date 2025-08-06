package widgets

import (
	"github.com/rivo/tview"
)

// BuildMainWidgets returns the list of box widgets and the main layout.
func BuildMainWidgets() ([]tview.Primitive, *tview.Flex) {

	// Widgets

	// Calendar
	w_calendar := tview.NewBox().SetBorder(true).SetTitle("[ 1 ]")
	w_calendar_fixed_size := 0
	w_calendar_proportion := 2
	w_calendar_focus := false

	// Courses
	w_courses := tview.NewBox().SetBorder(true).SetTitle("[ 2 ]")
	w_courses_fixed_size := 0
	w_courses_proportion := 1
	w_courses_focus := false

	// Todo/Tasks
	w_todo := tview.NewBox().SetBorder(true).SetTitle("[ 3 ]")
	w_todo_fixed_size := 0
	w_todo_proportion := 1
	w_todo_focus := false

	// Daily Schedule
	w_schedule := tview.NewBox().SetBorder(true).SetTitle("[ 4 ]")
	w_schedule_fixed_size := 0
	w_schedule_proportion := 1
	w_schedule_focus := false

	// Assignments
	w_assignments := tview.NewBox().SetBorder(true).SetTitle("[ 5 ]")
	w_assignments_fixed_size := 0
	w_assignments_proportion := 1
	w_assignments_focus := false

	// Layout structure
	mainBody := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewBox(), 1, 3, false).
		AddItem(
			tview.NewFlex().SetDirection(tview.FlexColumn).
				AddItem(w_calendar, w_calendar_fixed_size, w_calendar_proportion, w_calendar_focus).
				AddItem(w_courses, w_courses_fixed_size, w_courses_proportion, w_courses_focus).
				AddItem(w_todo, w_todo_fixed_size, w_todo_proportion, w_todo_focus),
			0, 1, false).
		AddItem(
			tview.NewFlex().SetDirection(tview.FlexColumn).
				AddItem(w_schedule, w_schedule_fixed_size, w_schedule_proportion, w_schedule_focus).
				AddItem(w_assignments, w_assignments_fixed_size, w_assignments_proportion, w_assignments_focus),
			0, 1, false)

		// Header

	quitPadding := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true).
		SetText(`


		`)

	quitText := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter).
		SetText("[white::b][ ESC ] To QUIT")

	quit := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(quitPadding, 2, 1, false).
		AddItem(quitText, 0, 2, false)

	title := tview.NewTextView().
		SetTextAlign(tview.AlignLeft).
		SetDynamicColors(true).
		SetText(`
__     __   _     ____    
\ \   / /  | |   / __ \ 
 \ \_/ /   | |  | |__| |
  \___/    |_|   \____/ 
`)

	header := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(quit, 0, 4, false). // narrow weight
		AddItem(title, 0, 6, false) // wider weight

	// Final layout with header and body
	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(header, 5, 0, false).
		AddItem(mainBody, 0, 1, false)

	widgets := []tview.Primitive{w_calendar, w_courses, w_todo, w_schedule, w_assignments}
	return widgets, flex
}
