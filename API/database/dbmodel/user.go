package dbmodel

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex" json:"username"`
	Email    string `gorm:"uniqueIndex;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	ColorID  *uint  `json:"color_id,omitempty"`
	Color    *Color `gorm:"constraint:OnDelete:SET NULL;"`

	UserGroups []UserGroup `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"user_groups"`
}

type UserRepository interface {
	Create(user *User) (*User, error)
	FindAll() ([]User, error)
	FindByID(id uint) (*User, error)
	FindByEmail(email string) (*User, error)
	FindByUsername(username string) (*User, error)
	UpdateByID(id uint, user *User) (*User, error)
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

	if err := userRepository.DB.Preload("Color").First(user, user.ID).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (userRepository *userRepository) FindAll() ([]User, error) {
	var users []User

	if err := userRepository.DB.
		Preload("Color").
		Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (userRepository *userRepository) FindByID(id uint) (*User, error) {
	var user User

	if err := userRepository.DB.
		Preload("Color").
		First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (userRepository *userRepository) FindByEmail(email string) (*User, error) {
	var user User

	if err := userRepository.DB.
		Preload("Color").
		Where("email = ?", email).
		First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (userRepository *userRepository) FindByUsername(username string) (*User, error) {
	var user User

	if err := userRepository.DB.
		Preload("Color").
		Where("username = ?", username).
		First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (userRepository *userRepository) UpdateByID(id uint, user *User) (*User, error) {
	if err := userRepository.DB.
		Model(&User{}).
		Where("id = ?", id).
		Updates(user).Error; err != nil {
		return nil, err
	}

	var updatedUser User
	if err := userRepository.DB.
		Preload("Color").
		First(&updatedUser, id).Error; err != nil {
		return nil, err
	}

	return &updatedUser, nil
}

func (userRepository *userRepository) DeleteByID(id uint) error {
	if err := userRepository.DB.Delete(&User{}, id).Error; err != nil {
		return err
	}
	return nil
}
