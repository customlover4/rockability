package generators

import "math/rand/v2"

func Float64TwoDays(rnd *rand.Rand) float64 {
	return rnd.Float64() * 100
}

func Float64ThreeDays(rnd *rand.Rand) float64 {
	return rnd.Float64() * 75
}

func Float64FiveDays(rnd *rand.Rand) float64 {
	return rnd.Float64() * 50
}

func Float64NineDays(rnd *rand.Rand) float64 {
	return rnd.Float64() * 25
}

func Float64FifteenDays(rnd *rand.Rand) float64 {
	return rnd.Float64() * 13.04
}

func Float64TwentyDays(rnd *rand.Rand) float64 {
	return rnd.Float64() * 10
}

func Float64Month(rnd *rand.Rand) float64 {
	return rnd.Float64() * 7
}

func Float64TwoMonth(rnd *rand.Rand) float64 {
	return rnd.Float64() * 3
}

func Float64ThreeMonth(rnd *rand.Rand) float64 {
	return rnd.Float64() * 2
}

func Float64HalfYear(rnd *rand.Rand) float64 {
	return rnd.Float64() * 1.09
}

func Float64Year(rnd *rand.Rand) float64 {
	return rnd.Float64() * 0.549
}

func Float64TwoYear(rnd *rand.Rand) float64 {
	return rnd.Float64() * 0.2728
}

func Float64ThreeYear(rnd *rand.Rand) float64 {
	return rnd.Float64() * 0.187
}
