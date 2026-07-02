package dbmodel

import (
	"gorm.io/gorm"
)

type Availability struct {
	gorm.Model
	UserID    uint      `json:"user_id"`
	User      *User     `gorm:"not null;constraint:OnDelete:CASCADE;"`
	BeginTime int64       `json:"begin_time"`
	EndTime   int64       `json:"end_time"`
}

type AvailabilityRepository interface {
	Create(availability *Availability) (*Availability, error)
	FindAll() ([]Availability, error)
	FindByID(id uint) (*Availability, error)
	FindByUserID(userID uint) ([]Availability, error)
	UpdateByID(id uint, availability *Availability) (*Availability, error)
	DeleteByID(id uint) error
}

type availabilityRepository struct {
	DB *gorm.DB
}

func NewAvailabilityRepository(db *gorm.DB) AvailabilityRepository {
	return &availabilityRepository{DB: db}
}

func (availabilityRepository *availabilityRepository) Create(availability *Availability) (*Availability, error) {
	if err := availabilityRepository.DB.Create(availability).Error; err != nil {
		return nil, err
	}

	if err := availabilityRepository.DB.
		Preload("User").
		First(availability, availability.ID).Error; err != nil {
		return nil, err
	}

	return availability, nil
}

func (availabilityRepository *availabilityRepository) FindAll() ([]Availability, error) {
	var availabilities []Availability

	if err := availabilityRepository.DB.
		Preload("User").
		Find(&availabilities).Error; err != nil {
		return nil, err
	}

	return availabilities, nil
}

func (availabilityRepository *availabilityRepository) FindByID(id uint) (*Availability, error) {
	var availability Availability

	if err := availabilityRepository.DB.
		Preload("User").
		First(&availability, id).Error; err != nil {
		return nil, err
	}

	return &availability, nil
}

func (availabilityRepository *availabilityRepository) FindByUserID(userID uint) ([]Availability, error) {
	var availabilities []Availability

	if err := availabilityRepository.DB.
		Preload("User").
		Where("user_id = ?", userID).
		Find(&availabilities).Error; err != nil {
		return nil, err
	}

	return availabilities, nil
}

func (availabilityRepository *availabilityRepository) UpdateByID(id uint, availability *Availability) (*Availability, error) {
	if err := availabilityRepository.DB.
		Model(&Availability{}).
		Where("id = ?", id).
		Updates(availability).Error; err != nil {
		return nil, err
	}

	var updated Availability
	if err := availabilityRepository.DB.
		Preload("User").
		First(&updated, id).Error; err != nil {
		return nil, err
	}

	return &updated, nil
}

func (availabilityRepository *availabilityRepository) DeleteByID(id uint) error {
	if err := availabilityRepository.DB.Delete(&Availability{}, id).Error; err != nil {
		return err
	}
	return nil
}
