package initializers // Declaring the package name as initializers

import ( // Importing necessary packages
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB  // Declaring a global variable DB of type *gorm.DB

func ConnectToDB() {  // Function to connect to the database
	var err error  // Declaring an error variable
	dsn := os.Getenv("DB_URL")  // Getting the database URL from the environment variables
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})  // Opening a connection to the database using gorm and the postgres driver

	if err != nil {  // If there is an error in connecting to the database
		log.Fatal("Failed to connect to Database")  // Log the error and stop the program
	}
}
