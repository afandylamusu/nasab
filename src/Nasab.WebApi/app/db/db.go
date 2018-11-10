package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Connect to database
func Connect() *gorm.DB {
	log.Println("Connecting to SQL Db...")

	db, err := gorm.Open("postgres", "host=myhost port=myport user=gorm dbname=gorm password=mypassword")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	db.AutoMigrate(&User{})

	db.AutoMigrate(&People{})

	return db
}
