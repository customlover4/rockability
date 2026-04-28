package game

import "game/tui"

func (g *Game) Improvization(a *tui.App) {

}

func (g *Game) SoloRepetition(a *tui.App) {

}

func (g *Game) BandRepetition(a *tui.App) {

}

func (g *Game) WriteSongs(a *tui.App) {

}

func (g *Game) RewriteSong(a *tui.App) {

}

func (g *Game) ShowSongs(a *tui.App) {

}

func (g *Game) RecordSongs(a *tui.App) {
	// some buffs and debuffs
}

func (g *Game) MusicMenuActions() {
	g.app.SetActions([]tui.Action{
		{
			Label: "Импровизировать на гитаре",
			Handle: func(a *tui.App) {
				g.Improvization(a)
			},
		},
		{
			Label: "Доработать идеи",
			Handle: func(a *tui.App) {
				g.RewriteSong(a)
			},
		},
		{
			Label: "Написать трек",
			Handle: func(a *tui.App) {
				g.WriteSongs(a)
			},
		},
		{
			Label: "Репетировать одному",
			Handle: func(a *tui.App) {
				g.SoloRepetition(a)
			},
		},
		{
			Label: "Репетировать с группой",
			Handle: func(a *tui.App) {
				g.BandRepetition(a)
			},
		},
		{
			Label: "Записать трек",
			Handle: func(a *tui.App) {
				g.RecordSongs(a)
			},
		},
		{
			Label: "Каталог песен",
			Handle: func(a *tui.App) {
				g.ShowSongs(a)
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
