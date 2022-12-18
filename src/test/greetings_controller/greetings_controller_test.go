package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	greetingsrequest "github.com/ramonmpacheco/poc-strategy-go/src/application/input/web/request/greetings_request"
	greetingsresponse "github.com/ramonmpacheco/poc-strategy-go/src/application/input/web/response/greetings_response"
	"github.com/ramonmpacheco/poc-strategy-go/src/application/input/web/routes"
	"github.com/stretchr/testify/assert"
)

func Test_Return_Msg_Based_On_Code(t *testing.T) {
	router := gin.Default()
	r := routes.Routes{
		Router: router,
	}
	r.LoadRoutes()

	request := greetingsrequest.GreetingsRequest{Code: "GM"}
	bodyRequest, _ := json.Marshal(request)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/strategy-go/v1/greetings", bytes.NewReader(bodyRequest))
	router.ServeHTTP(w, req)

	var response greetingsresponse.GreetingsResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, response.Msg, "Good Morning!")
}

func Test_Return_Not_Found_Msg(t *testing.T) {
	router := gin.Default()
	r := routes.Routes{
		Router: router,
	}
	r.LoadRoutes()

	request := greetingsrequest.GreetingsRequest{Code: "WC"}
	bodyRequest, _ := json.Marshal(request)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/strategy-go/v1/greetings", bytes.NewReader(bodyRequest))
	router.ServeHTTP(w, req)

	var response greetingsresponse.GreetingsErrorResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "greeting for code [WC] not found", response.Err)
}
