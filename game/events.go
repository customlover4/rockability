package game

import (
	"game/game/entities/event"
	"game/game/entities/music"
	"game/tui"
)

func (g *Game) NewIdeaEvent() {
	if g.rnd.Float64() < 0.1 {
		g.p.NewEvent(event.Event{
			Title: "Вдохновение!",
			Handle: func(a *tui.App) {
				a.ShowMessage(
					"Новая идея",
					"Вам пришла новая идея для песни!",
				)
				g.p.Ideas = append(g.p.Ideas, music.NewIdea())
			},
		})
	}
}
