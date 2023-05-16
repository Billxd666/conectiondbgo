package controllers

import (
	"github.com/Billxd666/conectiondbgo.git/initializers"
	"github.com/Billxd666/conectiondbgo.git/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	//Get data off req body
	var body struct {
		Title   string
		Content string
	}

	c.Bind(&body)

	//Create a post
	post := models.Post{Title: body.Title, Content: body.Content}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostIndex(c *gin.Context) {
	//Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"post": posts,
	})
}

func PostShow(c *gin.Context) {
	//Get id of url
	id := c.Param("id")

	//Get the posts
	var post models.Post
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"posts": post,
	})
}

func PostUpdate(c *gin.Context) {
	id := c.Param("id")
	//Get the data of req body
	var body struct {
		Title   string
		Content string
	}

	c.Bind(&body)

	//Find the post were updating
	var post models.Post
	initializers.DB.First(&post, id)
	//Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title:   body.Title,
		Content: body.Content,
	})

	//Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	//Get the id off the url
	id := c.Param("id")

	//Delete the posts
	initializers.DB.Delete(&models.Post{}, id)

	//respond
	c.Status(200)
}
