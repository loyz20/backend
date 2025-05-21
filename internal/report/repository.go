// internal/report/repository.go
package report

import (
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	GetStockReport(startDate, endDate time.Time, sku, name string) ([]StockReportRow, error)
	GetPurchaseReport(startDate, endDate time.Time, supplierName string) ([]PurchaseReportRow, error)
	GetSaleReport(startDate, endDate time.Time, customerName string) ([]SaleReportRow, error)
	GetFinancialReport(startDate, endDate time.Time) (*FinancialReport, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) GetStockReport(startDate, endDate time.Time, sku, name string) ([]StockReportRow, error) {
	var results []StockReportRow

	query := r.db.Table("products").
		Select(`
			products.id as product_id,
			products.name,
			products.sku,
			COUNT(product_batches.id) as total_batch,
			COALESCE(SUM(product_batches.remaining_quantity), 0) as total_stock,
			COALESCE(SUM(product_batches.remaining_quantity * products.purchase_price), 0) as total_value,
			MAX(product_batches.expiry_date) as latest_batch_expiry
		`).
		Joins("LEFT JOIN product_batches ON product_batches.product_id = products.id").
		Group("products.id")

	if !startDate.IsZero() {
		query = query.Where("product_batches.created_at >= ?", startDate)
	}
	if !endDate.IsZero() {
		query = query.Where("product_batches.created_at <= ?", endDate)
	}
	if sku != "" {
		query = query.Where("products.sku ILIKE ?", "%"+sku+"%")
	}
	if name != "" {
		query = query.Where("products.name ILIKE ?", "%"+name+"%")
	}

	err := query.Scan(&results).Error
	return results, err
}

func (r *repository) GetPurchaseReport(startDate, endDate time.Time, supplierName string) ([]PurchaseReportRow, error) {
	var results []PurchaseReportRow

	query := r.db.Table("orders").
		Select("orders.id as purchase_order_id, orders.order_number, suppliers.name as supplier_name, orders.order_date, orders.total_amount").
		Joins("JOIN suppliers ON orders.supplier_id = suppliers.id")

	if !startDate.IsZero() {
		query = query.Where("orders.order_date >= ?", startDate)
	}
	if !endDate.IsZero() {
		query = query.Where("orders.order_date <= ?", endDate)
	}
	if supplierName != "" {
		query = query.Where("suppliers.name ILIKE ?", "%"+supplierName+"%")
	}

	err := query.Order("orders.order_date DESC").Scan(&results).Error
	return results, err
}

func (r *repository) GetSaleReport(startDate, endDate time.Time, customerName string) ([]SaleReportRow, error) {
	var results []SaleReportRow

	query := r.db.Table("sales").
		Select("sales.id as sale_order_id, sales.order_number, customers.name as customer_name, sales.order_date, sales.total_amount").
		Joins("JOIN customers ON sales.customer_id = customers.id")

	if !startDate.IsZero() {
		query = query.Where("sales.order_date >= ?", startDate)
	}
	if !endDate.IsZero() {
		query = query.Where("sales.order_date <= ?", endDate)
	}
	if customerName != "" {
		query = query.Where("customers.name ILIKE ?", "%"+customerName+"%")
	}

	err := query.Order("sales.order_date DESC").Scan(&results).Error
	return results, err
}

func (r *repository) GetFinancialReport(startDate, endDate time.Time) (*FinancialReport, error) {
	var revenue float64
	var expense float64

	// Hitung total penjualan (revenue)
	err := r.db.Table("sales").
		Select("COALESCE(SUM(total_amount),0)").
		Where("order_date BETWEEN ? AND ?", startDate, endDate).
		Scan(&revenue).Error
	if err != nil {
		return nil, err
	}

	// Hitung total pembelian (expense)
	err = r.db.Table("orders").
		Select("COALESCE(SUM(total_amount),0)").
		Where("order_date BETWEEN ? AND ?", startDate, endDate).
		Scan(&expense).Error
	if err != nil {
		return nil, err
	}

	report := &FinancialReport{
		StartDate:    startDate,
		EndDate:      endDate,
		TotalRevenue: revenue,
		TotalExpense: expense,
		NetProfit:    revenue - expense,
	}

	return report, nil
}
