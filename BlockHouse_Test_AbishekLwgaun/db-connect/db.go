package db_connect

import (
	"BlockHouse_Test_AbishekLwgaun/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

// ConnectDB initializes the SQLite database connection.
func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("orders.db"), &gorm.Config{})

	// Handling error while making the connection
	if err != nil {
		log.Fatal("Cannot make connection with the database: ", err)
	}

	// Perform database schema migration
	err = DB.AutoMigrate(&models.Order{})
	if err != nil {
		log.Fatal("Error migrating database schema: ", err)
	}

	log.Println("Database Connection is Successfully Established")
}
