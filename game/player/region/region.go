package region

import (
	"game/tui"
	"time"
)

type Region struct {
	Title       string
	MoneyFormat string

	NewDayProcessor func(a *tui.App, newDate time.Time)
	// todo: make population
	// todo: make other x, like 0.25x by population for max cap on stadium
	// todo: or listeners of new album
}

func DefaultRegion() Region {
	return Region{
		Title:       "Не выбрано",
		MoneyFormat: "?",
		NewDayProcessor: func(a *tui.App, newDate time.Time) {
			// empty
		},
	}
}

// todo: make func
var ChoiceRegion = map[int]Region{
	1: {
		Title:       "USA",
		MoneyFormat: "$",
		NewDayProcessor: func(a *tui.App, newDate time.Time) {
			// empty
		},
	},
}
