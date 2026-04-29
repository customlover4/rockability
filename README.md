# Console Menu UI Framework

Это чистый TUI-каркас для консольной игры на Go.  
В нем нет зашитой игровой сцены, таверны, улицы или движка правил.  
Он оставляет тебе только интерфейсные кубики:

- левая колонка действий
- правая панель лога
- нижняя строка статуса
- popup сообщений
- popup выбора стрелками
- popup формы
- полноэкранная катсцена

## Структура

`main.go`  
Небольшой пример использования API.

`tui/entities.go`  
Публичные типы: `App`, `Screen`, `Action`, `StatusItem`, `ChoiceOption`, `FormField`, `Cutscene`, `Theme`.

`tui/app.go`  
Основной рендер и базовый API приложения.

`tui/dialogs.go`  
Все popup-окна: сообщение, выбор, форма.

`tui/cutscene.go`  
Полноэкранная психоделическая катсцена.

`tui/theme.go`  
Цвета и тема.

## Базовая идея

Ты хранишь свою игровую логику отдельно.  
UI ничего не знает про твои квесты, бои, инвентарь и сцены.

Ты просто:

1. создаешь `app := tui.NewApp()`
2. вызываешь `app.SetScreen(...)`
3. передаешь список действий с callback-ами
4. внутри callback-ов меняешь свои данные и снова обновляешь экран

## Основной API

### Создание и запуск

```go
app := tui.NewApp()
err := app.Run()
```

### Основной экран

```go
app.SetScreen(tui.Screen{
    ActionTitle: "Команды",
    LogTitle:    "События",
    StatusTitle: "Статус",
    Log: []string{
        "Игра запущена.",
    },
    Status: []tui.StatusItem{
        {Label: "HP", Value: "10/10"},
        {Label: "Gold", Value: "25"},
    },
    Actions: []tui.Action{
        {
            Label: "Осмотреться",
            Handle: func(app *tui.App) {
                app.AppendLog("Ты осматриваешься.")
            },
        },
    },
})
```

### Обновление лога

```go
app.AppendLog("Новая строка")
app.SetLog([]string{"Полная", "замена", "журнала"})
app.ClearLog()
```

### Обновление статуса

```go
app.SetStatus([]tui.StatusItem{
    {Label: "HP", Value: "8/10"},
    {Label: "Gold", Value: "30"},
})
```

### Замена действий

```go
app.SetActions([]tui.Action{
    {
        Label: "Открыть сундук",
        Handle: func(app *tui.App) {
            app.AppendLog("Сундук открыт.")
        },
    },
})
```

### Popup сообщения

```go
app.ShowMessage("Внимание", "У тебя недостаточно золота.")
```

### Popup выбора стрелками

```go
app.ShowChoicePopup("Куда идти?", "Выбери направление.", []tui.ChoiceOption{
    {
        Label: "На север",
        Handle: func(app *tui.App) {
            app.AppendLog("Ты идешь на север.")
        },
    },
    {
        Label: "На юг",
        Handle: func(app *tui.App) {
            app.AppendLog("Ты идешь на юг.")
        },
    },
})
```

### Popup формы

```go
app.ShowFormPopup("Создать персонажа", []tui.FormField{
    {Key: "name", Label: "Имя", DefaultValue: "Игрок"},
    {Key: "class", Label: "Класс", DefaultValue: "Странник"},
}, func(values map[string]string) {
    name := values["name"]
    class := values["class"]
    app.AppendLog("Создан персонаж: " + name + " / " + class)
})
```

### Катсцена

```go
app.PlayCutscene(tui.Cutscene{
    Title: "Видение",
    Messages: []string{
        "Мир дрожит.",
        "Шум становится цветом.",
    },
    Duration: 4 * time.Second,
})
```

## Полезный паттерн

Самый удобный способ использовать этот каркас:

1. хранить свою игру в собственной структуре, например `Game`
2. держать функцию `render()`
3. внутри `render()` вызывать `app.SetScreen(...)`
4. в `Handle` менять состояние игры и снова вызывать `render()`

Схема такая:

```go
type Game struct {
    hp int
    gold int
}
```

```go
func (g *Game) render(app *tui.App) {
    app.SetScreen(...)
}
```

Так у тебя логика игры остается отдельно, а `tui` остается только рендером.
