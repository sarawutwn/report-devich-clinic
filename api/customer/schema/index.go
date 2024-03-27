package CustomerSchema

type CreateNewCustomer struct {
	CustomerName string `json:"customer_name" validate:"required"`
	AdsID        string `json:"ads_id"`
}

type UpdateStatusIsActive struct {
	CustomerID string `json:"customer_id" validate:"required"`
	IsActive   bool   `json:"is_active"`
}

type GetCustomer struct {
	CustomerID       string  `json:"customer_id"`
	CustomerName     string  `json:"customer_name"`
	CustomerStatus   string  `json:"customer_status"`
	CustomerIsActive bool    `json:"customer_is_active"`
	CreatedAt        string  `json:"created_at"`
	AdsID            *string `json:"ads_id"`
}

type Advertisements struct {
	AdsName string `json:"ads_name"`
}
