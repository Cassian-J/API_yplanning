package dbmodel

import (
	"time"

	"gorm.io/gorm"
)

type Date struct {
	gorm.Model
	Title        string    `gorm:"not null" json:"title"`
	Body         string    `gorm:"not null" json:"body"`
	UserID       *uint     `json:"user_id"`
	User         *User     `gorm:"not null;constraint:OnDelete:CASCADE;"`
	BeginTime    time.Time `json:"begin_time"`
	EndTime      time.Time `json:"end_time"`
	Private      bool      `json:"private"`
	RecurrenceID *uint     `json:"recurrence_id"`
	Recurrence   *Date     `gorm:"constraint:OnDelete:SET NULL;"`
	ColorID      *uint     `json:"color_id"`
	Color        *Color    `gorm:"null;constraint:OnDelete:SET NULL;"`
}

type DateRepository interface {
	Create(date *Date) (*Date, error)
	FindAll() ([]Date, error)
	FindByID(id uint) (*Date, error)
	FindByUserID(userID uint) (*Date, error)
	FindByRecurrenceID(recurrenceID uint) (*Date, error)
	FindByDayRange(begin time.Time, end time.Time, userID uint) ([]Date, error)
	UpdateByID(id uint, date *Date) error
	DeleteByID(id uint) error
}

type dateRepository struct {
	DB *gorm.DB
}

func (dateRepository *dateRepository) Create(date *Date) (*Date, error) {
	if err := dateRepository.DB.Create(date).Error; err != nil {
		return nil, err
	}
	return date, nil
}

func (dateRepository *dateRepository) FindAll() ([]Date, error) {
	var dates []Date
	if err := dateRepository.DB.Find(&dates).Error; err != nil {
		return nil, err
	}
	return dates, nil
}

func (dateRepository *dateRepository) FindByID(id uint) (*Date, error) {
	var date Date
	if err := dateRepository.DB.First(&date, id).Error; err != nil {
		return nil, err
	}
	return &date, nil
}

func (dateRepository *dateRepository) FindByUserID(userID uint) (*Date, error) {
	var date Date
	if err := dateRepository.DB.Preload("User").First(&date, userID).Error; err != nil {
		return nil, err
	}
	return &date, nil
}

func (dateRepository *dateRepository) FindByRecurrenceID(recurrenceID uint) (*Date, error) {
	var date Date
	if err := dateRepository.DB.Preload("Recurrence").First(&date, recurrenceID).Error; err != nil {
		return nil, err
	}
	return &date, nil
}

func (dateRepository *dateRepository) FindByDayRange(begin time.Time, end time.Time, userID uint) ([]Date, error) {
	var dates []Date
	if err := dateRepository.DB.Preload("User").Where("begin_time >= ? AND end_time <= ? AND user_id = ?", begin, end, userID).Find(&dates).Error; err != nil {
		return nil, err
	}
	return dates, nil
}

func (dateRepository *dateRepository) UpdateByID(id uint, date *Date) error {
	if err := dateRepository.DB.Model(&Date{}).Where("id = ?", id).Updates(date).Error; err != nil {
		return err
	}
	return nil
}

func (dateRepository *dateRepository) DeleteByID(id uint) error {
	if err := dateRepository.DB.Delete(&Date{}, id).Error; err != nil {
		return err
	}
	return nil
}
