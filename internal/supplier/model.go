// internal/supplier/model.go
package supplier

import (
	"time"

	"gorm.io/gorm"
)

// Supplier represents a vendor supplying products
type Supplier struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"size:100;not null"`
	Phone     string         `json:"phone" gorm:"size:20"`
	Email     string         `json:"email" gorm:"size:100"`
	Address   string         `json:"address" gorm:"size:255"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
