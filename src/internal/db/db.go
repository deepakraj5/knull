package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("knull.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return db
}
