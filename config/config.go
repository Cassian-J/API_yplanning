package config

import (
	"yplanning/database"
	"yplanning/database/dbmodel"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type config struct {
	GroupRepository        dbmodel.GroupRepository
	UserRepository         dbmodel.UserRepository
	ColorRepository        dbmodel.ColorRepository
	AvailabilityRepository dbmodel.AvailabilityRepository
	DateRepository         dbmodel.DateRepository
	UserGroupRepository    dbmodel.UserGroupRepository
}

func New() (*config, error) {
	config := &config{}

	databaseSession, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
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
