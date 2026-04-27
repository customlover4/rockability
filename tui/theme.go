package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func DefaultTheme() Theme {
	return Theme{
		Background:         tcell.NewRGBColor(10, 12, 18),
		PanelBackground:    tcell.NewRGBColor(18, 22, 31),
		PanelText:          tcell.NewRGBColor(228, 229, 231),
		Border:             tcell.NewRGBColor(90, 124, 176),
		Title:              tcell.NewRGBColor(127, 214, 255),
		StatusAccent:       "gold",
		StatusValue:        "white",
		ListSelectedText:   tcell.ColorBlack,
		ListSelectedBg:     tcell.NewRGBColor(110, 224, 190),
		ModalBackground:    tcell.NewRGBColor(23, 27, 38),
		CutsceneBackground: tcell.ColorBlack,
	}
}

func applyTheme(theme Theme) {
	tview.Styles.PrimitiveBackgroundColor = theme.Background
	tview.Styles.ContrastBackgroundColor = theme.PanelBackground
	tview.Styles.PrimaryTextColor = theme.PanelText
	tview.Styles.BorderColor = theme.Border
	tview.Styles.TitleColor = theme.Title
}
