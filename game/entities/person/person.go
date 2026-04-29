package person

import (
	"game/game/entities/region"
	"time"
)

const (
	MaleGender = iota
	FemaleGender
)

// non-playable persons in the world
type Person struct {
	Name     string
	Age      int
	Gender   int8
	Region   region.Region
	BirthDay time.Time

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

func NewMale() {}
