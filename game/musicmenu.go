package game

import (
	"fmt"

	"game/game/player/music"
	"game/generators"
	"game/tui"
)

func (g *Game) MusicMenuActions() {
	g.app.SetActions([]tui.Action{
		{
			Label: "Импровизировать на гитаре",
			Handle: func(a *tui.App) {
				g.Improvization(a)
			},
		},
		{
			Label: "Доработать идеи",
			Handle: func(a *tui.App) {
				g.RewriteSong(a)
			},
		},
		{
			Label: "Написать трек",
			Handle: func(a *tui.App) {
				g.WriteSongs(a)
			},
		},
		{
			Label: "Репетировать одному",
			Handle: func(a *tui.App) {
				g.SoloRepetition(a)
			},
		},
		{
			Label: "Каталог треков",
			Handle: func(a *tui.App) {
				g.ShowSongs(a)
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

//#########################################################################
// RECORD SONG (INTO SONG MENU)

func (g *Game) RecordSong(a *tui.App, index int) {
	a.ShowChoicePopup(
		"Выберите способ записи",
		"Способ записи влияет на конечный результат",
		[]tui.ChoiceOption{
			{
				Label: "Дома на диктофон - Бесплатно",
				Handle: func(a *tui.App) {
					g.recordSongWithCoefficient(0.15, 0, a, index)
				},
			},
			{
				Label: "Студия на районе - 200 $",
				Handle: func(a *tui.App) {
					g.recordSongWithCoefficient(0.5, 200, a, index)
				},
			},
			{
				Label: "Лучшая студия в городе - 1000 $",
				Handle: func(a *tui.App) {
					g.recordSongWithCoefficient(1, 1000, a, index)
				},
			},
		},
	)
}

func (g *Game) recordSongWithCoefficient(coef float64, price int, a *tui.App, index int) {
	if index < 0 || index >= len(g.p.Band.Songs) {
		return
	}
	if !g.p.AddMoney(-price) {
		g.app.ShowMessage("Деньги", "У вас недостаточно денег для записи")
		return
	}

	g.p.Band.Songs[index].SetRecordQuality(
		(g.p.Band.Songs[index].Skill * coef),
	)
	g.p.Band.Songs[index].Record()

	g.app.AppendLog("Песня успешно записана")
	g.p.SkipDay(a)
}

//#########################################################################
// BAND REPETITION (INTO SONG MENU)

func (g *Game) BandRepetition(a *tui.App, index int) {
	g.p.Band.AddTeamwork(generators.Float64HalfYear(g.rnd))
	g.p.Band.Songs[index].AddSkill(generators.Float64ThreeDays(g.rnd))
	for _, member := range g.p.Band.Members {
		member.AddPlayingSkill(generators.Float64Year(g.rnd))
	}

	g.p.Stats.AddSingingSkill(generators.Float64Year(g.rnd))
	g.p.Stats.AddPlayingSkill(generators.Float64Year(g.rnd))

	g.app.AppendLog("Вы репетировали весь день")
	g.p.SkipDay(a)
}

//#########################################################################
// SHOW SONGS

func (g *Game) ShowSongs(a *tui.App) {
	if len(g.p.Band.Songs) == 0 {
		a.ShowMessage("Мои песни", "У тебя пока нет ни одной песни.")
		g.MusicMenuActions()
		return
	}

	actions := make([]tui.Action, 0, len(g.p.Ideas)+1)
	for i := range g.p.Band.Songs {
		index := i
		actions = append(actions, tui.Action{
			Label: g.p.Band.Songs[i].Render(),
			Handle: func(a *tui.App) {
				g.showSongInfo(a, index)
			},
		})
	}

	actions = append(actions, tui.Action{
		Label: "Назад",
		Handle: func(a *tui.App) {
			g.MusicMenuActions()
		},
	})

	g.app.SetActions(actions)
}

func (g *Game) forgetSong(a *tui.App, index int) {
	if index < 0 || index >= len(g.p.Band.Songs) {
		return
	}

	if g.p.Band.Songs[index].Released {
		a.ShowMessage(
			"Не могу", "Ты не можешь забыть песню если она уже выпущена",
		)
		return
	}

	g.p.Band.Songs = append(
		g.p.Band.Songs[:index], g.p.Band.Songs[index+1:]...,
	)
}

func (g *Game) showSongInfo(a *tui.App, index int) {
	if index < 0 || index >= len(g.p.Band.Songs) {
		return
	}

	s := g.p.Band.Songs[index]

	a.ShowChoicePopup(
		s.Title, s.InfoRender(),
		[]tui.ChoiceOption{
			{
				Label: "Репетировать",
				Handle: func(a *tui.App) {
					g.BandRepetition(a, index)
					g.ShowSongs(a)
				},
			},
			{
				Label: "Записать",
				Handle: func(a *tui.App) {
					g.RecordSong(a, index)
					g.ShowSongs(a)
				},
			},
			{
				Label: "Забыть",
				Handle: func(a *tui.App) {
					g.forgetSong(a, index)
					g.ShowSongs(a)
				},
			},
			{
				Label:  "Закрыть",
				Handle: func(a *tui.App) {},
			},
		},
	)
}

//#########################################################################
// WRITE SONGS

func (g *Game) WriteSongs(a *tui.App) {
	if len(g.p.Ideas) == 0 {
		a.ShowMessage("Написание трека", "У тебя пока нет идей для песен.")
		g.MusicMenuActions()
		return
	}

	actions := make([]tui.Action, 0, len(g.p.Ideas)+1)
	for i := range g.p.Ideas {
		index := i
		actions = append(actions, tui.Action{
			Label: g.p.Ideas[i].Render(),
			Handle: func(a *tui.App) {
				g.showIdeaSongConv(a, index)
			},
		})
	}

	actions = append(actions, tui.Action{
		Label: "Назад",
		Handle: func(a *tui.App) {
			g.MusicMenuActions()
		},
	})

	g.app.SetActions(actions)
}

func (g *Game) showIdeaSongConv(a *tui.App, index int) {
	if index < 0 || index >= len(g.p.Ideas) {
		return
	}

	a.ShowChoicePopup(
		"Идея",
		g.p.Ideas[index].Render(),
		[]tui.ChoiceOption{
			{
				Label: "Написать песню",
				Handle: func(a *tui.App) {
					g.writeSongFromIdea(a, index)
				},
			},
			{
				Label:  "Закрыть",
				Handle: func(a *tui.App) {},
			},
		},
	)
}

func (g *Game) writeSongFromIdea(a *tui.App, index int) {
	if index < 0 || index >= len(g.p.Ideas) {
		return
	}
	if g.p.Stats.Inspiration < 20 {
		a.ShowMessage(
			"Нет вдохновения",
			"Пока что у тебя нет вдохновения, попробуй позже. Необходимо >20",
		)
		return
	}

	a.ShowFormPopup("Название песни",
		[]tui.FormField{
			{
				Key:          "song_title",
				Label:        "Название песни",
				Value:        "",
				Placeholder:  "Название",
				DefaultValue: "Без названия",
			},
		}, func(values map[string]string) {
			tmp := g.p.Ideas[index]
			g.p.Ideas = append(g.p.Ideas[:index], g.p.Ideas[index+1:]...)

			g.p.Band.Songs = append(
				g.p.Band.Songs,
				music.NewSong(
					values["song_title"], g.p.Stats, g.rnd.Float64(),
					tmp.Potential,
				),
			)

			a.AppendLog("Песня написана")
			g.p.Stats.AddInspiration(-20)
			g.p.SkipDay(a)
			a.SetStatus(g.p.RenderStatus())
			g.MusicMenuActions()
		},
	)
}

//#########################################################################
// SOLO REPETITION

func (g *Game) SoloRepetition(a *tui.App) {
	g.p.Stats.AddSingingSkill(generators.Float64Year(g.rnd))
	g.p.Stats.AddPlayingSkill(generators.Float64Year(g.rnd))

	a.AppendLog("Ты упражнялся и улучшил свои навыки игры.")
	g.p.SkipDay(a)
}

//#########################################################################
// IMPROVIZATION

func (g *Game) Improvization(a *tui.App) {
	if g.p.Stats.Inspiration < 2 {
		a.ShowMessage(
			"Нет вдохновения",
			"Пока что у тебя нет вдохновения, попробуй позже. Необходимо >2",
		)
		return
	}

	gain := generators.Float64NineDays(g.rnd)
	g.p.ImprovizationXP += gain
	g.p.Stats.AddInspiration(-2)
	g.p.Stats.AddPlayingSkill(generators.Float64Year(g.rnd))
	g.p.Stats.AddMusicWritingSkill(generators.Float64Year(g.rnd))

	a.AppendLog("Ты импровизируешь на гитаре и продвигаешься к новой идее.")

	for g.p.ImprovizationXP >= 100 {
		g.p.ImprovizationXP -= 100

		g.p.Ideas = append(g.p.Ideas, music.NewIdea())

		a.AppendLog("Во время импровизации у тебя родилась новая идея для трека.")
	}

	g.p.SkipDay(a)
	a.SetStatus(g.p.RenderStatus())
}

// ########################################################################
// REWRITE SONG

func (g *Game) RewriteSong(a *tui.App) {
	if len(g.p.Ideas) == 0 {
		a.ShowMessage("Доработка идей", "У тебя пока нет идей для доработки.")
		g.MusicMenuActions()
		return
	}

	actions := make([]tui.Action, 0, len(g.p.Ideas)+1)
	for i := range g.p.Ideas {
		index := i
		actions = append(actions, tui.Action{
			Label: g.p.Ideas[i].Render(),
			Handle: func(a *tui.App) {
				g.showIdeaActions(a, index)
			},
		})
	}

	actions = append(actions, tui.Action{
		Label: "Назад",
		Handle: func(a *tui.App) {
			g.MusicMenuActions()
		},
	})

	g.app.SetActions(actions)
}

func (g *Game) showIdeaActions(a *tui.App, ideaIndex int) {
	if ideaIndex < 0 || ideaIndex >= len(g.p.Ideas) {
		return
	}

	a.ShowChoicePopup(
		"Идея",
		g.p.Ideas[ideaIndex].Render(),
		[]tui.ChoiceOption{
			{
				Label: "Доработать идею",
				Handle: func(a *tui.App) {
					g.improveIdea(a, ideaIndex)
				},
			},
			{
				Label: "Забыть идею",
				Handle: func(a *tui.App) {
					g.forgetIdea(a, ideaIndex)
				},
			},
			{
				Label:  "Закрыть",
				Handle: func(a *tui.App) {},
			},
		},
	)
}

// todo: maybe trubles with optimization :)
func (g *Game) improveIdea(a *tui.App, ideaIndex int) {
	if ideaIndex < 0 || ideaIndex >= len(g.p.Ideas) {
		return
	}
	if g.p.Stats.Inspiration < 5 {
		a.ShowMessage(
			"Нет вдохновения",
			"Пока что у тебя нет вдохновения, попробуй позже.  Необходимо >5",
		)
		return
	}

	idea := &g.p.Ideas[ideaIndex]
	improved := false

	if g.rnd.Float64() < 0.45 {
		boost := 2.0 + generators.Float64FifteenDays(g.rnd)
		idea.AddPotential(boost)
		a.AppendLog(fmt.Sprintf("Потенциал идеи вырос на %.1f.", boost))
		improved = true
	}

	if improved {
		a.AppendLog("Ты доработал идею и выжал из нее что-то стоящее.")
	} else {
		a.AppendLog("Ты попытался доработать идею, но в этот раз она не стала лучше.")
	}

	g.RewriteSong(a)
	g.p.Stats.AddInspiration(-5)
	g.p.SkipDay(a)
}

func (g *Game) forgetIdea(a *tui.App, ideaIndex int) {
	if ideaIndex < 0 || ideaIndex >= len(g.p.Ideas) {
		return
	}

	g.p.Ideas = append(g.p.Ideas[:ideaIndex], g.p.Ideas[ideaIndex+1:]...)
	a.AppendLog("Ты решил отпустить эту идею и больше к ней не возвращаться.")
	g.RewriteSong(a)
}
