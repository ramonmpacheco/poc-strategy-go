package greetingsstrategies

type GoodEveningStrategy struct {
}

func (gms *GoodEveningStrategy) Execute() string {
	return "Good Evening!"
}
