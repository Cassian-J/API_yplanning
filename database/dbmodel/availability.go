package dbmodel

import (
	"time"

	"gorm.io/gorm"
)

type Availability struct {
	gorm.Model
	UserID    *uint     `json:"user_id"`
	User      *User     `gorm:"not null;constraint:OnDelete:CASCADE;"`
	BeginTime time.Time `json:"begin_time"`
	EndTime   time.Time `json:"end_time"`
}

type AvailabilityRepository interface {
	Create(availability *Availability) (*Availability, error)
	FindAll() ([]Availability, error)
	FindByID(id uint) (*Availability, error)
	FindByUserID(userID uint) (*Availability, error)
	UpdateByID(id uint, availability *Availability) error
	DeleteByID(id uint) error
}

type availabilityRepository struct {
	DB *gorm.DB
}

func (availabilityRepository *availabilityRepository) Create(availability *Availability) (*Availability, error) {
	if err := availabilityRepository.DB.Create(availability).Error; err != nil {
		return nil, err
	}
	return availability, nil
}

func (availabilityRepository *availabilityRepository) FindAll() ([]Availability, error) {
	var availabilities []Availability
	if err := availabilityRepository.DB.Find(&availabilities).Error; err != nil {
		return nil, err
	}
	return availabilities, nil
}

func (availabilityRepository *availabilityRepository) FindByID(id uint) (*Availability, error) {
	var availability Availability
	if err := availabilityRepository.DB.First(&availability, id).Error; err != nil {
		return nil, err
	}
	return &availability, nil
}

func (availabilityRepository *availabilityRepository) FindByUserID(userID uint) (*Availability, error) {
	var availability Availability
	if err := availabilityRepository.DB.Preload("User").First(&availability, userID).Error; err != nil {
		return nil, err
	}
	return &availability, nil
}

func (availabilityRepository *availabilityRepository) UpdateByID(id uint, availability *Availability) error {
	if err := availabilityRepository.DB.Model(&Availability{}).Where("id = ?", id).Updates(availability).Error; err != nil {
		return err
	}
	return nil
}

func (availabilityRepository *availabilityRepository) DeleteByID(id uint) error {
	if err := availabilityRepository.DB.Delete(&Availability{}, id).Error; err != nil {
		return err
	}
	return nil
}
