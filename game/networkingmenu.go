package game

import "game/tui"

func (g *Game) MyMeetings(a *tui.App)   {}
func (g *Game) VisitClub(a *tui.App)    {}
func (g *Game) VisitGalery(a *tui.App)  {}
func (g *Game) VisitConcert(a *tui.App) {}
func (g *Game) WalkAround(a *tui.App)   {}

func (g *Game) NetworkingMenuActions() {
	g.app.SetActions([]tui.Action{
		{
			Label: "Мои знакомства",
			Handle: func(a *tui.App) {
				g.VisitClub(a)
			},
		},
		{
			Label: "Посетить клуб",
			Handle: func(a *tui.App) {
				g.VisitClub(a)
			},
		},
		{
			Label: "Посетить арт гелерею",
			Handle: func(a *tui.App) {
				g.VisitGalery(a)
			},
		},
		{
			Label: "Посетить концерт",
			Handle: func(a *tui.App) {
				g.VisitConcert(a)
			},
		},
		{
			Label: "Выйти на прогулку",
			Handle: func(a *tui.App) {
				g.WalkAround(a)
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
