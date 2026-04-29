package tui

import (
	"math/rand"
	"time"
)

type StatusItem struct {
	Label string
	Value string
}

type Stat struct {
	Name  string
	Value int
}

type Action struct {
	Label  string
	Handle func(*App)
}

type ChoiceOption struct {
	Label  string
	Handle func(*App)
}

type FormField struct {
	Key          string
	Label        string
	Value        string
	Placeholder  string
	DefaultValue string
}

type Cutscene struct {
	Title    string
	Messages []string
	Duration time.Duration
}

type Screen struct {
	ActionTitle string
	LogTitle    string
	StatusTitle string
	Actions     []Action
	Log         []string
	Status      []StatusItem
	Stats       []Stat
}

func newRNG() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}
