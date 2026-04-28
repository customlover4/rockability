package player

import (
	"game/tui"
	"time"
)

func (p *Player) ProcessStructs(a *tui.App) {
	p.stats.ProcessNewDay(a, p.agetime)
	p.region.NewDayProcessor(a, p.agetime)
	// todo: here new days generators
	// todo: for region and other structs data
}

func (p *Player) SkipDay(a *tui.App) {
	p.agetime = p.agetime.Add(time.Hour * 24)

	p.ProcessStructs(a)

	// todo: generate random events

	a.AppendLog("Настал новый день.")
	a.SetStatus(p.RenderStatus())
}
