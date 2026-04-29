package player

import (
	"game/tui"
	"time"
)

func (p *Player) ProcessStructs(a *tui.App) {
	p.Stats.ProcessNewDay(a, p.agetime)
	p.Region.ProcessNewDay(a, p.agetime)
	p.Band.ProcessNewDay(a, p.agetime)
}

func (p *Player) GlobalRandomEvents(a *tui.App) {}

func (p *Player) Events(a *tui.App) {

	for p.events.Len() != 0 {
		e := p.events.Front()
		ev := e.Value.(Event)
		ev.Handle(a)
		p.events.Remove(e)
	}
}

func (p *Player) SkipDay(a *tui.App) {
	p.agetime = p.agetime.Add(time.Hour * 24)

	p.ProcessStructs(a)
	p.GlobalRandomEvents(a)

	p.Events(a)
	a.SetStatus(p.RenderStatus())
}
