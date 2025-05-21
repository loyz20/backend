// internal/report/usecase.go
package report

import (
	"time"
)

type UseCase interface {
	GenerateStockReport(startDate, endDate time.Time, sku, name string) ([]StockReportRow, error)
	GeneratePurchaseReport(startDate, endDate time.Time, supplierName string) ([]PurchaseReportRow, error)
	GenerateSaleReport(startDate, endDate time.Time, customerName string) ([]SaleReportRow, error)
	GenerateFinancialReport(startDate, endDate time.Time) (*FinancialReport, error)
}

type usecase struct {
	repo Repository
}

func NewUseCase(repo Repository) UseCase {
	return &usecase{repo}
}

func (u *usecase) GenerateSaleReport(startDate, endDate time.Time, customerName string) ([]SaleReportRow, error) {
	return u.repo.GetSaleReport(startDate, endDate, customerName)
}

func (u *usecase) GenerateStockReport(startDate, endDate time.Time, sku, name string) ([]StockReportRow, error) {
	return u.repo.GetStockReport(startDate, endDate, sku, name)
}

func (u *usecase) GeneratePurchaseReport(startDate, endDate time.Time, supplierName string) ([]PurchaseReportRow, error) {
	return u.repo.GetPurchaseReport(startDate, endDate, supplierName)
}

func (u *usecase) GenerateFinancialReport(startDate, endDate time.Time) (*FinancialReport, error) {
	return u.repo.GetFinancialReport(startDate, endDate)
}
