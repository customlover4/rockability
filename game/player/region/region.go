package region

type Region struct {
	Title       string
	MoneyFormat string
}

func DefaultRegion() Region {
	return Region{
		Title:       "Не выбрано",
		MoneyFormat: "?",
	}
}

var ChoiceRegion = map[int]Region{
	1: {
		Title:       "USA",
		MoneyFormat: "$",
	},
}
