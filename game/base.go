package game

import "game/tui"

// todo: think about base actions
func (g *Game) BaseUpdateScreen() {
	g.app.SetScreen(tui.Screen{
		ActionTitle: "Действия",
		LogTitle:    "Вывод",
		StatusTitle: "Статус",
		Log: []string{
			"Твой путь начинается, тебе 18, только окончил школу,",
			"твоя цель стать лучшим музыкантом в мире.",
		},
		Status: g.p.RenderStatus(),
		Actions: []tui.Action{
			{
				Label: "Музыка",
				Handle: func(a *tui.App) {
					// todo: realiz
				},
			},
			{
				Label: "Группа",
				Handle: func(a *tui.App) {
					// todo: realiz
				},
			},
			{
				Label: "Финансы",
				Handle: func(a *tui.App) {
					// todo: realiz
				},
			},
			{
				Label: "Общение",
				Handle: func(a *tui.App) {
					// todo: realiz
				},
			},
			{
				Label: "Разное",
				Handle: func(a *tui.App) {
					// todo: realiz
				},
			},
			{
				Label: "Пропустить день",
				Handle: func(a *tui.App) {
					// todo: realiz
				},
			},
			{
				Label: "Выйти и сохранить",
				Handle: func(a *tui.App) {
					a.Stop()
					// todo: realiz
				},
			},
		},
	})
}
