// internal/supplier/repository.go
package supplier

import "gorm.io/gorm"

type Repository interface {
	GetAll() ([]Supplier, error)
	GetByID(id uint) (*Supplier, error)
	Create(s *Supplier) error
	Update(s *Supplier) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

// NewRepository returns a new Supplier repository
func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetAll() ([]Supplier, error) {
	var list []Supplier
	err := r.db.Find(&list).Error
	return list, err
}

func (r *repository) GetByID(id uint) (*Supplier, error) {
	var s Supplier
	err := r.db.First(&s, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &s, err
}

func (r *repository) Create(s *Supplier) error {
	return r.db.Create(s).Error
}

func (r *repository) Update(s *Supplier) error {
	return r.db.Save(s).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&Supplier{}, id).Error
}
