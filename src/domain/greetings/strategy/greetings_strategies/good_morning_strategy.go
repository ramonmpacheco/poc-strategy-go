package greetingsstrategies

type GoodMorningStrategy struct {
}

func (gms *GoodMorningStrategy) Execute() string {
	return "Good Morning!"
}
