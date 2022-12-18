package greetingscontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	greetingsrequest "github.com/ramonmpacheco/poc-strategy-go/src/application/input/web/request/greetings_request"
	greetingsresponse "github.com/ramonmpacheco/poc-strategy-go/src/application/input/web/response/greetings_response"
	greetingsusecase "github.com/ramonmpacheco/poc-strategy-go/src/domain/port/input/greetings_use_case"
)

type greetingsController struct {
	greetingsUseCase greetingsusecase.GreetingsUseCase
}

func NewGreetingsController(guc greetingsusecase.GreetingsUseCase) *greetingsController {
	return &greetingsController{
		greetingsUseCase: guc,
	}
}

func (gc *greetingsController) Greet(c *gin.Context) {
	var request greetingsrequest.GreetingsRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		errMsg := greetingsresponse.GreetingsErrorResponse{Err: err.Error()}
		c.JSON(http.StatusBadRequest, &errMsg)
		return
	}

	msg, err := gc.greetingsUseCase.Greet(request)
	if err != nil {
		errMsg := greetingsresponse.GreetingsErrorResponse{Err: err.Error()}
		c.JSON(http.StatusBadRequest, &errMsg)
		return
	}
	response := greetingsresponse.GreetingsResponse{Msg: msg}
	c.JSON(http.StatusOK, &response)
}
