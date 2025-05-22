package dashboard

type OrderDTO struct {
	ID           string  `json:"id"`
	CustomerName string  `json:"customerName"`
	Date         string  `json:"date"`
	Total        float64 `json:"total"`
}

type DashboardData struct {
	TotalOrders    int        `json:"totalOrders"`
	TotalProducts  int        `json:"totalProducts"`
	TotalCustomers int        `json:"totalCustomers"`
	RecentOrders   []OrderDTO `json:"recentOrders"`
}
