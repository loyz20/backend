// internal/product/usecase.go
package product

import "errors"

type UseCase interface {
	GetAll() ([]Product, error)
	GetByID(id uint) (*Product, error)
	Create(product *Product) error
	Update(id uint, input *Product) (*Product, error)
	Delete(id uint) error
}

type usecase struct {
	repo Repository
}

// NewUseCase creates product business logic layer
func NewUseCase(repo Repository) UseCase {
	return &usecase{repo}
}

func (u *usecase) GetAll() ([]Product, error) {
	return u.repo.FindAll()
}

func (u *usecase) GetByID(id uint) (*Product, error) {
	return u.repo.FindByID(id)
}

func (u *usecase) Create(product *Product) error {
	if product.Name == "" || product.SKU == "" {
		return errors.New("name and sku are required")
	}
	return u.repo.Create(product)
}

func (u *usecase) Update(id uint, input *Product) (*Product, error) {
	existing, err := u.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, errors.New("product not found")
	}

	// Update fields
	existing.Name = input.Name
	existing.SKU = input.SKU
	existing.Category = input.Category
	existing.Unit = input.Unit
	existing.PurchasePrice = input.PurchasePrice
	existing.SellingPrice = input.SellingPrice
	existing.Stock = input.Stock
	existing.ExpiryDate = input.ExpiryDate
	existing.Batches = input.Batches

	err = u.repo.Update(existing)
	return existing, err
}

func (u *usecase) Delete(id uint) error {
	return u.repo.Delete(id)
}
