package greetingscontext

import greetingsstrategies "github.com/ramonmpacheco/poc-strategy-go/src/domain/greetings/strategy/greetings_strategies"

type greetingsContext struct {
}

func NewGreetingsContext() greetingsContext {
	return greetingsContext{}
}

func (gc *greetingsContext) Greet(str greetingsstrategies.GreetingsStrategy) string {
	if str == nil {
		return "Invalid strategy: nil"
	}
	return str.Execute()
}
