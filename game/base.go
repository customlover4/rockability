package game

import "game/tui"

func (g *Game) BaseUpdateScreen() {
	g.app.SetScreen(tui.Screen{
		ActionTitle: "Действия",
		LogTitle:    "Вывод",
		StatusTitle: "Статус",
		Log: []string{
			"Твой путь начинается, тебе 18, на улице 1995,",
			"твоя цель стать лучшим музыкантом в мире.",
			"Ты создаешь свою группу и пока ты один, но все впереди.",
		},
		Status:  g.p.RenderStatus(),
		Actions: []tui.Action{},
	})
	g.BaseActions()
}

func (g *Game) BaseActions() {
	g.app.SetActions(
		[]tui.Action{
			{
				Label: "Инфо",
				Handle: func(a *tui.App) {
					g.InfoMenuActions()
				},
			},
			{
				Label: "Музыка",
				Handle: func(a *tui.App) {
					g.MusicMenuActions()
				},
			},
			{
				Label: "Группа (in progress)",
				Handle: func(a *tui.App) {
					g.BandMenuActions()
				},
			},
			{
				Label: "Финансы (in progress)",
				Handle: func(a *tui.App) {
					g.FinanceMenuActions()
				},
			},
			{
				Label: "Общение (in progress)",
				Handle: func(a *tui.App) {
					g.NetworkingMenuActions()
				},
			},
			{
				Label: "Разное (in progress)",
				Handle: func(a *tui.App) {
					g.OtherMenuActions()
				},
			},
			{
				Label: "Пропустить день (in progress)",
				Handle: func(a *tui.App) {
					g.p.SkipDay(a)
				},
			},
			{
				Label: "Выйти и сохранить (in progress)",
				Handle: func(a *tui.App) {
					a.Stop()
					// todo: realiz
				},
			},
		},
	)
}
