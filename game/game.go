package game

import (
	"game/game/player"
	"game/tui"
	"math/rand/v2"
	"time"
)

type Game struct {
	app *tui.App
	rnd *rand.Rand
	p   *player.Player
}

func New(app *tui.App) *Game {
	return &Game{
		app: app,
		rnd: rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), uint64(time.Now().Unix()))),

		p: player.DefaultPlayer(),
	}
}
