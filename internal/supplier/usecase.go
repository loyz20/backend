// internal/supplier/usecase.go
package supplier

import "errors"

type UseCase interface {
	GetAll() ([]Supplier, error)
	GetByID(id uint) (*Supplier, error)
	Create(s *Supplier) error
	Update(id uint, s *Supplier) (*Supplier, error)
	Delete(id uint) error
}

type usecase struct {
	repo Repository
}

// NewUseCase returns a new Supplier use case
func NewUseCase(repo Repository) UseCase {
	return &usecase{repo: repo}
}

func (u *usecase) GetAll() ([]Supplier, error) {
	return u.repo.GetAll()
}

func (u *usecase) GetByID(id uint) (*Supplier, error) {
	return u.repo.GetByID(id)
}

func (u *usecase) Create(s *Supplier) error {
	if s.Name == "" {
		return errors.New("supplier name is required")
	}
	return u.repo.Create(s)
}

func (u *usecase) Update(id uint, input *Supplier) (*Supplier, error) {
	existing, err := u.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, errors.New("supplier not found")
	}

	existing.Name = input.Name
	existing.Phone = input.Phone
	existing.Email = input.Email
	existing.Address = input.Address

	err = u.repo.Update(existing)
	return existing, err
}

func (u *usecase) Delete(id uint) error {
	return u.repo.Delete(id)
}
