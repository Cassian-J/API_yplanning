package dbmodel

import "gorm.io/gorm"

type UserGroup struct {
	UserID  uint `json:"user_id"`
	GroupID uint `json:"group_id"`
	ColorID uint `gorm:"null" json:"color_id"`
}

type UserGroupRepository interface {
	Create(userGroup *UserGroup) (*UserGroup, error)
	FindAll() ([]UserGroup, error)
	FindByUserID(userID uint) ([]UserGroup, error)
	FindByGroupID(groupID uint) ([]UserGroup, error)
	FindByUserIDAndGroupID(userID uint, groupID uint) (*UserGroup, error)
	UpdateColorByUserIDAndGroupID(userID uint, groupID uint, colorID uint) error
	DeleteByUserIDAndGroupID(userID uint, groupID uint) error
	DeleteByGroupID(groupID uint) error
}

type userGroupRepository struct {
	DB *gorm.DB
}

func NewUserGroupRepository(db *gorm.DB) UserGroupRepository {
	return &userGroupRepository{DB: db}
}

func (userGroupRepository *userGroupRepository) Create(userGroup *UserGroup) (*UserGroup, error) {
	if err := userGroupRepository.DB.Create(userGroup).Error; err != nil {
		return nil, err
	}
	return userGroup, nil
}

func (userGroupRepository *userGroupRepository) FindAll() ([]UserGroup, error) {
	var userGroups []UserGroup
	if err := userGroupRepository.DB.Find(&userGroups).Error; err != nil {
		return nil, err
	}
	return userGroups, nil
}

func (userGroupRepository *userGroupRepository) FindByUserID(userID uint) ([]UserGroup, error) {
	var userGroups []UserGroup
	if err := userGroupRepository.DB.Preload("User").Where("user_id = ?", userID).Find(&userGroups).Error; err != nil {
		return nil, err
	}
	return userGroups, nil
}

func (userGroupRepository *userGroupRepository) FindByGroupID(groupID uint) ([]UserGroup, error) {
	var userGroups []UserGroup
	if err := userGroupRepository.DB.Preload("Group").Where("group_id = ?", groupID).Find(&userGroups).Error; err != nil {
		return nil, err
	}
	return userGroups, nil
}

func (userGroupRepository *userGroupRepository) FindByUserIDAndGroupID(userID uint, groupID uint) (*UserGroup, error) {
	var userGroup UserGroup
	if err := userGroupRepository.DB.Preload("User").Preload("Group").Where("user_id = ? AND group_id = ?", userID, groupID).First(&userGroup).Error; err != nil {
		return nil, err
	}
	return &userGroup, nil
}

func (userGroupRepository *userGroupRepository) UpdateColorByUserIDAndGroupID(userID uint, groupID uint, colorID uint) error {
	if err := userGroupRepository.DB.Model(&UserGroup{}).Where("user_id = ? AND group_id = ?", userID, groupID).Update("color_id", colorID).Error; err != nil {
		return err
	}
	return nil
}

func (userGroupRepository *userGroupRepository) DeleteByUserIDAndGroupID(userID uint, groupID uint) error {
	if err := userGroupRepository.DB.Where("user_id = ? AND group_id = ?", userID, groupID).Delete(&UserGroup{}).Error; err != nil {
		return err
	}
	return nil
}

func (userGroupRepository *userGroupRepository) DeleteByGroupID(groupID uint) error {
	if err := userGroupRepository.DB.Where("group_id = ?", groupID).Delete(&UserGroup{}).Error; err != nil {
		return err
	}
	return nil
}
