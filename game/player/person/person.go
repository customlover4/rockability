package person

// non-playable persons in the world
type Person struct {
	Name         string
	PlayingSkill float64
}

func clampValue(value float64) float64 {
	if value < 0 {
		return 0
	} else if value > 100 {
		return 100
	}

	return value
}

func (p *Person) AddPlayingSkill(value float64) {
	p.PlayingSkill = clampValue(p.PlayingSkill + value)
}
