package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	greetingscontrollerfactory "github.com/ramonmpacheco/poc-strategy-go/src/application/input/web/greetings_controller/factory/greetings_controller_factory"
)

type Routes struct {
	Router *gin.Engine
}

func (r *Routes) LoadRoutes() (err error) {
	r.Router.Use(CORSMiddleware())
	rest := r.Router.Group("/strategy-go")
	loadGreetingsControllers(rest)
	return
}

func loadGreetingsControllers(rest *gin.RouterGroup) {
	err := greetingscontrollerfactory.AddGreetingsControllers(rest)
	if err != nil {
		log.Fatalf(err.Error())
	}

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
