// This is the main package, the entry point of the application
package main

// Importing necessary packages. initializers package is used for initializing the application like loading environment variables, connecting to the DB. models package contains the data structures.
import (
	"go-crud/initializers"
	"go-crud/models"
)

// init function is used to load the environment variables and connect to the DB before the application starts
func init() {
	initializers.LoadEnvVariables() // Load environment variables
	initializers.ConnectToDB() // Connect to the database

}

// main function is the entry point of the application
func main() {
	// AutoMigrate function is used to automatically create the table schema for the models. Here it will create a table schema for the Post model.
	initializers.DB.AutoMigrate(&models.Post{})
}
