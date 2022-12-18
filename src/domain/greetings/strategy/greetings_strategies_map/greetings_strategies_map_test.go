package greetingsstrategiesmap

import (
	"testing"

	greetingsstrategies "github.com/ramonmpacheco/poc-strategy-go/src/domain/greetings/strategy/greetings_strategies"
	"github.com/stretchr/testify/assert"
)

func Test_Return_Correct_Strategy(t *testing.T) {
	greetingsStrategiesMap := NewGreetingsStrategiesMap()

	goodMorning, err := greetingsStrategiesMap.GetBy("GM")
	assert.Nil(t, err)
	assert.IsType(t, &greetingsstrategies.GoodMorningStrategy{}, goodMorning)

	goodAfternoon, err := greetingsStrategiesMap.GetBy("GA")
	assert.Nil(t, err)
	assert.IsType(t, &greetingsstrategies.GoodAfternoonStrategy{}, goodAfternoon)

	goodEvening, err := greetingsStrategiesMap.GetBy("GE")
	assert.Nil(t, err)
	assert.IsType(t, &greetingsstrategies.GoodEveningStrategy{}, goodEvening)
}

func Test_Return_Error_When_Strategy_Not_Found(t *testing.T) {
	greetingsStrategiesMap := NewGreetingsStrategiesMap()

	wc, err := greetingsStrategiesMap.GetBy("WC")
	assert.NotNil(t, err)
	assert.Empty(t, wc)
}
