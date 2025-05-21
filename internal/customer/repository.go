package customer

import "gorm.io/gorm"

type Repository interface {
	GetAll() ([]Customer, error)
	GetByID(id uint) (*Customer, error)
	Create(customer *Customer) error
	Update(id uint, customer *Customer) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]Customer, error) {
	var customers []Customer
	if err := r.db.Find(&customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}

func (r *repository) GetByID(id uint) (*Customer, error) {
	var customer Customer
	if err := r.db.First(&customer, id).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *repository) Create(customer *Customer) error {
	return r.db.Create(customer).Error
}

func (r *repository) Update(id uint, updated *Customer) error {
	var existing Customer
	if err := r.db.First(&existing, id).Error; err != nil {
		return err
	}
	updated.ID = id
	return r.db.Save(updated).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&Customer{}, id).Error
}
