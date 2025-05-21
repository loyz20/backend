package sale

import (
	"errors"
	"fmt"

	"backend/internal/product"
)

type UseCase interface {
	GetAll() ([]Sale, error)
	GetByID(id uint) (*Sale, error)
	Create(so *Sale) error
	Update(so *Sale) error
	Delete(id uint) error
}

type usecase struct {
	repo           Repository
	productUseCase product.UseCase
}

func NewUseCase(repo Repository, productUC product.UseCase) UseCase {
	return &usecase{repo: repo, productUseCase: productUC}
}

func (u *usecase) GetAll() ([]Sale, error) {
	return u.repo.GetAll()
}

func (u *usecase) GetByID(id uint) (*Sale, error) {
	return u.repo.GetByID(id)
}

func (u *usecase) Create(so *Sale) error {
	if so.OrderNumber == "" {
		return errors.New("order number is required")
	}
	if so.CustomerID == 0 {
		return errors.New("customer id is required")
	}
	if len(so.Details) == 0 {
		return errors.New("order must have at least one detail")
	}

	// Validate batch stock and calculate total
	var total float64
	for i := range so.Details {
		// Check if batch stock available
		batch, err := u.productUseCase.GetBatchByID(so.Details[i].BatchID)
		if err != nil {
			return fmt.Errorf("failed to get batch: %w", err)
		}
		if batch == nil {
			return fmt.Errorf("batch %d not found", so.Details[i].BatchID)
		}
		if batch.Quantity < so.Details[i].Quantity {
			return fmt.Errorf("insufficient stock for batch %d", so.Details[i].BatchID)
		}

		so.Details[i].Subtotal = float64(so.Details[i].Quantity) * so.Details[i].SellingPrice
		total += so.Details[i].Subtotal
	}

	so.TotalAmount = total

	// Create SO first
	err := u.repo.Create(so)
	if err != nil {
		return err
	}

	// Reduce stock from batch
	for _, detail := range so.Details {
		err := u.productUseCase.DecreaseBatchStock(detail.BatchID, detail.Quantity)
		if err != nil {
			return fmt.Errorf("failed to decrease stock: %w", err)
		}
	}

	return nil
}

func (u *usecase) Update(so *Sale) error {
	// For simplicity, no stock adjustment on update
	return u.repo.Update(so)
}

func (u *usecase) Delete(id uint) error {
	// Optionally, restore stock here if SO is deleted
	return u.repo.Delete(id)
}
