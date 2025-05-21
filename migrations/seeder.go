// migrations/seeder.go
package migrations

import (
	"time"

	"gorm.io/gorm"

	"backend/internal/customer"
	"backend/internal/order"
	"backend/internal/product"
	"backend/internal/sale"
	"backend/internal/supplier"
	"backend/internal/user"
	"backend/internal/warehouse"
)

// Seed populates the database with initial data
func Seed(db *gorm.DB) error {
	// Users
	users := []user.User{
		{Name: "Admin", Email: "admin@example.com", PasswordHash: "$2a$10$hash", Role: "admin"},
		{Name: "Staff", Email: "staff@example.com", PasswordHash: "$2a$10$hash2", Role: "staff"},
	}
	for _, u := range users {
		db.FirstOrCreate(&u, user.User{Email: u.Email})
	}

	// Customers
	customers := []customer.Customer{
		{Name: "Apotek Sehat", Phone: "0812345678", Email: "apotek@sehat.com", Address: "Jakarta", SIPA: "SIPA123", ExpiredDate: parseTimePtr("2026-12-31T00:00:00Z")},
	}
	for _, c := range customers {
		db.FirstOrCreate(&c, customer.Customer{Email: c.Email})
	}

	// Suppliers
	suppliers := []supplier.Supplier{
		{Name: "PT Farma", Phone: "0823456789", Email: "contact@farma.com", Address: "Bandung"},
	}
	for _, s := range suppliers {
		db.FirstOrCreate(&s, supplier.Supplier{Name: s.Name})
	}

	// Warehouses
	warehouses := []warehouse.Warehouse{
		{Name: "Gudang Utama", Location: "Jakarta", Manager: "Budi"},
	}
	for _, w := range warehouses {
		db.FirstOrCreate(&w, warehouse.Warehouse{Name: w.Name})
	}

	// Products and Batches
	products := []product.Product{
		{
			Name:          "Paracetamol 500mg",
			SKU:           "PRC500",
			Category:      "Obat",
			Unit:          "Tablet",
			PurchasePrice: 1000,
			SellingPrice:  1500,
			Batches: []product.ProductBatch{
				{BatchNumber: "BATCH-A1", Quantity: 1000, RemainingQuantity: 1000, ExpiryDate: time.Now().AddDate(1, 0, 0)},
			},
		},
	}
	for _, p := range products {
		db.FirstOrCreate(&p, product.Product{SKU: p.SKU})
		for _, b := range p.Batches {
			b.ProductID = p.ID
			db.FirstOrCreate(&b, product.ProductBatch{BatchNumber: b.BatchNumber})
		}
	}

	// Purchase Orders
	po := order.Order{
		OrderNumber: "PO-001",
		SupplierID:  suppliers[0].ID,
		OrderDate:   time.Now(),
		Status:      "received",
		Details: []order.OrderDetail{
			{ProductID: products[0].ID, Quantity: 500, PurchasePrice: 1000},
		},
	}
	db.FirstOrCreate(&po, order.Order{OrderNumber: po.OrderNumber})

	// Sale Orders
	so := sale.Sale{
		OrderNumber: "SO-001",
		CustomerID:  customers[0].ID,
		OrderDate:   time.Now(),
		Status:      "pending",
		Details: []sale.SaleDetail{
			{ProductID: products[0].ID, BatchID: products[0].Batches[0].ID, Quantity: 100, SellingPrice: 1500},
		},
	}
	db.FirstOrCreate(&so, sale.Sale{OrderNumber: so.OrderNumber})

	return nil
}

// parseTimePtr is helper to parse RFC3339 string to *time.Time
func parseTimePtr(value string) *time.Time {
	t, _ := time.Parse(time.RFC3339, value)
	return &t
}
