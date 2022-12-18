package greetingsstrategies

type GoodAfternoonStrategy struct {
}

func (gas *GoodAfternoonStrategy) Execute() string {
	return "Good Afternoon!"
}
