package dbmodel

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string  `gorm:"uniqueIndex;not null" json:"username"`
	Email    string  `gorm:"uniqueIndex;not null" json:"email"`
	Password string  `gorm:"not null" json:"password"`
	Name     string  `json:"name"`
	Surname  string  `json:"surname"`
	ColorID  *uint   `json:"color_id"`
	Color    *Color  `gorm:"null;constraint:OnDelete:SET NULL;"`
	Groups   []Group `gorm:"many2many:user_group;" json:"groups"`
	Colors   []Color `gorm:"many2many:user_group;" json:"colors"`
}

type UserRepository interface {
	Create(user *User) (*User, error)
	FindAll() ([]User, error)
	FindByID(id uint) (*User, error)
	FindByEmail(email string) (*User, error)
	FindByUsername(username string) (*User, error)
	UpdateByID(id uint, user *User) error
	DeleteByID(id uint) error
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{DB: db}
}

func (userRepository *userRepository) Create(user *User) (*User, error) {
	if err := userRepository.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (userRepository *userRepository) FindAll() ([]User, error) {
	var users []User
	if err := userRepository.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (userRepository *userRepository) FindByID(id uint) (*User, error) {
	var user User
	if err := userRepository.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (userRepository *userRepository) FindByEmail(email string) (*User, error) {
	var user User
	if err := userRepository.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (userRepository *userRepository) FindByUsername(username string) (*User, error) {
	var user User
	if err := userRepository.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (userRepository *userRepository) UpdateByID(id uint, user *User) error {
	if err := userRepository.DB.Model(&User{}).Where("id = ?", id).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func (userRepository *userRepository) DeleteByID(id uint) error {
	if err := userRepository.DB.Delete(&User{}, id).Error; err != nil {
		return err
	}
	return nil
}
