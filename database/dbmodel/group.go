package dbmodel

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name      string  `gorm:"not null" json:"name"`
	CreatorID uint    `json:"creator_id"`
	Creator   *User   `gorm:"not null;constraint:OnDelete:CASCADE;"`
	Users     []User  `gorm:"many2many:user_group;" json:"users"`
	Colors    []Color `gorm:"many2many:user_group;" json:"colors"`
}

type GroupRepository interface {
	Create(group *Group) (*Group, error)
	FindAll() ([]Group, error)
	FindByID(id uint) (*Group, error)
	FindByCreatorID(creatorID uint) (*Group, error)
	UpdateByID(id uint, group *Group) (*Group, error)
	DeleteByID(id uint) error
}

type groupRepository struct {
	DB *gorm.DB
}

func NewGroupRepository(db *gorm.DB) GroupRepository {
	return &groupRepository{DB: db}
}

func (groupRepository *groupRepository) Create(group *Group) (*Group, error) {
	if err := groupRepository.DB.Create(group).Error; err != nil {
		return nil, err
	}
	return group, nil
}

func (groupRepository *groupRepository) FindAll() ([]Group, error) {
	var groups []Group
	if err := groupRepository.DB.Find(&groups).Error; err != nil {
		return nil, err
	}
	return groups, nil
}

func (groupRepository *groupRepository) FindByID(id uint) (*Group, error) {
	var group Group
	if err := groupRepository.DB.First(&group, id).Error; err != nil {
		return nil, err
	}
	return &group, nil
}

func (groupRepository *groupRepository) FindByCreatorID(creatorID uint) (*Group, error) {
	var group Group
	if err := groupRepository.DB.Preload("Creator").First(&group, creatorID).Error; err != nil {
		return nil, err
	}
	return &group, nil
}

func (groupRepository *groupRepository) UpdateByID(id uint, group *Group) (*Group, error) {
	if err := groupRepository.DB.Model(&Group{}).Where("id = ?", id).Updates(group).Error; err != nil {
		return nil, err
	}
	return group, nil
}

func (groupRepository *groupRepository) DeleteByID(id uint) error {
	if err := groupRepository.DB.Delete(&Group{}, id).Error; err != nil {
		return err
	}
	return nil
}
