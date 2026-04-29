package game

import "game/tui"

func (g *Game) PlayOutside() {}
func (g *Game) StreetShop()  {}

func (g *Game) OtherMenuActions() {
	g.app.SetActions([]tui.Action{
		{
			Label: "Выступать на улице",
			Handle: func(a *tui.App) {
				g.PlayOutside()
			},
		},
		{
			Label: "Уличный рынок",
			Handle: func(a *tui.App) {
				g.StreetShop()
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
