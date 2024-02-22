package controllers

import (
	"learning/go-crud/initializers"
	"learning/go-crud/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)

	if body.Title == "" {
		c.JSON(400, gin.H{
			"error": "Title is required",
		})
		return
	}

	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsShow(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Respond wit them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsIndex(c *gin.Context) {
	// Get id off url
	id := c.Param("id")
	// Show single post
	var post models.Post
	result := initializers.DB.First(&post, id)

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{
			"error": "Post not found",
		})
		return
	}

	// Respond wit them
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")

	// Get the data off req body
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	// Find the post were updating
	var post models.Post
	result := initializers.DB.First(&post, id)

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{
			"error": "Post not found",
		})
		return
	}

	// Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")

	// Delete the post
	result := initializers.DB.Delete(&models.Post{}, id)

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{
			"error": "Post not found",
		})
		return
	}

	// Respond
	c.Status(200)
}
