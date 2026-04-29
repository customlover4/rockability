package game

import (
	"game/tui"
)

func (g *Game) GlobalInfo(a *tui.App) {
	a.ShowMessage(
		"Основная информация",
		g.p.Stats.RenderOtherInfo(),
	)
}

func (g *Game) HealthInfo(a *tui.App) {
	a.ShowMessage(
		"Здоровье",
		g.p.Stats.RenderHealth(),
	)
}

func (g *Game) SkillInfo(a *tui.App) {
	a.ShowMessage(
		"Навыки",
		g.p.Stats.RenderStats(),
	)
}

func (g *Game) BandInfo(a *tui.App) {
	a.ShowMessage(
		"Группа",
		g.p.Band.RenderInfo(),
	)
}

func (g *Game) Inventory(a *tui.App) {
	// todo: open other menu
}

func (g *Game) InfoMenuActions() {
	g.app.SetActions([]tui.Action{
		{
			Label: "Общая информаци",
			Handle: func(a *tui.App) {
				g.GlobalInfo(a)
			},
		},
		{
			Label: "Здоровье",
			Handle: func(a *tui.App) {
				g.HealthInfo(a)
			},
		},
		{
			Label: "Навыки",
			Handle: func(a *tui.App) {
				g.SkillInfo(a)
			},
		},
		{
			Label: "Группа",
			Handle: func(a *tui.App) {
				g.BandInfo(a)
			},
		},
		{
			Label: "Инвентарь",
			Handle: func(a *tui.App) {
				g.Inventory(a)
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
