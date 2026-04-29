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
	Health      float64 // from 0 to 100
	Hapiness    float64 // from 0 to 100
	Inspiration float64 // from 0 to 100
	Popularity  float64 // from 0 to 100

	LyricsSkill       float64 // from 0 to 100
	MusicWritingSkill float64 // from 0 to 100
	PlayingSkill      float64 // from 0 to 100
	SingingSkill      float64 // from 0 to 100
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
		fmt.Sprintf("%s: %.0f", "Написание слов", s.LyricsSkill),
		fmt.Sprintf("%s: %.0f", "Написание музыки", s.MusicWritingSkill),
		fmt.Sprintf("%s: %.0f", "Навык игры на гитаре", s.PlayingSkill),
		fmt.Sprintf("%s: %.0f", "Навык исполнения", s.SingingSkill),
	}
	return strings.Join(skills, "\n")
}

func (s *Stats) RenderHealth() string {
	return fmt.Sprintf("%.0f", s.Health)
}

func (s *Stats) RenderOtherInfo() string {
	res := []string{
		fmt.Sprintf("Счастье: %.0f", s.Hapiness),
		fmt.Sprintf("Вдохновение: %.0f", s.Inspiration),
		fmt.Sprintf("Популярность: %.0f", s.Popularity),
	}
	return strings.Join(res, "\n")
}
