package tui

import (
	"fmt"
	"math/rand"
	"strings"
	"sync/atomic"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	HapinessIdx = iota
	HealthIdx
	InspirationIdx
)

type App struct {
	app            *tview.Application
	pages          *tview.Pages
	theme          Theme
	rng            *rand.Rand
	currentActions []Action
	cutsceneActive atomic.Bool
	onQuit         func()

	screen Screen

	description *tview.TextView
	status      *tview.TextView
	actionsList *tview.List
	statsView   *tview.TextView

	popupQueue  []func()
	popupActive bool
}

func NewApp() *App {
	theme := DefaultTheme()
	applyTheme(theme)

	app := &App{
		app:   tview.NewApplication(),
		pages: tview.NewPages(),
		theme: theme,
		rng:   newRNG(),
		screen: Screen{
			ActionTitle: "Действия",
			LogTitle:    "Журнал",
			StatusTitle: "Статус",
			Log: []string{
				"UI framework готов.",
				"Заполни actions, log и status своими данными.",
			},
			Stats: []Stat{
				HapinessIdx: {
					Name:  "Счастье",
					Value: 75,
				},
				HealthIdx: {
					Name:  "Здоровье",
					Value: 100,
				},
				InspirationIdx: {
					Name:  "Вдохновение",
					Value: 50,
				},
			},
		},

		popupQueue:  make([]func(), 0, 4),
		popupActive: false,
	}

	app.buildMainLayout()
	app.installGlobalInput()

	return app
}

func (a *App) Run() error {
	return a.app.SetRoot(a.pages, true).SetFocus(a.actionsList).Run()
}

func (a *App) SetOnQuit(fn func()) {
	a.onQuit = fn
}

func (a *App) Update(fn func()) {
	a.app.QueueUpdateDraw(func() {
		fn()
		a.render()
	})
}

func (a *App) SetScreen(screen Screen) {
	a.screen = screen
	a.render()
}

func (a *App) Screen() Screen {
	return a.screen
}

func (a *App) SetActions(actions []Action) {
	a.screen.Actions = actions
	a.render()
}

func (a *App) UpdateStatusStats(status []StatusItem, st []Stat) {
	a.screen.Status = append([]StatusItem(nil), status...)
	a.screen.Stats = append([]Stat(nil), st...)
	a.render()
}

func (a *App) SetStats(stats []Stat) {
	a.screen.Stats = append([]Stat(nil), stats...)
	a.render()
}

func (a *App) UpdateStat(index int, value int) {
	if index < 0 || index >= len(a.screen.Stats) {
		return
	}
	if value < 0 {
		value = 0
	}
	if value > 100 {
		value = 100
	}
	a.screen.Stats[index].Value = value
	a.render()
}

func (a *App) SetLog(lines []string) {
	a.screen.Log = append([]string(nil), lines...)
	a.render()
}

func (a *App) AppendLog(lines ...string) {
	a.screen.Log = append(a.screen.Log, lines...)
	a.render()
}

func (a *App) ClearLog() {
	a.screen.Log = nil
	a.render()
}

func (a *App) SetStatus(items []StatusItem) {
	a.screen.Status = append([]StatusItem(nil), items...)
	a.render()
}

func (a *App) SetTitles(actionTitle, logTitle, statusTitle string) {
	a.screen.ActionTitle = actionTitle
	a.screen.LogTitle = logTitle
	a.screen.StatusTitle = statusTitle
	a.render()
}

func (a *App) FocusActions() {
	a.app.SetFocus(a.actionsList)
}

func (a *App) enqueuePopup(showFunc func()) {
	if !a.popupActive {
		a.popupActive = true
		showFunc()
	} else {
		a.popupQueue = append(a.popupQueue, showFunc)
	}
}

func (a *App) dequeuePopup() {
	a.popupActive = false

	if len(a.popupQueue) > 0 {
		next := a.popupQueue[0]
		a.popupQueue = a.popupQueue[1:]
		a.popupActive = true
		next()
	}
}

func (a *App) isPopupActive() bool {
	return a.popupActive
}

