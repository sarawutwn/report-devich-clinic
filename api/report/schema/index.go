package ReportSchema

type ReportDashbord struct {
	CountNewUsers []CountNewUsers `json:"count_new_users"`
	CountUser     int             `json:"count_user"`
	PointAll      *int            `json:"point_all"`
	UsePoint      *int            `json:"use_point"`
	CountToday    int             `json:"count_today"`
	CountReward   []CountNewUsers `json:"count_reward"`
}

type RewardQuota struct {
	QuotaId    string   `json:"quota_id"`
	QuataCode  string   `json:"quota_code"`
	RedeemDate any      `json:"redeem_date"`
	CreatedAt  string   `json:"created_at"`
	CustomerId *string  `json:"customer_id"`
	RewardId   string   `json:"reward_id"`
	Customer   Customer `json:"customer"`
}

type UserTransaction struct {
	CustomerId        string `json:"customer_id"`
	TransactionID     string `json:"transaction_id"`
	TransactionType   string `json:"transaction_type"`
	CreatedAt         string `json:"created_at"`
	StoreId           string `json:"store_id"`
	CustomerName      string `json:"customer_name"`
	TransactionDetail string `json:"transaction_detail"`
	TransactionAmount int    `json:"transaction_amount"`
}

type Customer struct {
	CustomerId          string `json:"customer_id"`
	CustomerName        string `json:"customer_name"`
	CustomerTelephone   string `json:"customer_telephone"`
	CustomerToken       string `json:"customer_token"`
	CustomerPointAmount int    `json:"customer_point_amount"`
	CustomerGender      string `json:"customer_gender"`
	CustomerBirthOfDate string `json:"customer_birth_of_date"`
	CustomerTotalPrice  int    `json:"customer_total_price"`
	CustomerPoint       int    `json:"customer_point"`
	CustomerAge         int    `json:"customer_age"`
	CreatedAt           string `json:"created_at"`
	UpdatedAt           string `json:"updated_at"`
	StoreId             string `json:"store_id"`
}

type CountNewUsers struct {
	CountNewUsers int    `json:"count"`
	Date          string `json:"date"`
}

type ReportDashbordOnDate struct {
	Gender []Gender `json:"gender"`
	// Age         []Age         `json:"age"`
	Transaction []Transaction `json:"transaction"`
	Frequency   []DataCount   `json:"frequency"`
	BirthDate   []DateCount   `json:"birth_date"`
	NewUsers    []DateCount   `json:"new_users"`
}

type Gender struct {
	GenderName string `json:"gender_name"`
	Count      int    `json:"count"`
}

type Age struct {
	Age   string `json:"age"`
	Count int    `json:"count"`
}

type Transaction struct {
	Date            string `json:"date"`
	TransactionType string `json:"Transaction_type"`
	Count           int    `json:"count"`
}

type DateCount struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

type DataCount struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}
