package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.POST("/events", createEvent) // auth required
	server.PUT("/events/:id", updateEvent) // auth required
	server.DELETE("/events/:id", deleteEvent) // auth required
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/signup", signup)
	server.POST("/login", login)
}
