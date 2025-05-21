// internal/product/model.go
package product

import (
	"time"
)

// Product represents a pharmaceutical product
type Product struct {
	ID            uint       `gorm:"primaryKey" json:"id"`
	Name          string     `gorm:"size:100;not null" json:"name"`
	SKU           string     `gorm:"size:50;unique;not null" json:"sku"`
	Category      string     `gorm:"size:50" json:"category"`
	Unit          string     `gorm:"size:20" json:"unit"`
	PurchasePrice float64    `json:"purchase_price"`
	SellingPrice  float64    `json:"selling_price"`
	Stock         int        `json:"stock"`
	ExpiryDate    *time.Time `json:"expiry_date,omitempty"`
	Batches       []Batch    `gorm:"foreignKey:ProductID" json:"batches,omitempty"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

// Batch represents a batch/lot of a product
type Batch struct {
	ID         uint       `gorm:"primaryKey" json:"id"`
	ProductID  uint       `gorm:"index;not null" json:"product_id"`
	BatchCode  string     `gorm:"size:50;not null" json:"batch_code"`
	Quantity   int        `json:"quantity"`
	ExpiryDate *time.Time `json:"expiry_date,omitempty"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}
