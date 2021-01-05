package db

import (
	"os"

	"ema_generator/entity"

	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
	err error
)

func InitDatabase() {
	env := os.Getenv("ENV")

	if env == "production" {
		env = "production"
	} else {
		env = "development"
	}

	db, err = gorm.Open("mysql", os.Getenv("CONNECT"))

	if err != nil {
		panic(err)
	}

	autoMigration()
}

func autoMigration() {
	db.AutoMigrate(&entity.Ema{})
}

func GetDatabase() *gorm.DB {
	return db
}

func CloseDatabase() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}
