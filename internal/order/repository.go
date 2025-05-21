package order

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetAll() ([]Order, error)
	GetByID(id uint) (*Order, error)
	Create(po *Order) error
	Update(po *Order) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetAll() ([]Order, error) {
	var pos []Order
	err := r.db.Preload("Details").Find(&pos).Error
	return pos, err
}

func (r *repository) GetByID(id uint) (*Order, error) {
	var po Order
	err := r.db.Preload("Details").First(&po, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &po, err
}

func (r *repository) Create(po *Order) error {
	return r.db.Create(po).Error
}

func (r *repository) Update(po *Order) error {
	return r.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(po).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&Order{}, id).Error
}
