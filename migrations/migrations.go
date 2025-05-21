package migrations

import (
	"gorm.io/gorm"

	"backend/internal/customer"
	"backend/internal/order"
	"backend/internal/product"
	"backend/internal/sale"
	"backend/internal/supplier"
	"backend/internal/user"
	"backend/internal/warehouse"
)

// RunAll executes all migrations
func RunAll(db *gorm.DB) error {
	return db.AutoMigrate(
		&user.User{},
		&customer.Customer{},
		&supplier.Supplier{},
		&warehouse.Warehouse{},
		&product.Product{},
		&product.ProductBatch{},
		&order.Order{},
		&order.OrderDetail{},
		&sale.Sale{},
		&sale.SaleDetail{},
	)
}
