package game

import (
	"game/game/player"
	"game/tui"
)

type Game struct {
	app *tui.App

	p *player.Player
}

func New(app *tui.App) *Game {
	return &Game{
		app: app,

		p: player.DefaultPlayer(),
	}
}


