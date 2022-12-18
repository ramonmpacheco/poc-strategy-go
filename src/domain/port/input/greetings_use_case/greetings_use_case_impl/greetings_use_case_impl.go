package greetingsusecaseimpl

import (
	greetingsrequest "github.com/ramonmpacheco/poc-strategy-go/src/application/input/web/request/greetings_request"

	greetingscontext "github.com/ramonmpacheco/poc-strategy-go/src/domain/greetings/strategy/greetings_context"
	greetingsstrategiesmap "github.com/ramonmpacheco/poc-strategy-go/src/domain/greetings/strategy/greetings_strategies_map"
	greetingsusecase "github.com/ramonmpacheco/poc-strategy-go/src/domain/port/input/greetings_use_case"
)

type greetingsUseCaseImpl struct {
	greetingsStrategies greetingsstrategiesmap.GreetingsStrategiesMap
}

func NewGreetingsUseCaseImpl() greetingsusecase.GreetingsUseCase {
	return &greetingsUseCaseImpl{
		greetingsStrategies: greetingsstrategiesmap.NewGreetingsStrategiesMap(),
	}
}

func (guci *greetingsUseCaseImpl) Greet(request greetingsrequest.GreetingsRequest) (string, error) {
	s, err := guci.greetingsStrategies.GetBy(request.Code)
	if err != nil {
		return "", err
	}
	context := greetingscontext.NewGreetingsContext()
	return context.Greet(s), nil
}
