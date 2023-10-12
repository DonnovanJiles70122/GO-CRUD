package main

import (
	"go-crud/controllers"  // import controllers package
	"go-crud/initializers" // import initializers package

	"github.com/gin-gonic/gin" // import gin framework
)

// Initialize the environment variables and database connection
func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default() // Create a default gin router

	// Define the routes for the posts
	r.POST("/posts", controllers.PostsCreate) // Create a new post
	r.PUT("/posts/:id", controllers.PostUpdate) // Update a specific post
	r.GET("/posts", controllers.PostIndex) // Get all the posts
	r.GET("/posts/:id", controllers.PostShow) // Get a specific post
	r.DELETE("/posts/:id", controllers.PostDelete) // Delete a specific post

	r.Run() // Start the server
}

