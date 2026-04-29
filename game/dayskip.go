package game

import (
	"game/game/player"
	"time"
)

func (g *Game) ProcessStructs() {
	g.p.Stats.ProcessNewDay(g.app, g.p.Agetime)
	g.p.Region.ProcessNewDay(g.app, g.p.Agetime)
	g.p.Band.ProcessNewDay(g.app, g.p.Agetime)
}

func (p *Game) GlobalRandomEvents() {}

func (g *Game) Events() {
	for g.p.Events.Len() != 0 {
		e := g.p.Events.Front()
		ev := e.Value.(player.Event)
		ev.Handle(g.app)
		g.p.Events.Remove(e)
	}
}

func (g *Game) SkipDay(days int) {
	g.p.Agetime = g.p.Agetime.Add(time.Hour * 24)

	g.ProcessStructs()
	g.GlobalRandomEvents()

	g.Events()
	g.app.SetStatus(g.p.RenderStatus())
}
