package database

import (
	"github.com/alexohneander/zeitungsfritze/model"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() {
	var err error

	DB, err = gorm.Open(sqlite.Open("zeitungsfritze.db"), &gorm.Config{})

	if err != nil {
		log.Panic().Msg("failed to connect database")
	}

	log.Info().Msg("Connection Opened to Database")
}

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&model.Forward{}, &model.User{})
	if err != nil {
		log.Error().Msg("Error migrating database: " + err.Error())
	}

	log.Info().Msg("Database migrated")
	return err
}
