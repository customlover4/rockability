package game

import "game/tui"

func (g *Game) PlayOutside(a *tui.App) {}

func (g *Game) OtherMenuActions() {
	g.app.SetActions([]tui.Action{
		{
			Label: "Выступать на улице",
			Handle: func(a *tui.App) {
				g.VisitClub(a)
			},
		},
		// ...
		{
			Label: "Назад",
			Handle: func(a *tui.App) {
				g.BaseActions()
			},
		},
	})
}
