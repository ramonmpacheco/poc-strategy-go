package greetingsstrategiesmap

import (
	"fmt"

	greetingsstrategies "github.com/ramonmpacheco/poc-strategy-go/src/domain/greetings/strategy/greetings_strategies"
)

type GreetingsStrategiesMap interface {
	GetBy(code string) (greetingsstrategies.GreetingsStrategy, error)
}

type greetingsStrategiesMap struct {
	Strategies map[string]greetingsstrategies.GreetingsStrategy
}

func NewGreetingsStrategiesMap() GreetingsStrategiesMap {
	return &greetingsStrategiesMap{
		Strategies: getStrategies(),
	}
}

func (gsm *greetingsStrategiesMap) GetBy(code string) (greetingsstrategies.GreetingsStrategy, error) {
	strategy, contains := gsm.Strategies[code]
	if !contains {
		return nil, fmt.Errorf("greeting for code [%s] not found", code)
	}
	return strategy, nil
}

func getStrategies() map[string]greetingsstrategies.GreetingsStrategy {
	return map[string]greetingsstrategies.GreetingsStrategy{
		"GM": &greetingsstrategies.GoodMorningStrategy{},
		"GA": &greetingsstrategies.GoodAfternoonStrategy{},
		"GE": &greetingsstrategies.GoodEveningStrategy{},
	}
}
