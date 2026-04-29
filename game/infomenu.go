package game

import (
	"game/tui"
)

func (g *Game) GlobalInfo() {
	g.app.ShowMessage(
		"Основная информация",
		g.p.Stats.RenderOtherInfo(),
	)
}

func (g *Game) HealthInfo() {
	g.app.ShowMessage(
		"Здоровье",
		g.p.Stats.RenderHealth(),
	)
}

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
			Label: "Общая информаци",
			Handle: func(a *tui.App) {
				g.GlobalInfo()
			},
		},
		{
			Label: "Здоровье",
			Handle: func(a *tui.App) {
				g.HealthInfo()
			},
		},
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
