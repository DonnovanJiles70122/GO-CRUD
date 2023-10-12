package controllers

import (
	"go-crud/initializers"
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

// PostsCreate is the controller function for creating a new post
func PostsCreate(c *gin.Context) {
	
	// Define the request body structure
	var body struct {
		Body string
		Title string
	}

	// Bind the request body to the defined structure
	c.Bind(&body)

	// Create a new Post model with the request body
	post := models.Post{Title: body.Title, Body: body.Body}

	// Create the new post in the database
	result := initializers.DB.Create(&post)

	// If there is an error, return a 400 status
	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return the created post as JSON
	c.JSON(200, gin.H{
		"post": post,
	})
}

// PostIndex is the controller function for fetching all posts
func PostIndex(c *gin.Context) {
	// Define a slice to hold the posts
	var posts []models.Post

	// Fetch all posts from the database
	initializers.DB.Find(&posts)

	// Return the fetched posts as JSON
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

// PostShow is the controller function for fetching a single post
func PostShow(c *gin.Context) {
	// Fetch the id from the request parameters
	id := c.Param("id")

	// Define a Post model to hold the fetched post
	var post models.Post

	// Fetch the post from the database
	initializers.DB.First(&post, id)

	// Return the fetched post as JSON
	c.JSON(200, gin.H{
		"post": post,
	})
}

// PostUpdate is the controller function for updating a post
func PostUpdate(c *gin.Context) {
	// Fetch the id from the request parameters
	id := c.Param("id")

	// Define the request body structure
	var body struct {
		Body string
		Title string
	}

	// Bind the request body to the defined structure
	c.Bind(&body)

	// Define a Post model to hold the fetched post
	var post models.Post

	// Fetch the post from the database
	initializers.DB.First(&post, id)

	// Update the fetched post with the request body
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body: body.Body,
	})

	// Return the updated post as JSON
	c.JSON(200, gin.H{
		"post": post,
	})
}

// PostDelete is the controller function for deleting a post
func PostDelete(c *gin.Context) {
	// Fetch the id from the request parameters
	id := c.Param("id")

	// Delete the post from the database
	initializers.DB.Delete(&models.Post{}, id)

	// Return a 200 status
	c.Status(200)
}
