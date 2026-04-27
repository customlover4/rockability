package tui

import (
	"math/rand"
	"sync/atomic"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Theme struct {
	Background         tcell.Color
	PanelBackground    tcell.Color
	PanelText          tcell.Color
	Border             tcell.Color
	Title              tcell.Color
	StatusAccent       string
	StatusValue        string
	ListSelectedText   tcell.Color
	ListSelectedBg     tcell.Color
	ModalBackground    tcell.Color
	CutsceneBackground tcell.Color
}

type StatusItem struct {
	Label string
	Value string
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
}

type App struct {
	app            *tview.Application
	pages          *tview.Pages
	theme          Theme
	rng            *rand.Rand
	currentActions []Action
	cutsceneActive atomic.Bool
	onQuit         func()

	screen Screen

	description *tview.TextView
	status      *tview.TextView
	actionsList *tview.List
}

func newRNG() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}
