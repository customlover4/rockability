package game

import (
	"game/tui"
)

func (g *Game) SkillInfo() {
	g.app.ShowMessage(
		"Навыки",
		g.p.Stats.RenderStats(),
	)
}

func (g *Game) BandInfo() {
	g.app.ShowMessage(
		"Группа",
		g.p.Band.RenderInfo(),
	)
}

func (g *Game) Inventory() {
	// todo: open other menu
}

func (g *Game) InfoMenuActions() {
	g.app.SetActions([]tui.Action{
		{
			Label: "Навыки",
			Handle: func(a *tui.App) {
				g.SkillInfo()
			},
		},
		{
			Label: "Группа",
			Handle: func(a *tui.App) {
				g.BandInfo()
			},
		},
		{
			Label: "Инвентарь",
			Handle: func(a *tui.App) {
				g.Inventory()
			},
		},
		{
			Label: "Назад",
			Handle: func(a *tui.App) {
				g.BaseActions()
			},
		},
	})
}
