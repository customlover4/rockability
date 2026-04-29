package band

import (
	"fmt"
	"game/game/entities/label"
	"game/game/entities/manager"
	"game/game/entities/music"
	"game/game/entities/person"
	"game/tui"
	"strings"
	"time"
)

type Band struct {
	GroupName string

	Drummer    person.Person
	BassGuitar person.Person

	Teamwork   float64
	Popularity int8 // from 0 to 100
	Fans       int

	Manager *manager.Manager // todo
	Label   *label.Label     // todo

	SetList []music.Song

	Songs  []music.Song
	Albums []music.Album
}

func NewBand(name string) *Band {
	return &Band{
		GroupName:  name,
		Teamwork:   0,
		Popularity: 0,
		Fans:       0,

		Manager: nil,
		Label:   nil,

		SetList: make([]music.Song, 0),
		Songs:   make([]music.Song, 0),
		Albums:  make([]music.Album, 0),
	}
}

func clamp(value float64) float64 {
	if value <= 0 {
		return 0
	} else if value >= 100 {
		return 100
	}

	return value
}

func (b *Band) AddTeamwork(value float64) {
	b.Teamwork = clamp(b.Teamwork + value)
}

func (b *Band) AddPlayskill(value float64) {
	b.BassGuitar.PlayingSkill = clamp(b.BassGuitar.PlayingSkill + value)
	b.Drummer.PlayingSkill = clamp(b.Drummer.PlayingSkill + value)
}

func (b *Band) RenderInfo() string {
	var members []string = []string{
		b.Drummer.Name, b.BassGuitar.Name,
	}
	var info []string = []string{
		fmt.Sprintf("Название: %s", b.GroupName),
		fmt.Sprintf("Участники: %s", strings.Join(members, ", ")),
		fmt.Sprintf("Сыгранность: %.0f", b.Teamwork),
		fmt.Sprintf("Популярность: %d", b.Popularity),
		fmt.Sprintf("Фанаты: %d", b.Fans),
	}

	return strings.Join(info, "\n")
}

func (b *Band) ProcessNewDay(a *tui.App, newDate time.Time) {}
