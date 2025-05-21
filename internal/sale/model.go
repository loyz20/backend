package sale

import (
	"time"
)

type Sale struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	OrderNumber string    `json:"order_number" gorm:"unique;not null"`
	CustomerID  uint      `json:"customer_id" gorm:"not null"`
	OrderDate   time.Time `json:"order_date" gorm:"not null"`
	Status      string    `json:"status" gorm:"not null"` // e.g. pending, shipped, cancelled
	TotalAmount float64   `json:"total_amount"`

	Details []SaleDetail `json:"details" gorm:"foreignKey:SaleOrderID;constraint:OnDelete:CASCADE"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SaleDetail struct {
	ID           uint    `json:"id" gorm:"primaryKey"`
	SaleOrderID  uint    `json:"sale_order_id" gorm:"index"`
	ProductID    uint    `json:"product_id" gorm:"not null"`
	BatchID      uint    `json:"batch_id"` // wajib pilih batch yang stoknya cukup
	Quantity     int     `json:"quantity" gorm:"not null"`
	SellingPrice float64 `json:"selling_price" gorm:"not null"`
	Subtotal     float64 `json:"subtotal"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
