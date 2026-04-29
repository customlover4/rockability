package player

import (
	"container/list"
	"fmt"
	"game/game/entities/band"
	"game/game/entities/event"
	"game/game/entities/music"
	"game/game/entities/region"
	"game/game/entities/stats"
	"game/tui"
	"time"
)

const startDate = "1995-12-28"

type Player struct {
	Region region.Region

	playerName string
	Money      int
	Agetime    time.Time

	Stats           stats.Stats
	Band            *band.Band
	ImprovizationXP float64
	Ideas           []music.Idea

	Events *list.List
}

func DefaultPlayer() *Player {
	at, _ := time.Parse(time.DateOnly, startDate)

	return &Player{
		playerName: "Безымянный",
		Region:     region.DefaultRegion(),
		Money:      0,
		Agetime:    at,

		Stats:           stats.DefaultStats(),
		ImprovizationXP: 0,
		Ideas:           make([]music.Idea, 0),

		Events: list.New(),
	}
}

func (p *Player) RenderStatus() []tui.StatusItem {
	return []tui.StatusItem{
		{Label: "Имя", Value: p.playerName},
		{Label: "Возраст", Value: fmt.Sprintf("%d", p.Stats.Age)},
		{Label: "Регион", Value: p.Region.Title},
		{Label: "Деньги", Value: fmt.Sprintf(
			"%d $", p.Money,
		)},
		{Label: "Дата", Value: p.Agetime.Format(time.DateOnly)},
	}
}

// todo: save/load functions

func (p *Player) SetName(name string) {
	p.playerName = name
}

func (p *Player) SetRegion(region region.Region) {
	p.Region = region
}

func (p *Player) AddMoney(value int) bool {
	if p.Money+value < 0 {
		return false
	}

	p.Money += value
	return true
}

func (p *Player) NewEvent(e event.Event) {
	p.Events.PushBack(e)
}
