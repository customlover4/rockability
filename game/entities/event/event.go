package event

import (
	"game/tui"
)

type Event struct {
	Title  string
	Handle func(a *tui.App)
}
