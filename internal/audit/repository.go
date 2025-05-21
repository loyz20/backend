package audit

import (
	"gorm.io/gorm"
)

type Repository interface {
	CreateLog(log *AuditLog) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) CreateLog(log *AuditLog) error {
	return r.db.Create(log).Error
}
