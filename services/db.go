package services

import (
	"github.com/jinzhu/gorm"
	"github.com/auth-web-tokens/services/config"
	_ "github.com/lib/pq"
	"log"
)

var DB *gorm.DB

func InitDB() error {
	databaseString := config.Config.Database

	if db, err := open(databaseString); err == nil {
		DB = db
	} else {
		return err
	}
	return nil
}

func open(connectionString string) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", connectionString)

	if err != nil {
		return nil, err
	}

	db.SingularTable(true)
	if config.LogFile != nil {
		db.SetLogger(log.New(config.LogFile, "\r\n", 0))
	}
	return db, nil
}
