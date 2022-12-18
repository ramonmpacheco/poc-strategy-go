package greetingsresponse

type GreetingsResponse struct {
	Msg string `json:"msg"`
}

type GreetingsErrorResponse struct {
	Err string `json:"err"`
}
