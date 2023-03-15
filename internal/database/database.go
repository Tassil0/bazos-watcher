package database

import (
	"github.com/Tassil0/bazos-watcher/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	db *gorm.DB
)

func Init() {
	dsn := "bazos:kokos@tcp(localhost:5420)/bazos?charset=utf8mb4&parseTime=True&loc=Local"

	dbOpen, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db = dbOpen
}

func GetQueries() (queries []models.Query, err error) {
	err = db.Find(&queries).Error
	return
}

func GetQueryItems(queryId uint32) {
}
