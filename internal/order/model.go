package order

//
import (
	"time"
)

type Order struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	OrderNumber string    `json:"order_number" gorm:"unique;not null"`
	SupplierID  uint      `json:"supplier_id" gorm:"not null"`
	OrderDate   time.Time `json:"order_date" gorm:"not null"`
	Status      string    `json:"status" gorm:"not null"` // e.g. pending, received, cancelled
	TotalAmount float64   `json:"total_amount"`

	Details   []OrderDetail `json:"details" gorm:"foreignKey:PurchaseOrderID;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

type OrderDetail struct {
	ID              uint    `json:"id" gorm:"primaryKey"`
	PurchaseOrderID uint    `json:"purchase_order_id" gorm:"index"`
	ProductID       uint    `json:"product_id" gorm:"not null"`
	BatchNumber     string  `json:"batch_number"` // Optional: batch nomor jika sudah tahu
	Quantity        int     `json:"quantity" gorm:"not null"`
	PurchasePrice   float64 `json:"purchase_price" gorm:"not null"`
	Subtotal        float64 `json:"subtotal"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
