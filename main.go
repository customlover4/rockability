package main

import (
	"fmt"
	"log"
	"time"

	"game/tui"
)

func main() {
	app := tui.NewApp()

	playerName := "Игрок"
	className := "Странник"
	coins := 25
	day := 1

	var refresh func()
	refresh = func() {
		app.SetScreen(tui.Screen{
			ActionTitle: "Команды",
			LogTitle:    "Вывод",
			StatusTitle: "Параметры",
			Log: []string{
				"Это чистый UI-каркас без игровой сцены.",
				"Меняй actions, log и status так, как удобно твоей игре.",
			},
			Status: []tui.StatusItem{
				{Label: "Имя", Value: playerName},
				{Label: "Класс", Value: className},
				{Label: "Монеты", Value: fmt.Sprintf("%d", coins)},
				{Label: "День", Value: fmt.Sprintf("%d", day)},
			},
			Actions: []tui.Action{
				{
					Label: "Добавить строку в лог",
					Handle: func(app *tui.App) {
						app.AppendLog("Ты добавил новую строку в журнал через callback действия.")
					},
				},
				{
					Label: "Показать сообщение",
					Handle: func(app *tui.App) {
						app.ShowMessage("Сообщение", "Это обычный popup для уведомлений.")
					},
				},
				{
					Label: "Открыть выбор стрелками",
					Handle: func(app *tui.App) {
						app.ShowChoicePopup("Выбор", "Это заготовка выбора. Стрелки двигают курсор, Enter подтверждает.", []tui.ChoiceOption{
							{
								Label: "Получить 5 монет",
								Handle: func(app *tui.App) {
									coins += 5
									app.AppendLog("Ты выбрал вариант с наградой.")
									refresh()
								},
							},
							{
								Label: "Потратить день",
								Handle: func(app *tui.App) {
									day++
									app.AppendLog("Ты выбрал вариант, который продвигает время.")
									refresh()
								},
							},
						})
					},
				},
				{
					Label: "Открыть форму",
					Handle: func(app *tui.App) {
						app.ShowFormPopup("Редактировать параметры", []tui.FormField{
							{Key: "name", Label: "Имя", Value: playerName, DefaultValue: "Игрок"},
							{Key: "class", Label: "Класс", Value: className, DefaultValue: "Странник"},
						}, func(values map[string]string) {
							playerName = values["name"]
							className = values["class"]
							app.AppendLog("Форма сохранена.")
							refresh()
						})
					},
				},
				{
					Label: "Запустить катсцену",
					Handle: func(app *tui.App) {
						app.PlayCutscene(tui.Cutscene{
							Title: "Тестовая катсцена",
							Messages: []string{
								"Это не игровая логика, а просто визуальный слой.",
								"Ты можешь вызывать его из любого своего сценария.",
							},
							Duration: 4 * time.Second,
						})
					},
				},
			},
		})
	}

	refresh()

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
