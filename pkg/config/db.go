package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	var err error

	db, err := gorm.Open(postgres.Open("host=localhost user=root password=root dbname=gotbet-api port=5432"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database. Error: ", err)
	}

	return db
}
