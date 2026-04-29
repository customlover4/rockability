package game

import "game/tui"

func (g *Game) Frilance(a *tui.App)    {}
func (g *Game) MyProperty(a *tui.App)  {}
func (g *Game) BuyProperty(a *tui.App) {}
func (g *Game) Supermarket(a *tui.App) {}
func (g *Game) MusicShop(a *tui.App)   {}

func (g *Game) FinanceMenuActions() {
	g.app.SetActions([]tui.Action{
		{
			Label: "Подработка",
			Handle: func(a *tui.App) {
				g.Frilance(a)
			},
		},
		{
			Label: "Ваше имущество",
			Handle: func(a *tui.App) {
				g.MyProperty(a)
			},
		},
		{
			Label: "Покупка имущества",
			Handle: func(a *tui.App) {
				g.BuyProperty(a)
			},
		},
		{
			Label: "Гипермаркет",
			Handle: func(a *tui.App) {
				g.Supermarket(a)
			},
		},
		{
			Label: "Музыкальный магазин",
			Handle: func(a *tui.App) {
				g.MusicShop(a)
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
