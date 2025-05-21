package product

import "errors"

type UseCase interface {
	GetAll() ([]Product, error)
	GetByID(id uint) (*Product, error)
	Create(product *Product) error
	Update(id uint, input *Product) (*Product, error)
	Delete(id uint) error

	// method batch
	GetBatchByID(productID uint) (*ProductBatch, error)
	GetBatchesByProductID(productID uint) ([]ProductBatch, error)
	CreateBatch(batch *ProductBatch) error
	UpdateBatch(batch *ProductBatch) error
	DeleteBatch(id uint) error
	DecreaseBatchStock(batchID uint, qty int) error
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
	existing.Batches = input.Batches

	err = u.repo.Update(existing)
	return existing, err
}

func (u *usecase) Delete(id uint) error {
	return u.repo.Delete(id)
}

// ==================== implementasi batch =======================

func (u *usecase) GetBatchByID(productID uint) (*ProductBatch, error) {
	return u.repo.FindBatchByID(productID)
}

func (u *usecase) GetBatchesByProductID(productID uint) ([]ProductBatch, error) {
	return u.repo.FindBatchesByProductID(productID)
}

func (u *usecase) CreateBatch(batch *ProductBatch) error {
	if batch.BatchNumber == "" {
		return errors.New("batch number is required")
	}
	return u.repo.CreateBatch(batch)
}

func (u *usecase) UpdateBatch(batch *ProductBatch) error {
	return u.repo.UpdateBatch(batch)
}

func (u *usecase) DeleteBatch(id uint) error {
	return u.repo.DeleteBatch(id)
}

func (u *usecase) DecreaseBatchStock(batchID uint, qty int) error {
	batch, err := u.repo.FindBatchByID(batchID)
	if err != nil {
		return err
	}
	if batch == nil {
		return errors.New("batch not found")
	}

	if batch.RemainingQuantity < qty {
		return errors.New("insufficient stock")
	}

	batch.Quantity -= qty

	return u.repo.UpdateBatch(batch)
}
