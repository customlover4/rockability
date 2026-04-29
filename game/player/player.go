package player

import (
	"container/list"
	"fmt"
	"game/game/player/band"
	"game/game/player/music"
	"game/game/player/region"
	"game/game/player/stats"
	"game/tui"
	"time"
)

const startDate = "1995-12-28"

type Player struct {
	Region region.Region

	playerName string
	money      int
	agetime    time.Time

	Stats           stats.Stats
	Band            *band.Band
	ImprovizationXP float64
	Ideas           []music.Idea

	events *list.List
}

func DefaultPlayer() *Player {
	at, _ := time.Parse(time.DateOnly, startDate)

	return &Player{
		playerName: "Безымянный",
		Region:     region.DefaultRegion(),
		money:      0,
		agetime:    at,

		Stats:           stats.DefaultStats(),
		ImprovizationXP: 0,
		Ideas:           make([]music.Idea, 0),

		events: list.New(),
	}
}

func (p *Player) RenderStatus() []tui.StatusItem {
	return []tui.StatusItem{
		{Label: "Имя", Value: p.playerName},
		{Label: "Возраст", Value: fmt.Sprintf("%d", p.Stats.Age)},
		{Label: "Регион", Value: p.Region.Title},
		{Label: "Деньги", Value: fmt.Sprintf(
			"%d $", p.money,
		)},
		{Label: "Дата", Value: p.agetime.Format(time.DateOnly)},
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
	if p.money+value < 0 {
		return false
	}

	p.money += value
	return true
}

func (p *Player) NewEvent(e Event) {
	p.events.PushBack(e)
}
