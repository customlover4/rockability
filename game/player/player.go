package player

import (
	"fmt"
	"game/game/player/region"
	"game/game/player/stats"
	"game/tui"
	"time"
)

const startDate = "1995-12-28"

type Player struct {
	region region.Region

	playerName string
	groupName  string
	money      int
	agetime    time.Time

	stats stats.Stats
}

func DefaultPlayer() *Player {
	at, _ := time.Parse(time.DateOnly, startDate)

	return &Player{
		playerName: "Безымянный",
		groupName:  "Безымянный",
		region:     region.DefaultRegion(),
		money:      0,
		stats:      stats.DefaultStats(),
		agetime:    at,
	}
}

func (p *Player) RenderStatus() []tui.StatusItem {
	return []tui.StatusItem{
		{Label: "Имя", Value: p.playerName},
		{Label: "Возраст", Value: fmt.Sprintf("%d", p.stats.Age)},
		{Label: "Группа", Value: p.groupName},
		{Label: "Регион", Value: p.region.Title},
		{Label: "Деньги", Value: fmt.Sprintf(
			"%d %s", p.money, p.region.MoneyFormat,
		)},
		{Label: "Дата", Value: p.agetime.Format(time.DateOnly)},
	}
}

// todo: save/load functions
