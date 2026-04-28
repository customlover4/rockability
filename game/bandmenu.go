package game

import "game/tui"

func (g *Game) SetListEdit(a *tui.App) {}
func (g *Game) NewMember(a *tui.App)   {}
func (g *Game) ManagerMenu(a *tui.App) {}
func (g *Game) LabelMenu(a *tui.App)   {}
func (g *Game) MerchMenu(a *tui.App)   {}
func (g *Game) NewGig(a *tui.App)      {}
func (g *Game) NewTour(a *tui.App)     {}
func (g *Game) Discography(a *tui.App) {}

func (g *Game) BandMenuActions() {
	band := []tui.Action{}
	// todo: add peoples to context menu
	// todo: make struct for band

	all := []tui.Action{
		{
			Label: "Сет-лист",
			Handle: func(a *tui.App) {
				g.SetListEdit(a)
			},
		},
		{
			Label: "Найти участника",
			Handle: func(a *tui.App) {
				g.NewMember(a)
			},
		},
		{
			Label: "Менеджер",
			Handle: func(a *tui.App) {
				g.ManagerMenu(a)
			},
		},
		{
			Label: "Лейбл",
			Handle: func(a *tui.App) {
				g.LabelMenu(a)
			},
		},
		{
			Label: "Мерч",
			Handle: func(a *tui.App) {
				g.MerchMenu(a)
			},
		},
		{
			Label: "Организовать выступление",
			Handle: func(a *tui.App) {
				g.NewGig(a)
			},
		},
		{
			Label: "Организовать тур",
			Handle: func(a *tui.App) {
				g.NewTour(a)
			},
		},
		{
			Label: "Дискография",
			Handle: func(a *tui.App) {
				g.NewTour(a)
			},
		},
		{
			Label: "Назад",
			Handle: func(a *tui.App) {
				g.BaseActions()
			},
		},
	}

	band = append(band, all...)

	g.app.SetActions(band)
}
