package tui

import (
	"fmt"
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (a *App) PlayCutscene(cutscene Cutscene) {
	// TODO: try it
	if !a.cutsceneActive.CompareAndSwap(false, true) {
		return
	}

	view := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignLeft).
		SetWrap(false)
	view.SetBackgroundColor(a.theme.CutsceneBackground)
	view.SetTextColor(a.theme.PanelText)
	view.SetBorder(false)

	a.pages.AddPage("cutscene", view, true, true)
	a.app.SetFocus(view)

	stop := make(chan struct{})
	closeCutscene := func() {
		select {
		case <-stop:
			return
		default:
			close(stop)
		}
		a.cutsceneActive.Store(false)
		a.pages.RemovePage("cutscene")
		a.app.SetFocus(a.actionsList)
	}

	view.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEscape, tcell.KeyEnter:
			closeCutscene()
			return nil
		}
		return event
	})

	go func() {
		ticker := time.NewTicker(90 * time.Millisecond)
		defer ticker.Stop()

		timeout := time.NewTimer(cutscene.Duration)
		defer timeout.Stop()

		tick := 0
		for {
			select {
			case <-stop:
				return
			case <-timeout.C:
				a.app.QueueUpdateDraw(func() {
					closeCutscene()
				})
				return
			case <-ticker.C:
				a.app.QueueUpdateDraw(func() {
					if a.cutsceneActive.Load() {
						_, _, width, height := view.GetInnerRect()
						frame := a.buildCutsceneFrame(cutscene.Title, cutscene.Messages, tick, width, height)
						view.SetBackgroundColor(tcell.NewRGBColor(
							int32((tick*37)%255),
							int32((tick*67+80)%255),
							int32((tick*97+160)%255),
						))
						view.SetText(frame)
						tick++
					}
				})
			}
		}
	}()
}

func (a *App) buildCutsceneFrame(title string, messages []string, tick, width, height int) string {
	palette := []string{"red", "orange", "yellow", "green", "aqua", "blue", "purple", "fuchsia", "white"}
	backgrounds := []string{"maroon", "olive", "navy", "purple", "teal", "black", "red", "blue"}
	noiseRunes := []rune("▓▒░█@#$%&*+=~<>?/\\|")

	centerLine := func(text string, width int) string {
		runes := []rune(text)
		if len(runes) >= width {
			return string(runes[:width])
		}
		left := (width - len(runes)) / 2
		right := width - len(runes) - left
		return strings.Repeat(" ", left) + text + strings.Repeat(" ", right)
	}

	if width < 12 {
		width = 12
	}
	if height < 6 {
		height = 6
	}

	titleRow := height / 5
	messageRow := height / 2
	hintRow := height - 2

	var b strings.Builder
	for i := 0; i < height; i++ {
		switch i {
		case titleRow:
			line := centerLine(title, width)
			color := palette[tick%len(palette)]
			bg := backgrounds[(tick+i)%len(backgrounds)]
			b.WriteString(fmt.Sprintf("[%s:%s:b]%s[-:-:-]", color, bg, line))
		case messageRow:
			msg := ""
			if len(messages) > 0 {
				msg = messages[tick%len(messages)]
			}
			line := centerLine(msg, width)
			bg := backgrounds[(tick+i+2)%len(backgrounds)]
			b.WriteString(fmt.Sprintf("[white:%s:b]%s[-:-:-]", bg, line))
		case hintRow:
			line := centerLine("Enter или Esc чтобы пропустить", width)
			bg := backgrounds[(tick+i+4)%len(backgrounds)]
			b.WriteString(fmt.Sprintf("[gray:%s]%s[-:-:-]", bg, line))
		default:
			for j := 0; j < width; j++ {
				color := palette[(tick+i+j)%len(palette)]
				bg := backgrounds[(tick*2+i/2+j/3)%len(backgrounds)]
				ch := noiseRunes[a.rng.Intn(len(noiseRunes))]
				if len(messages) > 0 && (i+j+tick)%17 == 0 {
					msgRunes := []rune(messages[(i+tick)%len(messages)])
					if len(msgRunes) > 0 {
						ch = msgRunes[(j+tick)%len(msgRunes)]
					}
				}
				b.WriteString(fmt.Sprintf("[%s:%s]%c", color, bg, ch))
			}
			b.WriteString("[-:-:-]")
		}
		if i < height-1 {
			b.WriteByte('\n')
		}
	}
	return b.String()
}
