// internal/report/entity.go
package report

import "time"

type StockReportRow struct {
	ProductID         uint    `json:"product_id"`
	Name              string  `json:"name"`
	SKU               string  `json:"sku"`
	TotalStock        int     `json:"total_stock"`
	TotalBatch        int     `json:"total_batch"`
	TotalValue        float64 `json:"total_value"`
	LatestBatchExpiry string  `json:"latest_batch_expiry,omitempty"`
}

type PurchaseReportRow struct {
	PurchaseOrderID uint      `json:"purchase_order_id"`
	OrderNumber     string    `json:"order_number"`
	SupplierName    string    `json:"supplier_name"`
	OrderDate       time.Time `json:"order_date"`
	TotalAmount     float64   `json:"total_amount"`
}

type SaleReportRow struct {
	SaleOrderID  uint      `json:"sale_order_id"`
	OrderNumber  string    `json:"order_number"`
	CustomerName string    `json:"customer_name"`
	OrderDate    time.Time `json:"order_date"`
	TotalAmount  float64   `json:"total_amount"`
}

type FinancialReport struct {
	StartDate    time.Time `json:"start_date"`
	EndDate      time.Time `json:"end_date"`
	TotalRevenue float64   `json:"total_revenue"` // dari penjualan
	TotalExpense float64   `json:"total_expense"` // dari pembelian
	NetProfit    float64   `json:"net_profit"`    // totalRevenue - totalExpense
}
