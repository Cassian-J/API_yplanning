package config

import (
	"os"
	"yplanning/database"
	"yplanning/database/dbmodel"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	GroupRepository        dbmodel.GroupRepository
	UserRepository         dbmodel.UserRepository
	ColorRepository        dbmodel.ColorRepository
	AvailabilityRepository dbmodel.AvailabilityRepository
	DateRepository         dbmodel.DateRepository
	UserGroupRepository    dbmodel.UserGroupRepository
}

func New() (*Config, error) {
	config := &Config{}

	databaseSession, err := gorm.Open(postgres.Open(os.Getenv("CONNECTION_STRING")), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	database.Migrate(databaseSession)

	config.GroupRepository = dbmodel.NewGroupRepository(databaseSession)
	config.UserRepository = dbmodel.NewUserRepository(databaseSession)
	config.ColorRepository = dbmodel.NewColorRepository(databaseSession)
	config.AvailabilityRepository = dbmodel.NewAvailabilityRepository(databaseSession)
	config.DateRepository = dbmodel.NewDateRepository(databaseSession)
	config.UserGroupRepository = dbmodel.NewUserGroupRepository(databaseSession)
	return config, nil
}
