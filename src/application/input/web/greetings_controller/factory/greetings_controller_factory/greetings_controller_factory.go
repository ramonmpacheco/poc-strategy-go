package greetingscontrollerfactory

import (
	"github.com/gin-gonic/gin"
	greetingscontroller "github.com/ramonmpacheco/poc-strategy-go/src/application/input/web/greetings_controller"
	greetingsusecaseimpl "github.com/ramonmpacheco/poc-strategy-go/src/domain/port/input/greetings_use_case/greetings_use_case_impl"
)

func AddGreetingsControllers(rest *gin.RouterGroup) error {
	usecase := greetingsusecaseimpl.NewGreetingsUseCaseImpl()
	controller := greetingscontroller.NewGreetingsController(usecase)
	rest.POST("/v1/greetings", controller.Greet)
	return nil
}
