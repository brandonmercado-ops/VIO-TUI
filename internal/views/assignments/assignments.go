package assignments

import (
	"VIO/internal/model"
	"fmt"
	"sort"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func AssignmentsPage(app *tview.Application, data *model.AppData, returnTo func()) tview.Primitive {

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

	// Main body

	assignmentsBody := tview.NewTextView().SetDynamicColors(true)
	assignmentsBody.SetText(renderAssignments(data))

	mainBody := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(tview.NewBox(), 0, 1, false).
		AddItem(assignmentsBody, 0, 6, false).
		AddItem(tview.NewBox(), 0, 1, false)

	// Footer

	footer := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true).
		SetText("[ + ] ADD   [ - ] REMOVE   [ E ] EDIT")

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

func renderAssignments(data *model.AppData) string {
	assignments := append([]model.Assignment(nil), data.Assignments...)

	sort.Slice(assignments, func(i, j int) bool {
		left := assignments[i].DueAt
		right := assignments[j].DueAt

		if left == nil {
			return false
		}
		if right == nil {
			return true
		}

		return left.Before(*right)
	})

	if len(assignments) == 0 {
		return "No assignments loaded yet."
	}

	var lines []string
	for _, assignment := range assignments {
		courseLabel := assignment.CourseID
		for _, course := range data.Courses {
			if course.ID == assignment.CourseID {
				courseLabel = course.Code
				break
			}
		}

		statusParts := []string{}
		if assignment.HasSubmitted {
			statusParts = append(statusParts, "submitted")
		}
		if assignment.IsMissing {
			statusParts = append(statusParts, "missing")
		}
		if assignment.IsLate {
			statusParts = append(statusParts, "late")
		}
		if len(statusParts) == 0 {
			statusParts = append(statusParts, "open")
		}

		line := fmt.Sprintf("[white::b]%s[::-]  [gray](%s)[-]", assignment.Name, courseLabel)

		if assignment.DueAt != nil {
			line += fmt.Sprintf("\nDue: %s", assignment.DueAt.Format("Mon Jan 2, 3:04 PM"))
		}

		line += fmt.Sprintf("\nStatus: %s", strings.Join(statusParts, ", "))

		if assignment.PointsPossible > 0 {
			line += fmt.Sprintf("\nPoints: %.0f", assignment.PointsPossible)
		}

		lines = append(lines, line)
	}

	return strings.Join(lines, "\n\n")
}
