package config

import (
	"fmt"
	"log"
	"os"
	"yplanning/database"
	"yplanning/database/dbmodel"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Constants struct {
	Port             string `yaml:"port"`
	JWTSecret        string `yaml:"jwt_secret"`
	ConnectionString string `yaml:"connection_string"`
	ADDRESS_IP       string `yaml:"address_ip"`
}

type Config struct {
	Constants

	//repositories
	GroupRepository        dbmodel.GroupRepository
	UserRepository         dbmodel.UserRepository
	ColorRepository        dbmodel.ColorRepository
	AvailabilityRepository dbmodel.AvailabilityRepository
	DateRepository         dbmodel.DateRepository
	UserGroupRepository    dbmodel.UserGroupRepository
}

func initEnv(fileName string) (Constants, error) {
	if _, err := os.Stat(fileName); err == nil {
		if err := godotenv.Load(fileName); err != nil {
			log.Println("warning: failed to load .env:", err)
		}
	} else if !os.IsNotExist(err) {
		log.Println("warning: error checking .env:", err)
	}

	var constants Constants

	constants.Port = os.Getenv("PORT")
	constants.JWTSecret = os.Getenv("JWT_SECRET_KEY")
	constants.ConnectionString = os.Getenv("CONNECTION_STRING")

	if constants.Port == "" {
		return Constants{}, fmt.Errorf("missing required env var PORT")
	}
	if constants.JWTSecret == "" {
		return Constants{}, fmt.Errorf("missing required env var JWT_SECRET_KEY")
	}
	if constants.ConnectionString == "" {
		return Constants{}, fmt.Errorf("missing required env var CONNECTION_STRING")
	}

	return constants, nil
}

func New() (*Config, error) {
	config := Config{}

	// Constants
	constants, err := initEnv(".env")

	config.Constants = constants
	if err != nil {
		return &config, err
	}

	databaseSession, err := gorm.Open(postgres.Open(config.ConnectionString), &gorm.Config{})
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
	return &config, nil
}
