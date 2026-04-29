package game

import "game/tui"

func (g *Game) SetListEdit() {}
func (g *Game) NewMember()   {}
func (g *Game) ManagerMenu() {}
func (g *Game) LabelMenu()   {}
func (g *Game) MerchMenu()   {}
func (g *Game) NewGig()      {}
func (g *Game) NewTour()     {}
func (g *Game) Discography() {}

func (g *Game) BandMenuActions() {
	band := []tui.Action{}
	// todo: add peoples to context menu
	// todo: make struct for band

	all := []tui.Action{
		{
			Label: "Сет-лист",
			Handle: func(a *tui.App) {
				g.SetListEdit()
			},
		},
		{
			Label: "Найти участника",
			Handle: func(a *tui.App) {
				g.NewMember()
			},
		},
		{
			Label: "Менеджер",
			Handle: func(a *tui.App) {
				g.ManagerMenu()
			},
		},
		{
			Label: "Лейбл",
			Handle: func(a *tui.App) {
				g.LabelMenu()
			},
		},
		{
			Label: "Мерч",
			Handle: func(a *tui.App) {
				g.MerchMenu()
			},
		},
		{
			Label: "Организовать выступление",
			Handle: func(a *tui.App) {
				g.NewGig()
			},
		},
		{
			Label: "Организовать тур",
			Handle: func(a *tui.App) {
				g.NewTour()
			},
		},
		{
			Label: "Дискография",
			Handle: func(a *tui.App) {
				g.NewTour()
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
