package greetingscontext

import (
	"testing"

	greetingsstrategies "github.com/ramonmpacheco/poc-strategy-go/src/domain/greetings/strategy/greetings_strategies"
	"github.com/stretchr/testify/assert"
)

func Test_Call_Strategy(t *testing.T) {
	context := NewGreetingsContext()
	msg := context.Greet(
		&greetingsstrategies.GoodAfternoonStrategy{},
	)
	assert.EqualValues(t, "Good Afternoon!", msg)
}

func Test_Get_Invalid_Msg_When_SrOn_Call_Strategy(t *testing.T) {
	context := NewGreetingsContext()
	msg := context.Greet(nil)
	assert.EqualValues(t, "Invalid strategy: nil", msg)
}
