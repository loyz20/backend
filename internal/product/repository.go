// internal/product/repository.go
package product

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Product, error)
	FindByID(id uint) (*Product, error)
	Create(product *Product) error
	Update(product *Product) error
	Delete(id uint) error

	FindBatchByID(id uint) (*ProductBatch, error)
	FindBatchesByProductID(productID uint) ([]ProductBatch, error)
	CreateBatch(batch *ProductBatch) error
	UpdateBatch(batch *ProductBatch) error
	DeleteBatch(id uint) error
}

type repository struct {
	db *gorm.DB
}

// NewRepository creates a new product repository
func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Product, error) {
	var products []Product
	err := r.db.Preload("Batches").Find(&products).Error
	return products, err
}

func (r *repository) FindByID(id uint) (*Product, error) {
	var product Product
	err := r.db.Preload("Batches").First(&product, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &product, err
}

func (r *repository) Create(product *Product) error {
	return r.db.Create(product).Error
}

func (r *repository) Update(product *Product) error {
	return r.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(product).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&Product{}, id).Error
}

// Batch CRUD
func (r *repository) FindBatchByID(id uint) (*ProductBatch, error) {
	var batch ProductBatch
	err := r.db.First(&batch, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &batch, err
}

func (r *repository) FindBatchesByProductID(productID uint) ([]ProductBatch, error) {
	var batches []ProductBatch
	err := r.db.Where("product_id = ?", productID).Find(&batches).Error
	return batches, err
}

func (r *repository) CreateBatch(batch *ProductBatch) error {
	return r.db.Create(batch).Error
}

func (r *repository) UpdateBatch(batch *ProductBatch) error {
	return r.db.Save(batch).Error
}

func (r *repository) DeleteBatch(id uint) error {
	return r.db.Delete(&ProductBatch{}, id).Error
}
