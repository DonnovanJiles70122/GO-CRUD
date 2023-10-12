// Package models contains the data models or structures used in the application
package models

// Importing the gorm library. Gorm is a developer-friendly ORM library for handling database operations in Go.
import "gorm.io/gorm"

// Defining a struct for Post. This will be used to map the post data in the database.
type Post struct {
	gorm.Model // This is a basic GoLang struct which includes fields ID, CreatedAt, UpdatedAt, DeletedAt.
	Title string // Title of the post
	Body string  // Body of the post
}
