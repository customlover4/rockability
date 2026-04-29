package game

import "game/tui"

func (g *Game) MyMeetings()   {}
func (g *Game) VisitClub()    {}
func (g *Game) VisitGalery()  {}
func (g *Game) VisitConcert() {}
func (g *Game) WalkAround()   {}

func (g *Game) NetworkingMenuActions() {
	g.app.SetActions([]tui.Action{
		{
			Label: "Мои знакомства",
			Handle: func(a *tui.App) {
				g.VisitClub()
			},
		},
		{
			Label: "Посетить клуб",
			Handle: func(a *tui.App) {
				g.VisitClub()
			},
		},
		{
			Label: "Посетить арт гелерею",
			Handle: func(a *tui.App) {
				g.VisitGalery()
			},
		},
		{
			Label: "Посетить концерт",
			Handle: func(a *tui.App) {
				g.VisitConcert()
			},
		},
		{
			Label: "Выйти на прогулку",
			Handle: func(a *tui.App) {
				g.WalkAround()
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
