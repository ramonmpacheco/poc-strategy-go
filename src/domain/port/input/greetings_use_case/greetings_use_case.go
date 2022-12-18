package greetingsusecase

import greetingsrequest "github.com/ramonmpacheco/poc-strategy-go/src/application/input/web/request/greetings_request"

type GreetingsUseCase interface {
	Greet(code greetingsrequest.GreetingsRequest) (string, error)
}
