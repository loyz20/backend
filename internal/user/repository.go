package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]User, error)
	FindByID(id uint) (*User, error)
	FindByEmail(email string) (*User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

// NewRepository returns a new user repository
func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *repository) FindByID(id uint) (*User, error) {
	var user User
	result := r.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *repository) FindByEmail(email string) (*User, error) {
	var user User
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *repository) Create(user *User) error {
	return r.db.Create(user).Error
}

func (r *repository) Update(user *User) error {
	return r.db.Save(user).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&User{}, id).Error
}
