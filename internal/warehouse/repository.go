package warehouse

import "gorm.io/gorm"

type Repository interface {
	Create(w *Warehouse) error
	GetAll() ([]Warehouse, error)
	GetByID(id uint) (*Warehouse, error)
	Update(w *Warehouse) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(w *Warehouse) error {
	return r.db.Create(w).Error
}

func (r *repository) GetAll() ([]Warehouse, error) {
	var list []Warehouse
	err := r.db.Find(&list).Error
	return list, err
}

func (r *repository) GetByID(id uint) (*Warehouse, error) {
	var w Warehouse
	err := r.db.First(&w, id).Error
	if err != nil {
		return nil, err
	}
	return &w, nil
}

func (r *repository) Update(w *Warehouse) error {
	return r.db.Save(w).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&Warehouse{}, id).Error
}
