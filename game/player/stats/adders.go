package stats

func clampStatValue(value float64) float64 {
	if value < 0 {
		return 0
	}
	if value > 100 {
		return 100
	}
	return value
}

func (s *Stats) AddHealth(value float64) {
	s.Health = clampStatValue(s.Health + value)
}

func (s *Stats) AddHapiness(value float64) {
	s.Hapiness = clampStatValue(s.Hapiness + value)
}

func (s *Stats) AddInspiration(value float64) {
	s.Inspiration = clampStatValue(s.Inspiration + value)
}

func (s *Stats) AddPopularity(value float64) {
	s.Popularity = clampStatValue(s.Popularity + value)
}

func (s *Stats) AddLyricsSkill(value float64) {
	s.LyricsSkill = clampStatValue(s.LyricsSkill + value)
}

func (s *Stats) AddMusicWritingSkill(value float64) {
	s.MusicWritingSkill = clampStatValue(s.MusicWritingSkill + value)
}

func (s *Stats) AddPlayingSkill(value float64) {
	s.PlayingSkill = clampStatValue(s.PlayingSkill + value)
}

func (s *Stats) AddSingingSkill(value float64) {
	s.SingingSkill = clampStatValue(s.SingingSkill + value)
}
