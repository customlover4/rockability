package game

import (
	"game/game/player/region"
	"game/tui"
)

func (g *Game) StartUpUpdateScreen() {
	g.app.SetScreen(tui.Screen{
		ActionTitle: "Команды",
		LogTitle:    "Вывод",
		StatusTitle: "Параметры",
		Log: []string{
			"Добро пожаловать в rockability!",
			"v0.0.1 by customlover4",
		},
		Status: g.p.RenderStatus(),
		Actions: []tui.Action{
			{
				Label:  "Новая игра",
				Handle: g.NewGameStartup,
			},
			{
				Label:  "Загрузить игру",
				Handle: g.LoadGameStartup,
			},
		},
	})
}

func (g *Game) ChoiceRegion(a *tui.App) {
	regions := make([]tui.ChoiceOption, 0, len(region.ChoiceRegion))
	for _, v := range region.ChoiceRegion {
		regions = append(regions, tui.ChoiceOption{
			Label: v.Title,
			Handle: func(a *tui.App) {
				g.p.SetRegion(v)
				g.BaseUpdateScreen()
			},
		})
	}

	a.ShowChoicePopup(
		"Выберите регион", "Регион начала игры", regions,
	)
}

func (g *Game) NewGameStartup(a *tui.App) {
	a.ShowFormPopup("Создание игры", []tui.FormField{
		tui.FormField{
			Key:          "player_name",
			Label:        "Сценическое имя",
			Value:        "",
			Placeholder:  "Имя",
			DefaultValue: "John Doe",
		},
		tui.FormField{
			Key:          "group_name",
			Label:        "Название группы",
			Value:        "",
			Placeholder:  "Название группы",
			DefaultValue: "rockability",
		},
	}, func(values map[string]string) {
		g.p.SetName(values["player_name"])
		g.p.SetGroupName(values["group_name"])
		g.ChoiceRegion(a)
	})
}

func (g *Game) LoadGameStartup(a *tui.App) {
	a.ShowMessage("Загрузка игры", "временно недоступно")
	a.AppendLog("Загрузка игры пока недоступна")
}
