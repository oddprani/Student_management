package database

import (
	"log"
	"studentmanager/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DBInstance *gorm.DB

func Initialize() {
	var err error

	DSN := config.ServerConfig.DataSourceName
	DBInstance, err = gorm.Open(sqlite.Open(DSN), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	// Migrate the schema
	DBInstance.AutoMigrate(&User{}, &Session{})

	log.Println("Database initialized")
}
