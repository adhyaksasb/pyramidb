package controllers

import (
	"github.com/adhyaksasb/pyramidb/initializers"
	model "github.com/adhyaksasb/pyramidb/models"

	"github.com/gin-gonic/gin"
)

func TestUser (c *gin.Context) {
	c.JSON(200, gin.H{
		"data": gin.H{
			"id": c.Param("id"),
		},
	})
}

func CreatePosts (c *gin.Context) {
	// Get data off req body
	var body struct {
		Body string
		Title string
	}

	c.Bind(&body)

	// Create a post
	post := model.Post{Title: body.Title, Body: body.Body}

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

func IndexPosts (c *gin.Context) {
	// Get the posts
	var posts []model.Post
	initializers.DB.Find(&posts)

	// Return it
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func ShowPost (c *gin.Context) {
	// Get id parmas
	id := c.Param("id")

	// Get the post
	var post model.Post
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})
}

func UpdatePost (c *gin.Context) {
	// Get id parmas
	id := c.Param("id")

	var body struct {
		Body string
		Title string
	}

	c.Bind(&body)

	// Get the post
	var post model.Post
	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(model.Post{
		Title: body.Title,
		Body: body.Body,
	})

	c.JSON(200, gin.H{
		"post": post,
	})
}

func DeletePost (c *gin.Context) {
	// Get id parmas
	id := c.Param("id")

	var body struct {
		Body string
		Title string
	}

	c.Bind(&body)

	// Get the post
	initializers.DB.Delete(&model.Post{}, id)

	c.JSON(200, gin.H{
		"message": "Post deleted successfully",
	})
}