package dashboard

import (
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	FetchDashboardData() (*DashboardData, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) FetchDashboardData() (*DashboardData, error) {
	var orderCount int64
	r.db.Table("orders").Count(&orderCount)

	var productCount int64
	r.db.Table("products").Count(&productCount)

	var customerCount int64
	r.db.Table("customers").Count(&customerCount)

	// Ambil 5 pesanan terbaru, JOIN dengan customers untuk dapatkan nama
	var recent []OrderDTO
	rows, err := r.db.Table("orders").
		Select(`sales.id,
		        customers.name as customer_name,
		        sales.created_at,
		        sales.totalamount`).
		Joins(`LEFT JOIN customers 
		       ON customers.id = sales.customer_id`).
		Order("sales.created_at desc").
		Limit(5).
		Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id, customerName string
		var createdAt time.Time
		var totalPrice float64
		if err := rows.Scan(&id, &customerName, &createdAt, &totalPrice); err != nil {
			return nil, err
		}

		recent = append(recent, OrderDTO{
			ID:           id,
			CustomerName: customerName,
			Date:         createdAt.Format("2006-01-02"),
			Total:        totalPrice,
		})
	}

	return &DashboardData{
		TotalOrders:    int(orderCount),
		TotalProducts:  int(productCount),
		TotalCustomers: int(customerCount),
		RecentOrders:   recent,
	}, nil
}
