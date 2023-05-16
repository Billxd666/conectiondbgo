package main

import (
	"github.com/Billxd666/conectiondbgo.git/initializers"
	"github.com/Billxd666/conectiondbgo.git/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.Post{})
}
