package schedule

import (
	"VIO/internal/asciiart"
	"VIO/internal/model"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func SchedulePage(app *tview.Application, data *model.AppData, returnTo func()) tview.Primitive {

	// Header

	quitPadding := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true).
		SetText(`


		`)

	quitText := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignRight).
		SetText("[white::b][ ESC ] To RETURN TO MAIN")

	quit := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(quitPadding, 2, 1, false).
		AddItem(quitText, 0, 2, false)

	title := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)

	currMonth := time.Now().Month().String()
	currDay := strconv.Itoa(time.Now().Day())

	title.SetText(asciiart.GetMonthHeader(currMonth) + asciiart.GetDayHeader(currDay))

	header := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(quit, 0, 2, false). // narrow weight
		AddItem(title, 0, 6, false) // wider weight
	// Main body

	// Schedule Box

	scheduleBody := tview.NewTextView().SetDynamicColors(true)
	scheduleBody.SetText(renderUpcomingFeed(data))

	mainBody := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(tview.NewBox(), 0, 1, false).
		AddItem(scheduleBody, 0, 5, false).
		AddItem(tview.NewBox(), 0, 1, false)

	// Footer

	footer := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true).
		SetText("[ + ] ADD   [ - ] REMOVE   [ E ] EDIT")

	page := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(header, 0, 3, false).
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

func renderUpcomingFeed(data *model.AppData) string {
	type item struct {
		label string
		due   time.Time
	}

	now := time.Now()
	var todayItems []item
	var pastItems []item

	for _, assignment := range data.Assignments {
		if assignment.DueAt == nil {
			continue
		}

		due := assignment.DueAt.In(time.Local)
		entry := item{
			label: fmt.Sprintf("Assignment: %s", assignment.Name),
			due:   due,
		}

		if sameLocalDay(due, now) {
			todayItems = append(todayItems, entry)
		} else if due.Before(now) {
			pastItems = append(pastItems, entry)
		}
	}

	sort.Slice(todayItems, func(i, j int) bool {
		return todayItems[i].due.Before(todayItems[j].due)
	})
	sort.Slice(pastItems, func(i, j int) bool {
		return pastItems[i].due.Before(pastItems[j].due)
	})

	allItems := append(todayItems, pastItems...)

	if len(allItems) == 0 {
		return "No dated items loaded yet."
	}

	var lines []string
	for _, it := range todayItems {
		lines = append(lines, fmt.Sprintf("%s\n%s", it.label, it.due.Format("Mon Jan 2, 3:04 PM")))
	}

	return strings.Join(lines, "\n\n")
}

func sameLocalDay(a, b time.Time) bool {
	la := a.In(time.Local)
	lb := b.In(time.Local)
	ay, am, ad := la.Date()
	by, bm, bd := lb.Date()
	return ay == by && am == bm && ad == bd
}