func (a *App) buildMainLayout() {
	a.description = tview.NewTextView().
		SetDynamicColors(true).
		SetWrap(true).
		SetWordWrap(true)
	a.description.SetBorder(true)
	a.description.SetBackgroundColor(a.theme.PanelBackground)
	a.description.SetTextColor(a.theme.PanelText)
	a.description.SetBorderColor(a.theme.Border)
	a.description.SetTitleColor(a.theme.Title)

	a.status = tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignLeft)
	a.status.SetBorder(true)
	a.status.SetBackgroundColor(a.theme.PanelBackground)
	a.status.SetTextColor(a.theme.PanelText)
	a.status.SetBorderColor(a.theme.Border)
	a.status.SetTitleColor(a.theme.Title)

	a.actionsList = tview.NewList().ShowSecondaryText(false)
	a.actionsList.SetBorder(true)
	a.actionsList.SetBackgroundColor(a.theme.PanelBackground)
	a.actionsList.SetMainTextColor(a.theme.PanelText)
	a.actionsList.SetSelectedTextColor(a.theme.ListSelectedText)
	a.actionsList.SetSelectedBackgroundColor(a.theme.ListSelectedBg)
	a.actionsList.SetBorderColor(a.theme.Border)
	a.actionsList.SetTitleColor(a.theme.Title)

	a.statsView = tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignLeft).
		SetWordWrap(false)
	a.statsView.SetBorder(true)
	a.statsView.SetBackgroundColor(a.theme.PanelBackground)
	a.statsView.SetTextColor(a.theme.PanelText)
	a.statsView.SetBorderColor(a.theme.Border)
	a.statsView.SetTitleColor(a.theme.Title)
	a.statsView.SetTitle("Показатели")

	leftColumn := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(a.actionsList, 0, 3, true). // 3 части высоты — список действий
		AddItem(a.statsView, 0, 1, false)   // 1 часть высоты — показатели

	topLayout := tview.NewFlex().
		AddItem(leftColumn, 40, 1, true).   // фикс 40 символов ширина
		AddItem(a.description, 0, 3, false) // остальная ширина — лог

	mainLayout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(topLayout, 0, 1, true). // гибкая высота
		AddItem(a.status, 3, 0, false)  // фикс 3 строки внизу
	mainLayout.SetBackgroundColor(a.theme.Background)

	a.pages.AddPage("main", mainLayout, true, true)
	a.render()
}

func (a *App) render() {
	if a.screen.ActionTitle == "" {
		a.screen.ActionTitle = "Действия"
	}
	if a.screen.LogTitle == "" {
		a.screen.LogTitle = "Журнал"
	}
	if a.screen.StatusTitle == "" {
		a.screen.StatusTitle = "Статус"
	}

	a.actionsList.SetTitle(a.screen.ActionTitle)
	a.description.SetTitle(a.screen.LogTitle)
	a.status.SetTitle(a.screen.StatusTitle)

	a.renderActions()
	a.renderLog()
	a.renderStatus()
	a.renderStats()
}

func (a *App) renderActions() {
	current := a.actionsList.GetCurrentItem()
	a.actionsList.Clear()
	a.currentActions = append([]Action(nil), a.screen.Actions...)
	for i := range a.currentActions {
		index := i
		a.actionsList.AddItem(a.currentActions[i].Label, "", 0, func() {
			if a.currentActions[index].Handle != nil {
				a.currentActions[index].Handle(a)
			}
		})
	}

	if len(a.currentActions) == 0 {
		return
	}
	if current < 0 {
		current = 0
	}
	if current >= len(a.currentActions) {
		current = len(a.currentActions) - 1
	}
	a.actionsList.SetCurrentItem(current)
}

func (a *App) renderLog() {
	a.description.SetText(strings.Join(a.screen.Log, "\n\n"))
	a.description.ScrollToEnd()
}

func (a *App) renderStatus() {
	parts := make([]string, 0, len(a.screen.Status))
	for _, item := range a.screen.Status {
		if item.Label == "" {
			continue
		}
		parts = append(parts, fmt.Sprintf("[%s]%s:[%s] %s", a.theme.StatusAccent, item.Label, a.theme.StatusValue, item.Value))
	}
	a.status.SetText(strings.Join(parts, "   "))
}

func (a *App) renderStats() {
	if len(a.screen.Stats) == 0 {
		a.screen.Stats = []Stat{
			{Name: "Счастье", Value: 0},
			{Name: "Здоровье", Value: 0},
			{Name: "Вдохновение", Value: 0},
		}
	}

	lines := make([]string, 0, 3)
	for _, stat := range a.screen.Stats {
		lines = append(
			lines, fmt.Sprintf("%s: %d/100", stat.Name, stat.Value),
		)
	}
	a.statsView.SetText(strings.Join(lines, "\n"))
}

func (a *App) installGlobalInput() {
	a.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if a.cutsceneActive.Load() {
			return event
		}

		switch event.Key() {
		case tcell.KeyEscape:
			if a.pages.HasPage("popup") {
				a.pages.RemovePage("popup")
				a.app.SetFocus(a.actionsList)
				a.dequeuePopup()
				return nil
			}
			if a.pages.HasPage("modal") {
				a.pages.RemovePage("modal")
				a.app.SetFocus(a.actionsList)
				a.dequeuePopup()
				return nil
			}
			if a.onQuit != nil {
				a.onQuit()
			}
			a.app.Stop()
			return nil
		}

		switch event.Rune() {
		case 'q', 'Q':
			if a.onQuit != nil {
				a.onQuit()
			}
			a.app.Stop()
			return nil
		}

		return event
	})
}

func (a *App) Stop() {
	a.app.Stop()
}
