package main

import (
	"learning/go-crud/controllers"
	"learning/go-crud/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/posts", controllers.PostsCreate)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.GET("/posts", controllers.PostsShow)
	r.GET("/posts/:id", controllers.PostsIndex)
	r.DELETE("/posts/:id", controllers.PostsDelete)

	r.Run()
}
