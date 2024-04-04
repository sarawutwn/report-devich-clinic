package CustomerSchema

type CreateNewCustomer struct {
	CustomerName string `json:"customer_name" validate:"required"`
	AdsID        string `json:"ads_id"`
}

type UpdateStatusIsActive struct {
	CustomerID string `json:"customer_id" validate:"required"`
	IsActive   bool   `json:"is_active"`
}

type CreateInvoice struct {
	CustomerID  string   `json:"customer_id" validate:"required"`
	InvoiceDate string   `json:"invoice_date" validate:"required"`
	AdsID       []string `json:"ads_id" validate:"required"`
	Price       int      `json:"price"`
}

type UpdateBranch struct {
	AvgAmount string `json:"branch_avg_today"`
	DuoAmount int    `json:"branch_duo_amount"`
	BranchID  string `json:"branch_id" validate:"required"`
}

type AdsInvoice struct {
	AdsID    string
	AdsName  string
	AdsPrice int
}

type GetCustomer struct {
	CustomerID        string  `json:"customer_id"`
	CustomerName      string  `json:"customer_name"`
	CustomerStatus    string  `json:"customer_status"`
	CustomerIsActive  bool    `json:"customer_is_active"`
	CustomerPayAmount *int    `json:"customer_pay_amount"`
	CustomerPrice     *int    `json:"customer_price"`
	CreatedAt         string  `json:"created_at"`
	AdsID             *string `json:"ads_id"`
}

type Advertisements struct {
	AdsName string `json:"ads_name"`
}

type InvoiceItemData struct {
	InvoiceItemID   string      `json:"invoice_item_id"`
	InvoiceItemName string      `json:"invoice_item_name"`
	Invoice         InvoiceData `json:"invoice"`
}

type InvoiceData struct {
	InvoiceID          string       `json:"invoice_id"`
	InvoiceDate        string       `json:"invoice_date"`
	InvoiceTotalPrice  string       `json:"invoice_total_price"`
	InvoiceTotalAmount string       `json:"invoice_total_amount"`
	Customer           CustomerData `json:"customer"`
}

type CustomerData struct {
	CustomerName   string  `json:"customer_name"`
	CustomerStatus string  `json:"customer_status"`
	AdsID          *string `json:"ads_id"`
}

type GetCustomerReport struct {
	CustomerID        string  `json:"customer_id"`
	CustomerName      string  `json:"customer_name"`
	CustomerStatus    string  `json:"customer_status"`
	CustomerIsActive  bool    `json:"customer_is_active"`
	CustomerPayAmount *int    `json:"customer_pay_amount"`
	CustomerPrice     *int    `json:"customer_price"`
	CreatedAt         string  `json:"created_at"`
	AdsID             *string `json:"ads_id"`
	AdsName           *string `json:"ads_name"`
	AdsShortName      *string `json:"ads_short_name"`
}

type BranchData struct {
	BranchID        string `json:"branch_id"`
	BranchName      string `json:"branch_name"`
	BranchAvgToday  string `json:"branch_avg_today"`
	BranchAmount    int    `json:"branch_amount"`
	BranchDuoName   string `json:"branch_duo_name"`
	BranchDuoAmount int    `json:"branch_duo_amount"`
	BranchMeAmount  int    `json:"branch_me_amount"`
}
