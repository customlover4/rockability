package band

import (
	"fmt"
	"game/game/player/music"
	"game/game/player/music/label"
	"game/game/player/music/manager"
	"game/game/player/person"
	"game/tui"
	"strings"
	"time"
)

type Band struct {
	GroupName string

	Members []person.Person

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
		Members:    make([]person.Person, 0),
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

func (b *Band) AddTeamwork(value float64) {
	if b.Teamwork + value < 0 {
		b.Teamwork = 0
	} else if b.Teamwork + value > 100 {
		b.Teamwork = 100
	}

	b.Teamwork += value
}

func (b *Band) RenderInfo() string {
	var members []string
	for _, v := range b.Members {
		members = append(members, v.Name)
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
