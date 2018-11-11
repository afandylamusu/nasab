package db

import (
	"log"

	"github.com/jinzhu/gorm"

	// include postgres dialects
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Connect to database
func Connect() *gorm.DB {
	log.Println("Connecting to SQL Db...")

	conn, err := gorm.Open("postgres", "host=myhost port=myport user=gorm dbname=gorm password=mypassword")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	conn.AutoMigrate(&User{})

	conn.AutoMigrate(&People{})

	return conn
}
