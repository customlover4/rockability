// tui/theme.go
package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Theme struct {
	Background         tcell.Color
	PanelBackground    tcell.Color
	PanelText          tcell.Color
	Border             tcell.Color
	Title              tcell.Color
	StatusAccent       string
	StatusValue        string
	ListSelectedText   tcell.Color
	ListSelectedBg     tcell.Color
	ModalBackground    tcell.Color
	CutsceneBackground tcell.Color

	StatHappyColor   string
}

func DefaultTheme() Theme {
	return Theme{
		Background:         tcell.NewRGBColor(18, 18, 22),
		PanelBackground:    tcell.NewRGBColor(32, 32, 40),
		ModalBackground:    tcell.NewRGBColor(28, 24, 32),
		CutsceneBackground: tcell.ColorBlack,

		PanelText: tcell.NewRGBColor(220, 220, 225),

		Border: tcell.NewRGBColor(140, 40, 50),
		Title:  tcell.NewRGBColor(255, 80, 90),

		StatusAccent: "gold",
		StatusValue:  "white",

		ListSelectedText: tcell.ColorBlack,
		ListSelectedBg:   tcell.NewRGBColor(220, 60, 70),

		StatHappyColor:   "orange",
	}
}

func applyTheme(theme Theme) {
	tview.Styles.PrimitiveBackgroundColor = theme.Background
	tview.Styles.ContrastBackgroundColor = theme.PanelBackground
	tview.Styles.PrimaryTextColor = theme.PanelText
	tview.Styles.BorderColor = theme.Border
	tview.Styles.TitleColor = theme.Title
}
