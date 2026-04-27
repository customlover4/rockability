package main

import (
	"log"

	"game/game"
	"game/tui"
)

func main() {
	app := tui.NewApp()

	g := game.New(app)
	g.Render()

	

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
