package stats

import (
	"fmt"
	"game/tui"
	"strings"
	"time"
)

const BirthMonth = time.December
const BirthDay = 28

type Stats struct {
	Age         int8
	Health      int8 // from 0 to 100
	Hapiness    int8 // from 0 to 100
	Inspiration int8 // from 0 to 100
	Popularity  int8 // from 0 to 100

	LyricsSkill       int8 // from 0 to 100
	MusicWritingSkill int8 // from 0 to 100
	PlayingSkill      int8 // from 0 to 100
	SingingSkill      int8 // from 0 to 100
}

func DefaultStats() Stats {
	return Stats{
		Age:         18,
		Health:      100,
		Hapiness:    100,
		Inspiration: 50,
		Popularity:  0,

		LyricsSkill:       10,
		MusicWritingSkill: 10,
		PlayingSkill:      10,
		SingingSkill:      10,
	}
}

func (s *Stats) ProcessNewDay(a *tui.App, newDate time.Time) {
	if newDate.Month() == BirthMonth && newDate.Day() == BirthDay {
		a.AppendLog("Сегодня твой день рождения, отпразднуй его!")
		s.Age += 1
	}
}

func (s *Stats) RenderStats() string {
	var skills = []string{
		fmt.Sprintf("%s: %d", "Написание слов", s.LyricsSkill),
		fmt.Sprintf("%s: %d", "Написание музыки", s.MusicWritingSkill),
		fmt.Sprintf("%s: %d", "Навык игры на гитаре", s.PlayingSkill),
		fmt.Sprintf("%s: %d", "Навык исполнения", s.SingingSkill),
	}
	return strings.Join(skills, "\n")
}

func (s *Stats) RenderHealth() string {
	if s.Health >= 90 {
		return "великолепное"
	} else if s.Health >= 60 {
		return "удовлетворительное"
	} else if s.Health >= 30 {
		return "плохое"
	} else if s.Health >= 10 {
		return "ужасное"
	} else {
		return "присмерти"
	}
}
