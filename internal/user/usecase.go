package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UseCase interface {
	Register(user *User, password string) error
	Login(email, password string) (*User, error)
	GetAll() ([]User, error)
	GetByID(id uint) (*User, error)
	Update(id uint, input *User) error
	Delete(id uint) error
}

type usecase struct {
	repo Repository
}

// NewUseCase returns user usecase
func NewUseCase(repo Repository) UseCase {
	return &usecase{repo}
}

func (u *usecase) Register(user *User, password string) error {
	// Check existing
	_, err := u.repo.FindByEmail(user.Email)
	if err == nil {
		return errors.New("email already registered")
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hash)

	return u.repo.Create(user)
}

func (u *usecase) Login(email, password string) (*User, error) {
	user, err := u.repo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}
	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, errors.New("invalid email or password")
	}
	return user, nil
}

func (u *usecase) GetAll() ([]User, error) {
	return u.repo.FindAll()
}

func (u *usecase) GetByID(id uint) (*User, error) {
	return u.repo.FindByID(id)
}

func (u *usecase) Update(id uint, input *User) error {
	user, err := u.repo.FindByID(id)
	if err != nil {
		return err
	}
	// Update fields
	user.Name = input.Name
	user.Role = input.Role
	return u.repo.Update(user)
}

func (u *usecase) Delete(id uint) error {
	return u.repo.Delete(id)
}
