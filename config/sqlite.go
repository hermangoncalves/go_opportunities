package config

import (
	"fmt"
	"os"

	"github.com/hermangoncalves/go_opportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	dbPath := "./db/main.db"
	
	// Check if the database file exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		fmt.Println("database file not found, creating...")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}


	// Create DB and connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&schemas.Opening{})

	// return the DB
	return db, nil
}