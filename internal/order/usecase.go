package order

import (
	"errors"
)

type UseCase interface {
	GetAll() ([]Order, error)
	GetByID(id uint) (*Order, error)
	Create(po *Order) error
	Update(po *Order) error
	Delete(id uint) error
}

type usecase struct {
	repo Repository
}

func NewUseCase(repo Repository) UseCase {
	return &usecase{repo: repo}
}

func (u *usecase) GetAll() ([]Order, error) {
	return u.repo.GetAll()
}

func (u *usecase) GetByID(id uint) (*Order, error) {
	return u.repo.GetByID(id)
}

func (u *usecase) Create(po *Order) error {
	if po.OrderNumber == "" {
		return errors.New("order number is required")
	}
	if po.SupplierID == 0 {
		return errors.New("supplier id is required")
	}
	if len(po.Details) == 0 {
		return errors.New("order must have at least one detail")
	}

	// Calculate total amount
	var total float64
	for i := range po.Details {
		po.Details[i].Subtotal = float64(po.Details[i].Quantity) * po.Details[i].PurchasePrice
		total += po.Details[i].Subtotal
	}
	po.TotalAmount = total

	return u.repo.Create(po)
}

func (u *usecase) Update(po *Order) error {
	existing, err := u.repo.GetByID(po.ID)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("purchase order not found")
	}
	// Update total amount on details
	var total float64
	for i := range po.Details {
		po.Details[i].Subtotal = float64(po.Details[i].Quantity) * po.Details[i].PurchasePrice
		total += po.Details[i].Subtotal
	}
	po.TotalAmount = total

	return u.repo.Update(po)
}

func (u *usecase) Delete(id uint) error {
	return u.repo.Delete(id)
}
