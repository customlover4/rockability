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
	Region region.Region

	playerName string
	groupName  string
	money      int
	agetime    time.Time

	Stats stats.Stats
}

func DefaultPlayer() *Player {
	at, _ := time.Parse(time.DateOnly, startDate)

	return &Player{
		playerName: "Безымянный",
		groupName:  "Безымянный",
		Region:     region.DefaultRegion(),
		money:      0,
		Stats:      stats.DefaultStats(),
		agetime:    at,
	}
}

func (p *Player) RenderStatus() []tui.StatusItem {
	return []tui.StatusItem{
		{Label: "Имя", Value: p.playerName},
		{Label: "Возраст", Value: fmt.Sprintf("%d", p.Stats.Age)},
		{Label: "Группа", Value: p.groupName},
		{Label: "Регион", Value: p.Region.Title},
		{Label: "Деньги", Value: fmt.Sprintf(
			"%d %s", p.money, p.Region.MoneyFormat,
		)},
		{Label: "Дата", Value: p.agetime.Format(time.DateOnly)},
	}
}

// todo: save/load functions
