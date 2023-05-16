package main

import (
	"github.com/Billxd666/conectiondbgo.git/controllers"
	"github.com/Billxd666/conectiondbgo.git/initializers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConectToDB()
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostIndex)
	r.GET("/posts/:id", controllers.PostShow)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.DELETE("/posts/:id", controllers.PostDelete)
	r.Run() // listen and serve on 0.0.0.0:8080
}
