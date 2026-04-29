package music

import (
	"fmt"
	"game/game/entities/stats"
	"strings"
)

// Idea is a raw music fragment that may later become a full song.
type Idea struct {
	Potential float64
}

type Song struct {
	Title string

	Buzz       float64 // random from potential, hidden value
	Popularity float64
	Quality    float64 // from skills (lyrics + music / 2)

	Skill float64

	Recorded      bool
	RecordQuality float64
	Released      bool
}

type Album struct {
	Title      string
	Songs      []Song
	Popularity int
	Released   bool
}

func NewIdea() Idea {
	return Idea{
		Potential: 15,
	}
}

func NewSong(title string, st stats.Stats, rnd float64, potential float64) Song {
	buzz := (rnd * potential) * 1.2
	if buzz > 100 {
		buzz = 100
	}
	return Song{
		Title:         title,
		Buzz:          buzz,
		Quality:       (st.MusicWritingSkill + st.LyricsSkill) / 2,
		Skill:         0,
		Popularity:    0,
		Recorded:      false,
		RecordQuality: 0.0,
		Released:      false,
	}
}

// func NewAlbum(title string, songs []Song) Album {
// 	return Album{
// 		Title:      title,
// 		Songs:      append([]Song(nil), songs...),
// 		Popularity: 0,
// 		Released:   false,
// 	}
// }

func (i Idea) Render() string {
	return fmt.Sprintf(
		"Идея | потенциал: %.1f",
		i.Potential,
	)
}

func clampValue(value float64) float64 {
	if value < 0 {
		return 0
	}
	if value > 100 {
		return 100
	}
	return value
}

func (i *Idea) AddPotential(value float64) {
	i.Potential = clampValue(i.Potential + value)
}

func (s *Song) SetRecordQuality(value float64) {
	s.RecordQuality = clampValue(value + s.RecordQuality)
}

func (s *Song) AddSkill(value float64) {
	s.Skill = clampValue(s.Skill + value)
}

func (s *Song) Record() {
	s.Recorded = true
}

func (s Song) Render() string {
	title := s.Title
	if strings.TrimSpace(title) == "" {
		title = "Без названия"
	}
	return fmt.Sprintf(
		"%s | Q: %.0f | S: %.0f | P: %.0f",
		title,
		s.Quality,
		s.Skill,
		s.Popularity,
	)
}

func (s Song) InfoRender() string {
	rec := ""
	if s.Recorded {
		rec = fmt.Sprintf(
			"Записана | Качество записи: %.0f", s.RecordQuality,
		)
	} else {
		rec = "Не записана"
	}
	rel := ""
	if s.Released {
		rel = "Выпущена"
	} else {
		rel = "Не выпущена"
	}
	var res []string = []string{
		fmt.Sprintf(
			"Популярность: %.0f | Качество: %.0f | Навык игры: %.0f",
			s.Popularity, s.Quality, s.Skill,
		),
		rec, rel,
	}

	return strings.Join(res, "\n")
}

// func (s *Song) Release() {
// 	if s.State != SongStateRecorded {
// 		return
// 	}
// 	s.State = SongStateReleased
// 	if s.Buzz < 100 {
// 		s.Buzz += 10
// 	}
// 	if s.Popularity < 100 {
// 		s.Popularity += 5
// 	}
// }

// func (a Album) Render() string {
// 	title := a.Title
// 	if strings.TrimSpace(title) == "" {
// 		title = "Без названия"
// 	}
// 	released := "нет"
// 	if a.Released {
// 		released = "да"
// 	}
// 	return fmt.Sprintf(
// 		"%s | треков: %d | популярность: %d | выпущен: %s",
// 		title,
// 		len(a.Songs),
// 		a.Popularity,
// 		released,
// 	)
// }

// func (a *Album) Release() {
// 	a.Released = true
// 	if a.Popularity < 100 {
// 		a.Popularity += 10
// 	}
// }

// func RenderSongsList(songs []Song) string {
// 	if len(songs) == 0 {
// 		return "Каталог песен пуст."
// 	}

// 	lines := make([]string, 0, len(songs))
// 	for i, song := range songs {
// 		lines = append(lines, fmt.Sprintf("%d. %s", i+1, song.Render()))
// 	}
// 	return strings.Join(lines, "\n")
// }

// func RenderAlbumsList(albums []Album) string {
// 	if len(albums) == 0 {
// 		return "Каталог альбомов пуст."
// 	}

// 	lines := make([]string, 0, len(albums))
// 	for i, album := range albums {
// 		lines = append(lines, fmt.Sprintf("%d. %s", i+1, album.Render()))
// 	}
// 	return strings.Join(lines, "\n")
// }

// func (s *Song) promoteStateIfReady() {
// 	if s.ReadyToRecord() && s.State == SongStateDraft {
// 		s.State = SongStateReady
// 	}
// }
