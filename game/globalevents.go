package game

import (
	"fmt"
	"game/game/entities/event"
	"game/tui"
	"strings"
)

func (g *Game) RequiredGlobalEvents() {
	g.EmergencyHealth()
}

func (g *Game) EmergencyHealth() {
	if g.p.Stats.Health > 0 {
		return
	}
	g.p.NewEvent(event.Event{
		Title: "Здоровье на нуле",
		Handle: func(a *tui.App) {
			strs := []string{
				"Твое здоровье слишком низкое, тебя срочно госпитализировали!",
				"Пока здоровье не поднимется ты останешься в больнице!",
				"Есть шанс что ты умрешь, надейся на врачей",
			}
			text := strings.Join(strs, "\n")
			a.ShowMessage(
				"Ты попал в больницу!",
				text,
			)

			errors := 0
			coef := 0.6
			attemps := 0
			for errors < 5 && g.p.Stats.Health < 20 {
				if g.rnd.Float64() < coef {
					a.ShowMessage(
						"Больница", " У врачей получается тебя стабилизировать.",
					)
					g.p.Stats.AddHealth(5)
				} else {
					a.ShowMessage(
						"Больница", " Твое состояние ухудшается.",
					)
					errors++
					coef -= 0.1
				}
				g.SkipDaysWithoutEvents(1)
				attemps++
			}

			a.ShowMessage(
				"Больница",
				fmt.Sprintf(
					"Врачи боролись за вашу жизнь %d подряд",
					attemps,
				),
			)
			if errors >= 5 {
				a.ShowMessage(
					"Вы погибли.",
					"Ваша история заканчивается на больничной койке.",
				)
				a.SetActions([]tui.Action{}) // todo: make end actions, like upload history
			} else {
				a.ShowMessage(
					"Победа!", "Вы выжили! В следующий раз будьте осторожны.",
				)
			}
		},
	})
}
