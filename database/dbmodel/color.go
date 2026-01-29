package dbmodel

import "gorm.io/gorm"

type Color struct {
	gorm.Model
	HexCode string  `gorm:"uniqueIndex;not null;size:7" json:"hex_code"`
	Name    string  `gorm:"not null" json:"name"`
	Users   []User  `gorm:"many2many:user_group;" json:"users"`
	Groups  []Group `gorm:"many2many:user_group;" json:"groups"`
}

type ColorRepository interface {
	Create(color *Color) (*Color, error)
	FindAll() ([]Color, error)
	FindByID(id uint) (*Color, error)
	FindByHexCode(hexCode string) (*Color, error)
	UpdateByID(id uint, color *Color) error
	DeleteByID(id uint) error
}

type colorRepository struct {
	DB *gorm.DB
}

func NewColorRepository(db *gorm.DB) ColorRepository {
	return &colorRepository{DB: db}
}

func (colorRepository *colorRepository) Create(color *Color) (*Color, error) {
	if err := colorRepository.DB.Create(color).Error; err != nil {
		return nil, err
	}
	return color, nil
}

func (colorRepository *colorRepository) FindAll() ([]Color, error) {
	var colors []Color
	if err := colorRepository.DB.Find(&colors).Error; err != nil {
		return nil, err
	}
	return colors, nil
}

func (colorRepository *colorRepository) FindByID(id uint) (*Color, error) {
	var color Color
	if err := colorRepository.DB.First(&color, id).Error; err != nil {
		return nil, err
	}
	return &color, nil
}

func (colorRepository *colorRepository) FindByHexCode(hexCode string) (*Color, error) {
	var color Color
	if err := colorRepository.DB.Where("hex_code = ?", hexCode).First(&color).Error; err != nil {
		return nil, err
	}
	return &color, nil
}

func (colorRepository *colorRepository) UpdateByID(id uint, color *Color) error {
	if err := colorRepository.DB.Model(&Color{}).Where("id = ?", id).Updates(color).Error; err != nil {
		return err
	}
	return nil
}

func (colorRepository *colorRepository) DeleteByID(id uint) error {
	if err := colorRepository.DB.Delete(&Color{}, id).Error; err != nil {
		return err
	}
	return nil
}
