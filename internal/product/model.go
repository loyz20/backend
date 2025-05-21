package product

import (
	"time"
)

type Product struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	Name          string         `gorm:"not null" json:"name"`
	SKU           string         `gorm:"unique;not null" json:"sku"`
	Category      string         `json:"category"`
	Unit          string         `json:"unit"`
	PurchasePrice float64        `json:"purchase_price"`
	SellingPrice  float64        `json:"selling_price"`
	CreatedAt     int64          `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     int64          `gorm:"autoUpdateTime" json:"updated_at"`
	Batches       []ProductBatch `gorm:"foreignKey:ProductID" json:"batches,omitempty"`
}

type ProductBatch struct {
	ID                uint      `gorm:"primaryKey" json:"id"`
	ProductID         uint      `gorm:"index;not null" json:"product_id"`
	BatchNumber       string    `gorm:"not null" json:"batch_number"`
	Quantity          int       `json:"quantity"`           // total qty masuk batch ini
	RemainingQuantity int       `json:"remaining_quantity"` // stok sisa untuk batch ini
	ExpiryDate        time.Time `json:"expiry_date"`
	PurchaseOrderID   *uint     `gorm:"index" json:"purchase_order_id,omitempty"` // dari PO
	CreatedAt         int64     `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt         int64     `gorm:"autoUpdateTime" json:"updated_at"`
}
