package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	db *gorm.DB
)

func Init() {
	dsn := "bazos:kokos@tcp(127.0.0.1:5420)/bazos?charset=utf8mb4&parseTime=True&loc=Local"

	dbOpen, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db = dbOpen
}

func GetQueries() {
	db.Find()
}

func GetQueryItems(queryId uint32) {
}
