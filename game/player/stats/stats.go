package stats

type Stats struct {
	Age          int8
	Health       int8 // from 0 to 100
	MentalHealth int8 // from 0 to 100

	LyricsSkill       int8 // from 0 to 100
	MusicWritingSkill int8 // from 0 to 100
	PlayingSkill      int8 // from 0 to 100
	SingingSkill      int8 // from 0 to 100
}

func DefaultStats() Stats {
	return Stats{
		Age:          18,
		Health:       100,
		MentalHealth: 100,

		LyricsSkill:       10,
		MusicWritingSkill: 10,
		PlayingSkill:      10,
		SingingSkill:      10,
	}
}
