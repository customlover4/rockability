package tui

import (
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (a *App) ShowMessage(title, text string) {
	show := func() {
		modal := tview.NewModal().
			SetText(text).
			AddButtons([]string{"OK"}).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				a.pages.RemovePage("modal")
				a.app.SetFocus(a.actionsList)
				a.dequeuePopup()
			})
		modal.SetBackgroundColor(a.theme.ModalBackground)
		modal.SetTextColor(a.theme.PanelText)
		modal.SetButtonBackgroundColor(a.theme.PanelBackground)
		modal.SetButtonTextColor(a.theme.PanelText)
		modal.SetBorder(true)
		modal.SetBorderColor(a.theme.Border)
		modal.SetTitleColor(a.theme.Title)
		modal.SetTitle(title)
		a.pages.AddPage("modal", modal, true, true)
		a.app.SetFocus(modal)
	}

	a.enqueuePopup(show)
}

func (a *App) ShowChoicePopup(title, text string, options []ChoiceOption) {
	if len(options) == 0 {
		return
	}

	show := func() {
		description := tview.NewTextView().
			SetDynamicColors(true).
			SetWrap(true).
			SetWordWrap(true)
		description.SetText(text)
		description.SetBorder(true)
		description.SetTitle("Описание")
		description.SetBackgroundColor(a.theme.ModalBackground)
		description.SetTextColor(a.theme.PanelText)
		description.SetBorderColor(a.theme.Border)
		description.SetTitleColor(a.theme.Title)

		list := tview.NewList().ShowSecondaryText(false)
		list.SetBackgroundColor(a.theme.PanelBackground)
		list.SetMainTextColor(a.theme.PanelText)
		list.SetSelectedTextColor(a.theme.ListSelectedText)
		list.SetSelectedBackgroundColor(a.theme.ListSelectedBg)
		list.SetHighlightFullLine(true)

		closePopup := func() {
			a.pages.RemovePage("popup")
			a.app.SetFocus(a.actionsList)
			a.dequeuePopup()
		}

		for _, option := range options {
			selected := option
			list.AddItem(selected.Label, "", 0, func() {
				closePopup()
				if selected.Handle != nil {
					selected.Handle(a)
				}
			})
		}

		list.SetBorder(true).SetTitle(title)
		list.SetBorderColor(a.theme.Border)
		list.SetTitleColor(a.theme.Title)

		content := tview.NewFlex().SetDirection(tview.FlexRow)
		content.SetBackgroundColor(a.theme.ModalBackground)
		content.AddItem(description, 5, 0, false)
		content.AddItem(list, 0, 1, true)

		content.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			switch event.Key() {
			case tcell.KeyEscape:
				closePopup()
				return nil
			}
			return event
		})

		column := tview.NewFlex().SetDirection(tview.FlexRow)
		column.SetBackgroundColor(a.theme.Background)
		column.AddItem(nil, 0, 1, false)
		column.AddItem(content, 16, 1, true)
		column.AddItem(nil, 0, 1, false)

		popup := tview.NewFlex()
		popup.SetBackgroundColor(a.theme.Background)
		popup.AddItem(nil, 0, 1, false)
		popup.AddItem(column, 68, 1, true)
		popup.AddItem(nil, 0, 1, false)

		a.pages.AddPage("popup", popup, true, true)
		a.app.SetFocus(list)
	}

	a.enqueuePopup(show)
}

func (a *App) ShowFormPopup(title string, fields []FormField, onSubmit func(values map[string]string)) {
	if len(fields) == 0 {
		return
	}

	show := func() {
		form := tview.NewForm()
		inputs := make([]*tview.InputField, 0, len(fields))
		promptFocus := 0
		totalFocusItems := len(fields) + 2

		setPromptFocus := func(index int) {
			if index < 0 {
				index = totalFocusItems - 1
			}
			if index >= totalFocusItems {
				index = 0
			}
			promptFocus = index
			form.SetFocus(index)
			a.app.SetFocus(form)
		}

		closePopup := func() {
			a.pages.RemovePage("popup")
			a.app.SetFocus(a.actionsList)
			a.dequeuePopup()
		}

		for idx, field := range fields {
			input := tview.NewInputField().SetLabel(field.Label + ": ")
			input.SetPlaceholder(field.Placeholder)
			input.SetText(field.Value)
			input.SetFieldBackgroundColor(a.theme.Background)
			input.SetFieldTextColor(a.theme.PanelText)
			input.SetLabelColor(a.theme.Title)

			nextIndex := idx + 1
			input.SetDoneFunc(func(key tcell.Key) {
				switch key {
				case tcell.KeyEnter, tcell.KeyTab:
					setPromptFocus(nextIndex)
				case tcell.KeyEscape:
					closePopup()
				}
			})

			inputs = append(inputs, input)
			form.AddFormItem(input)
		}

		form.AddButton("Сохранить", func() {
			values := make(map[string]string, len(fields))
			for i, field := range fields {
				value := strings.TrimSpace(inputs[i].GetText())
				if value == "" {
					value = field.DefaultValue
				}
				values[field.Key] = value
			}
			closePopup()
			if onSubmit != nil {
				onSubmit(values)
			}
		})

		form.AddButton("Отмена", func() {
			closePopup()
		})

		form.SetBorder(true).SetTitle(title)
		form.SetButtonsAlign(tview.AlignCenter)
		form.SetBackgroundColor(a.theme.ModalBackground)
		form.SetBorderColor(a.theme.Border)
		form.SetTitleColor(a.theme.Title)
		form.SetButtonBackgroundColor(a.theme.PanelBackground)
		form.SetButtonTextColor(a.theme.PanelText)
		form.SetLabelColor(a.theme.Title)
		form.SetFieldBackgroundColor(a.theme.Background)
		form.SetFieldTextColor(a.theme.PanelText)
		form.SetCancelFunc(closePopup)
		form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			switch event.Key() {
			case tcell.KeyTab, tcell.KeyRight, tcell.KeyDown:
				setPromptFocus(promptFocus + 1)
				return nil
			case tcell.KeyBacktab, tcell.KeyLeft, tcell.KeyUp:
				setPromptFocus(promptFocus - 1)
				return nil
			}
			return event
		})

		column := tview.NewFlex().SetDirection(tview.FlexRow)
		column.SetBackgroundColor(a.theme.Background)
		column.AddItem(nil, 0, 1, false)
		column.AddItem(form, len(fields)+8, 1, true)
		column.AddItem(nil, 0, 1, false)

		popup := tview.NewFlex()
		popup.SetBackgroundColor(a.theme.Background)
		popup.AddItem(nil, 0, 1, false)
		popup.AddItem(column, 62, 1, true)
		popup.AddItem(nil, 0, 1, false)

		a.pages.AddPage("popup", popup, true, true)
		setPromptFocus(0)
	}
	
	a.enqueuePopup(show)
}
