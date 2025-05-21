package warehouse

import "errors"

type UseCase interface {
	Create(w *Warehouse) error
	GetAll() ([]Warehouse, error)
	GetByID(id uint) (*Warehouse, error)
	Update(w *Warehouse) (*Warehouse, error)
	Delete(id uint) error
}

type useCase struct {
	repo Repository
}

func NewUseCase(repo Repository) UseCase {
	return &useCase{repo}
}

func (u *useCase) Create(w *Warehouse) error {
	if w.Name == "" {
		return errors.New("name are required")
	}
	return u.repo.Create(w)
}

func (u *useCase) GetAll() ([]Warehouse, error) {
	return u.repo.GetAll()
}

func (u *useCase) GetByID(id uint) (*Warehouse, error) {
	return u.repo.GetByID(id)
}

func (u *useCase) Update(w *Warehouse) (*Warehouse, error) {
	existing, err := u.repo.GetByID(w.ID)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, errors.New("warehouse not found")
	}
	existing.Name = w.Name
	existing.Location = w.Location
	existing.Manager = w.Manager
	if err := u.repo.Update(existing); err != nil {
		return nil, err
	}
	return existing, nil
}

func (u *useCase) Delete(id uint) error {
	return u.repo.Delete(id)
}
