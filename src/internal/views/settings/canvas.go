package settings

import (
	"VIO/internal/store"
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func CanvasSettingsPage(
	app *tview.Application,
	current store.CanvasConfig,
	status string,
	onSave func(store.CanvasConfig),
	onCancel func(),
) tview.Primitive {
	domainField := tview.NewInputField().SetLabel("Canvas Domain: ")
	domainField.SetText(current.Domain)

	tokenField := tview.NewInputField().SetLabel("Canvas API Key: ")
	tokenField.SetMaskCharacter('*')
	tokenField.SetText(current.Token)

	pollField := tview.NewInputField().SetLabel("Poll Minutes: ")
	if current.PollMinutes <= 0 {
		current.PollMinutes = 5
	}
	pollField.SetText(fmt.Sprintf("%d", current.PollMinutes))

	statusBox := tview.NewTextView().SetDynamicColors(true)
	statusBox.SetBorder(true)
	statusBox.SetTitle("[ STATUS ]")
	if strings.TrimSpace(status) == "" {
		status = "Canvas setup is optional.\nLeave the fields blank if you want local JSON only."
	}
	statusBox.SetText(status)

	form := tview.NewForm()
	form.AddFormItem(domainField)
	form.AddFormItem(tokenField)
	form.AddFormItem(pollField)
	form.AddButton("Save", func() {
		cfg := store.CanvasConfig{
			Domain:      strings.TrimSpace(domainField.GetText()),
			Token:       strings.TrimSpace(tokenField.GetText()),
			PollMinutes: parsePollMinutes(pollField.GetText()),
		}
		onSave(cfg)
	})
	form.AddButton("Skip/Back", onCancel)
	form.SetBorder(true)
	form.SetTitle("[ CANVAS SETTINGS ]")

	header := tview.NewTextView().SetDynamicColors(true)
	header.SetText("[white::b][ ESC ] Back\n\nCanvas setup is optional. You can skip it and keep using local data only.\nUse [white::b]Tab[::-] / [white::b]Shift+Tab[::-] to move between fields.")

	body := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(header, 5, 0, false).
		AddItem(form, 0, 2, true).
		AddItem(statusBox, 0, 1, false)

	body.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEsc {
			onCancel()
			return nil
		}
		return event
	})

	return body
}

func parsePollMinutes(text string) int {
	text = strings.TrimSpace(text)
	if text == "" {
		return 5
	}

	var minutes int
	_, err := fmt.Sscanf(text, "%d", &minutes)
	if err != nil || minutes <= 0 {
		return 5
	}

	return minutes
}
