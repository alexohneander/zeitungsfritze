package database

import (
	"github.com/alexohneander/zeitungsfritze/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("zeitungsfritze.db"), &gorm.Config{})
	return db, err
}

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&model.Forward{})

	return err
}
