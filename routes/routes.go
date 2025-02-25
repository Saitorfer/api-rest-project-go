package routes

import (
	"github.com/gin-gonic/gin"
	"practise.com/rest-api-go/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	//events
	server.GET("/events", GetEvents)
	server.GET("/events/:id", GetEvent)
	//we protect the createEvent (we need to be sure the token is valid)
	//so we use the middleware Authenticate
	server.POST("/events", middlewares.Authenticate, CreateEvent)
	server.PUT("/events/:id", middlewares.Authenticate, UpdateEvent)
	server.DELETE("/events/:id", middlewares.Authenticate, DeleteEvent)
	//user
	server.POST("/signup", Signup)
	server.POST("/login", Login)
}
