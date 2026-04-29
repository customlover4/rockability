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
				g.Improvization()
			},
		},
		{
			Label: "Доработать идеи",
			Handle: func(a *tui.App) {
				g.RewriteSong()
			},
		},
		{
			Label: "Написать трек",
			Handle: func(a *tui.App) {
				g.WriteSongs()
			},
		},
		{
			Label: "Репетировать одному",
			Handle: func(a *tui.App) {
				g.SoloRepetition()
			},
		},
		{
			Label: "Каталог треков",
			Handle: func(a *tui.App) {
				g.ShowSongs()
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

func (g *Game) RecordSong(index int) {
	g.app.ShowChoicePopup(
		"Выберите способ записи",
		"Способ записи влияет на конечный результат",
		[]tui.ChoiceOption{
			{
				Label: "Дома на диктофон - Бесплатно",
				Handle: func(a *tui.App) {
					g.recordSongWithCoefficient(0.15, 0, index)
				},
			},
			{
				Label: "Студия на районе - 200 $",
				Handle: func(a *tui.App) {
					g.recordSongWithCoefficient(0.5, 200, index)
				},
			},
			{
				Label: "Лучшая студия в городе - 1000 $",
				Handle: func(a *tui.App) {
					g.recordSongWithCoefficient(1, 1000, index)
				},
			},
		},
	)
}

func (g *Game) recordSongWithCoefficient(coef float64, price int, index int) {
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
	g.SkipDay(1)
}

//#########################################################################
// BAND REPETITION (INTO SONG MENU)

func (g *Game) BandRepetition(index int) {
	g.p.Band.AddTeamwork(generators.Float64HalfYear(g.rnd))
	g.p.Band.Songs[index].AddSkill(generators.Float64ThreeDays(g.rnd))
	for _, member := range g.p.Band.Members {
		member.AddPlayingSkill(generators.Float64Year(g.rnd))
	}

	g.p.Stats.AddSingingSkill(generators.Float64Year(g.rnd))
	g.p.Stats.AddPlayingSkill(generators.Float64Year(g.rnd))

	g.app.AppendLog("Вы репетировали весь день")
	g.SkipDay(1)
}

//#########################################################################
// SHOW SONGS

func (g *Game) ShowSongs() {
	if len(g.p.Band.Songs) == 0 {
		g.app.ShowMessage("Мои песни", "У тебя пока нет ни одной песни.")
		g.MusicMenuActions()
		return
	}

	actions := make([]tui.Action, 0, len(g.p.Ideas)+1)
	for i := range g.p.Band.Songs {
		index := i
		actions = append(actions, tui.Action{
			Label: g.p.Band.Songs[i].Render(),
			Handle: func(a *tui.App) {
				g.showSongInfo(index)
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

func (g *Game) forgetSong(index int) {
	if index < 0 || index >= len(g.p.Band.Songs) {
		return
	}

	if g.p.Band.Songs[index].Released {
		g.app.ShowMessage(
			"Не могу", "Ты не можешь забыть песню если она уже выпущена",
		)
		return
	}

	g.p.Band.Songs = append(
		g.p.Band.Songs[:index], g.p.Band.Songs[index+1:]...,
	)
}

func (g *Game) showSongInfo(index int) {
	if index < 0 || index >= len(g.p.Band.Songs) {
		return
	}

	s := g.p.Band.Songs[index]

	g.app.ShowChoicePopup(
		s.Title, s.InfoRender(),
		[]tui.ChoiceOption{
			{
				Label: "Репетировать",
				Handle: func(a *tui.App) {
					g.BandRepetition(index)
					g.ShowSongs()
				},
			},
			{
				Label: "Записать",
				Handle: func(a *tui.App) {
					g.RecordSong(index)
					g.ShowSongs()
				},
			},
			{
				Label: "Забыть",
				Handle: func(a *tui.App) {
					g.forgetSong(index)
					g.ShowSongs()
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

func (g *Game) WriteSongs() {
	if len(g.p.Ideas) == 0 {
		g.app.ShowMessage("Написание трека", "У тебя пока нет идей для песен.")
		g.MusicMenuActions()
		return
	}

	actions := make([]tui.Action, 0, len(g.p.Ideas)+1)
	for i := range g.p.Ideas {
		index := i
		actions = append(actions, tui.Action{
			Label: g.p.Ideas[i].Render(),
			Handle: func(a *tui.App) {
				g.showIdeaSongConv(index)
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

func (g *Game) showIdeaSongConv(index int) {
	if index < 0 || index >= len(g.p.Ideas) {
		return
	}

	g.app.ShowChoicePopup(
		"Идея",
		g.p.Ideas[index].Render(),
		[]tui.ChoiceOption{
			{
				Label: "Написать песню",
				Handle: func(a *tui.App) {
					g.writeSongFromIdea(index)
				},
			},
			{
				Label:  "Закрыть",
				Handle: func(a *tui.App) {},
			},
		},
	)
}

func (g *Game) writeSongFromIdea(index int) {
	if index < 0 || index >= len(g.p.Ideas) {
		return
	}
	if g.p.Stats.Inspiration < 20 {
		g.app.ShowMessage(
			"Нет вдохновения",
			"Пока что у тебя нет вдохновения, попробуй позже. Необходимо >20",
		)
		return
	}

	g.app.ShowFormPopup("Название песни",
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

			g.app.AppendLog("Песня написана")
			g.p.Stats.AddInspiration(-20)
			g.SkipDay(1)
			g.app.SetStatus(g.p.RenderStatus())
			g.MusicMenuActions()
		},
	)
}

//#########################################################################
// SOLO REPETITION

func (g *Game) SoloRepetition() {
	g.p.Stats.AddSingingSkill(generators.Float64Year(g.rnd))
	g.p.Stats.AddPlayingSkill(generators.Float64Year(g.rnd))

	g.app.AppendLog("Ты упражнялся и улучшил свои навыки игры.")
	g.SkipDay(1)
}

//#########################################################################
// IMPROVIZATION

func (g *Game) Improvization() {
	if g.p.Stats.Inspiration < 2 {
		g.app.ShowMessage(
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

	g.app.AppendLog("Ты импровизируешь на гитаре и продвигаешься к новой идее.")

	for g.p.ImprovizationXP >= 100 {
		g.p.ImprovizationXP -= 100

		g.p.Ideas = append(g.p.Ideas, music.NewIdea())

		g.app.AppendLog("Во время импровизации у тебя родилась новая идея для трека.")
	}

	g.SkipDay(1)
	g.app.SetStatus(g.p.RenderStatus())
}

// ########################################################################
// REWRITE SONG

func (g *Game) RewriteSong() {
	if len(g.p.Ideas) == 0 {
		g.app.ShowMessage("Доработка идей", "У тебя пока нет идей для доработки.")
		g.MusicMenuActions()
		return
	}

	actions := make([]tui.Action, 0, len(g.p.Ideas)+1)
	for i := range g.p.Ideas {
		index := i
		actions = append(actions, tui.Action{
			Label: g.p.Ideas[i].Render(),
			Handle: func(a *tui.App) {
				g.showIdeaActions(index)
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

func (g *Game) showIdeaActions(ideaIndex int) {
	if ideaIndex < 0 || ideaIndex >= len(g.p.Ideas) {
		return
	}

	g.app.ShowChoicePopup(
		"Идея",
		g.p.Ideas[ideaIndex].Render(),
		[]tui.ChoiceOption{
			{
				Label: "Доработать идею",
				Handle: func(a *tui.App) {
					g.improveIdea(ideaIndex)
				},
			},
			{
				Label: "Забыть идею",
				Handle: func(a *tui.App) {
					g.forgetIdea(ideaIndex)
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
func (g *Game) improveIdea(ideaIndex int) {
	if ideaIndex < 0 || ideaIndex >= len(g.p.Ideas) {
		return
	}
	if g.p.Stats.Inspiration < 5 {
		g.app.ShowMessage(
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
		g.app.AppendLog(fmt.Sprintf("Потенциал идеи вырос на %.1f.", boost))
		improved = true
	}

	if improved {
		g.app.AppendLog("Ты доработал идею и выжал из нее что-то стоящее.")
	} else {
		g.app.AppendLog("Ты попытался доработать идею, но в этот раз она не стала лучше.")
	}

	g.RewriteSong()
	g.p.Stats.AddInspiration(-5)
	g.SkipDay(1)
}

func (g *Game) forgetIdea(ideaIndex int) {
	if ideaIndex < 0 || ideaIndex >= len(g.p.Ideas) {
		return
	}

	g.p.Ideas = append(g.p.Ideas[:ideaIndex], g.p.Ideas[ideaIndex+1:]...)
	g.app.AppendLog("Ты решил отпустить эту идею и больше к ней не возвращаться.")
	g.RewriteSong()
}
