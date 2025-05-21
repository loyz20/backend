package sale

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetAll() ([]Sale, error)
	GetByID(id uint) (*Sale, error)
	Create(so *Sale) error
	Update(so *Sale) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetAll() ([]Sale, error) {
	var sos []Sale
	err := r.db.Preload("Details").Find(&sos).Error
	return sos, err
}

func (r *repository) GetByID(id uint) (*Sale, error) {
	var so Sale
	err := r.db.Preload("Details").First(&so, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &so, err
}

func (r *repository) Create(so *Sale) error {
	return r.db.Create(so).Error
}

func (r *repository) Update(so *Sale) error {
	return r.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(so).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&Sale{}, id).Error
}
