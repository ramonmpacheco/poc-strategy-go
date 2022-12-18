package greetingsusecaseimpl

import (
	"testing"

	greetingsrequest "github.com/ramonmpacheco/poc-strategy-go/src/application/input/web/request/greetings_request"
	"github.com/stretchr/testify/assert"
)

func Test_Use_Strategy_According_To_Code(t *testing.T) {

	usecase := NewGreetingsUseCaseImpl()
	morningResult, e := usecase.Greet(greetingsrequest.GreetingsRequest{Code: "GM"})
	assert.Nil(t, e)
	afternoonResult, e := usecase.Greet(greetingsrequest.GreetingsRequest{Code: "GA"})
	assert.Nil(t, e)
	eveningResult, e := usecase.Greet(greetingsrequest.GreetingsRequest{Code: "GE"})
	assert.Nil(t, e)

	assert.EqualValues(t, "Good Morning!", morningResult)
	assert.EqualValues(t, "Good Afternoon!", afternoonResult)
	assert.EqualValues(t, "Good Evening!", eveningResult)
}

func Test_Not_Use_Strategy_When_Code_Not_Found(t *testing.T) {

	usecase := NewGreetingsUseCaseImpl()
	result, e := usecase.Greet(greetingsrequest.GreetingsRequest{Code: "WC"})
	assert.Empty(t, result)
	assert.NotNil(t, e)

	assert.EqualValues(t, "greeting for code [WC] not found", e.Error())
}
