package game

import "game/tui"

func (g *Game) Frilance()    {}
func (g *Game) MyProperty()  {}
func (g *Game) BuyProperty() {}
func (g *Game) Supermarket() {}
func (g *Game) MusicShop()   {}

func (g *Game) FinanceMenuActions() {
	g.app.SetActions([]tui.Action{
		{
			Label: "Подработка",
			Handle: func(a *tui.App) {
				g.Frilance()
			},
		},
		{
			Label: "Ваше имущество",
			Handle: func(a *tui.App) {
				g.MyProperty()
			},
		},
		{
			Label: "Покупка имущества",
			Handle: func(a *tui.App) {
				g.BuyProperty()
			},
		},
		{
			Label: "Гипермаркет",
			Handle: func(a *tui.App) {
				g.Supermarket()
			},
		},
		{
			Label: "Музыкальный магазин",
			Handle: func(a *tui.App) {
				g.MusicShop()
			},
		},
		{
			Label: "Назад",
			Handle: func(a *tui.App) {
				g.BaseActions()
			},
		},
	})
}
