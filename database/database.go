package database

import (
	"log"
	"yplanning/database/dbmodel"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase(db *gorm.DB) {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	Migrate(DB)

	log.Println("Database connected and migrated")
}

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&dbmodel.User{},
		&dbmodel.Availability{},
		&dbmodel.Color{},
		&dbmodel.Date{},
		&dbmodel.Group{},
		&dbmodel.UserGroup{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	err = db.SetupJoinTable(&dbmodel.User{}, "Groups", &dbmodel.UserGroup{})
	if err != nil {
		log.Fatal("Failed to setup join table:", err)
	}
	log.Println("Database migrated successfully")
}
