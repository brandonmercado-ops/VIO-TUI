package tasks

import (
	"VIO/internal/model"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func TasksPage(app *tview.Application, data *model.AppData, returnTo func()) tview.Primitive {

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

	// TODO, IN PROGRESS, and DONE task sections

	inProgress := tview.NewTextView().SetDynamicColors(true)
	inProgress.SetBorder(true)
	inProgress.SetTitle("[ IN PROGRESS ]")

	overdue := tview.NewTextView().SetDynamicColors(true)
	overdue.SetBorder(true)
	overdue.SetTitle("[ OVERDUE ]")

	complete := tview.NewTextView().SetDynamicColors(true)
	complete.SetBorder(true)
	complete.SetTitle("[ COMPLETE ]")

	inProgress.SetText(renderTaskColumn(data.Tasks, "in_progress"))
	overdue.SetText(renderTaskColumn(data.Tasks, "overdue"))
	complete.SetText(renderTaskColumn(data.Tasks, "complete"))

	// Main body with paddings

	mainBody := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(tview.NewBox(), 0, 1, false).
		AddItem(inProgress, 0, 8, false).
		AddItem(tview.NewBox(), 0, 1, false).
		AddItem(overdue, 0, 8, false).
		AddItem(tview.NewBox(), 0, 1, false).
		AddItem(complete, 0, 7, false).
		AddItem(tview.NewBox(), 0, 1, false)

	// Footer below all 3 task columns
	footer := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true).
		SetText("[ + ] ADD   [ - ] REMOVE   [ E ] EDIT")

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

func renderTaskColumn(tasks []model.Task, status string) string {
	filtered := make([]model.Task, 0)
	for _, task := range tasks {
		if effectiveTaskStatus(task) == status {
			filtered = append(filtered, task)
		}
	}

	sort.Slice(filtered, func(i, j int) bool {
		left := filtered[i].DueAt
		right := filtered[j].DueAt
		if left == nil {
			return false
		}
		if right == nil {
			return true
		}
		return left.Before(*right)
	})

	var lines []string
	for _, task := range filtered {
		line := task.Title
		if task.CourseID != "" {
			line += fmt.Sprintf("\n[gray]%s[-]", task.CourseID)
		}
		if task.DueAt != nil {
			line += fmt.Sprintf("\nDue: %s", task.DueAt.Format("Mon Jan 2, 3:04 PM"))
		}
		lines = append(lines, line)
	}

	if len(lines) == 0 {
		return "No tasks here yet."
	}

	return strings.Join(lines, "\n\n")
}

func effectiveTaskStatus(task model.Task) string {
	status := strings.ToLower(strings.TrimSpace(task.Status))
	if status == "complete" {
		return "complete"
	}
	if task.DueAt != nil && task.DueAt.Before(time.Now()) {
		return "overdue"
	}
	return "in_progress"
}
