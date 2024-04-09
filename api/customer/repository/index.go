package CustomerRepository

import (
	CustomerSchema "backend-app/api/customer/schema"
	"backend-app/database"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

func CreateCustomerADS(customer_name string, adsID string) error {
	db := database.DB
	newUUID := uuid.New()
	if adsID != "" {
		query := `
			INSERT INTO customers(customer_id, customer_name, customer_status, customer_is_active, created_at, updated_at, ads_id)
			VALUES($1, $2, $3, $4, $5, $6, $7)
		`
		_, err := db.Exec(query, newUUID.String(), customer_name, "ไม่ตอบ", true, time.Now(), time.Now(), adsID)
		if err != nil {
			return err
		}
		return nil
	} else {
		query := `
			INSERT INTO customers(customer_id, customer_name, customer_status, customer_is_active, created_at, updated_at)
			VALUES($1, $2, $3, $4, $5, $6)
		`
		_, err := db.Exec(query, newUUID.String(), customer_name, "ไม่ตอบ", true, time.Now(), time.Now())
		if err != nil {
			return err
		}
		return nil
	}
}

func UpdateStatusReading(customerID string) ([]CustomerSchema.GetCustomer, error) {
	db := database.DB
	query := `UPDATE customers SET customer_status=$1 WHERE customer_id=$2;`
	_, err := db.Exec(query, "อ่านแล้ว", customerID)
	if err != nil {
		return nil, err
	}
	return GetCustomerAll()
}

func UpdateBranchData(AvgAmount string, DuoAmount int, BranchID string) error {
	db := database.DB
	query := `UPDATE branches SET branch_avg_today=$1, branch_duo_amount=$2 WHERE branch_id=$3`
	_, err := db.Exec(query, AvgAmount, DuoAmount, BranchID)
	if err != nil {
		return err
	}
	return nil
}

func UpdateStatusIsActive(customerID string, isActive bool) ([]CustomerSchema.GetCustomer, error) {
	db := database.DB
	query := `UPDATE customers SET customer_is_active=$1 WHERE customer_id=$2;`
	_, err := db.Exec(query, isActive, customerID)
	if err != nil {
		return nil, err
	}
	return GetCustomerAll()
}

func GetCustomerIDToday() ([]string, error) {
	db := database.DB
	rows, err := db.Query(`
		SELECT
		    customer_id
		FROM customers
		WHERE DATE(created_at) = CURRENT_DATE
		ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	var customerID []string
	for rows.Next() {
		var customer string
		err = rows.Scan(&customer)
		if err != nil {
			return nil, err
		}
		customerID = append(customerID, customer)
	}
	return customerID, nil
}

func GetReportInvoice() ([]CustomerSchema.InvoiceItemData, error) {
	db := database.DB
	rows, err := db.Query(`
		SELECT
			it.invoice_item_id,
			it.invoice_item_name,
			i.invoice_id,
			i.invoice_date,
			i.invoice_total_price,
			i.invoice_total_amount,
			c.customer_name,
			c.customer_status,
			c.ads_id
		FROM invoice_item it 
		LEFT JOIN invoice i ON i.invoice_id = it.invoice_id
		LEFT JOIN customers c ON c.customer_id = i.customer_id
		WHERE DATE(it.created_at) = CURRENT_DATE
		ORDER BY it.created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	var invoiceItem []CustomerSchema.InvoiceItemData
	for rows.Next() {
		var InvoiceItemData CustomerSchema.InvoiceItemData
		var InvoiceData CustomerSchema.InvoiceData
		var CustomerData CustomerSchema.CustomerData
		err = rows.Scan(&InvoiceItemData.InvoiceItemID, &InvoiceItemData.InvoiceItemName, &InvoiceData.InvoiceID, &InvoiceData.InvoiceDate, &InvoiceData.InvoiceTotalPrice, &InvoiceData.InvoiceTotalAmount, &CustomerData.CustomerName, &CustomerData.CustomerStatus, &CustomerData.AdsID)
		if err != nil {
			return nil, err
		}
		InvoiceData.Customer = CustomerData
		InvoiceItemData.Invoice = InvoiceData
		invoiceItem = append(invoiceItem, InvoiceItemData)
	}
	return invoiceItem, nil
}

func GetCustomerAds() ([]CustomerSchema.GetCustomerReport, error) {
	db := database.DB
	rows, err := db.Query(`
		SELECT
		    c.customer_id,
			c.customer_name,
			c.customer_status,
			c.customer_is_active,
			c.customer_pay_amount,
			c.customer_price,
			c.created_at,
			c.ads_id,
			a.ads_name,
			a.ads_short_name
		FROM customers c
		LEFT JOIN advertisements a ON c.ads_id = a.ads_id
		WHERE DATE(c.created_at) = CURRENT_DATE
		ORDER BY c.created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	var GetCustomerReport []CustomerSchema.GetCustomerReport
	for rows.Next() {
		var customerReport CustomerSchema.GetCustomerReport
		err = rows.Scan(&customerReport.CustomerID, &customerReport.CustomerName, &customerReport.CustomerStatus, &customerReport.CustomerIsActive, &customerReport.CustomerPayAmount, &customerReport.CustomerPrice, &customerReport.CreatedAt, &customerReport.AdsID, &customerReport.AdsName, &customerReport.AdsShortName)
		if err != nil {
			return nil, err
		}
		GetCustomerReport = append(GetCustomerReport, customerReport)
	}
	return GetCustomerReport, nil
}

func GetBranch() (*CustomerSchema.BranchData, error) {
	db := database.DB
	rows, err := db.Query(`
		SELECT
			branch_id,
		    branch_name,
			branch_amount,
			branch_duo_name,
			branch_duo_amount,
			branch_me_amount,
			branch_avg_today
		FROM branches;
	`)
	if err != nil {
		return nil, err
	}
	var branches CustomerSchema.BranchData
	for rows.Next() {
		rows.Scan(&branches.BranchID, &branches.BranchName, &branches.BranchAmount, &branches.BranchDuoName, &branches.BranchDuoAmount, &branches.BranchMeAmount, &branches.BranchAvgToday)
	}
	return &branches, nil
}

func GetCustomerAll() ([]CustomerSchema.GetCustomer, error) {
	db := database.DB
	rows, err := db.Query(`
		SELECT
		    customer_id,
			customer_name,
			customer_status,
			customer_is_active,
			customer_pay_amount,
			customer_price,
			created_at,
			ads_id
		FROM customers
		WHERE DATE(created_at) = CURRENT_DATE
		ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}

	var customers []CustomerSchema.GetCustomer
	for rows.Next() {
		var customer CustomerSchema.GetCustomer
		err = rows.Scan(&customer.CustomerID, &customer.CustomerName, &customer.CustomerStatus, &customer.CustomerIsActive, &customer.CustomerPayAmount, &customer.CustomerPrice, &customer.CreatedAt, &customer.AdsID)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}
	return customers, nil
}

func GetAdsList(AdsID []string) ([]CustomerSchema.AdsInvoice, error) {
	db := database.DB
	placeholders := make([]string, len(AdsID))
	parameters := make([]interface{}, len(AdsID))
	for i, id := range AdsID {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
		parameters[i] = id
	}
	queryAds := fmt.Sprintf(`SELECT ads_id, ads_name, ads_price FROM advertisements WHERE ads_id IN (%s)`, strings.Join(placeholders, ", "))
	rows, err := db.Query(queryAds, parameters...)
	if err != nil {
		return nil, err
	}
	var AdsList []CustomerSchema.AdsInvoice
	for rows.Next() {
		var Ads CustomerSchema.AdsInvoice
		err = rows.Scan(&Ads.AdsID, &Ads.AdsName, &Ads.AdsPrice)
		if err != nil {
			return nil, err
		}
		AdsList = append(AdsList, Ads)
	}

	return AdsList, nil
}

func CreateInvoice(CustomerID string, InvoiceDate string, AdsList []CustomerSchema.AdsInvoice, PayAmount int) error {
	db := database.DB
	query := `
		INSERT INTO invoice(invoice_id, invoice_date, invoice_total_price, invoice_total_amount, created_at, customer_id)
		VALUES($1, $2, $3, $4, $5, $6)
	`
	newUUID := uuid.New()
	total_price := 0
	for i := range AdsList {
		total_price += AdsList[i].AdsPrice
	}
	fmt.Println(total_price)
	fmt.Println(PayAmount)
	payer := 0
	if total_price-PayAmount == 0 {
		payer += total_price
	} else {
		payer = total_price - PayAmount
	}
	fmt.Println(payer)
	_, err := db.Exec(query, newUUID.String(), InvoiceDate, total_price, payer, time.Now(), CustomerID)
	if err != nil {
		return nil
	}
	query = "INSERT INTO invoice_item(invoice_item_id, invoice_item_name, invoice_price, created_at, invoice_id) VALUES "
	var values []interface{}
	for i, item := range AdsList {
		query += fmt.Sprintf("($%d,$%d,$%d,$%d,$%d),", i*5+1, i*5+2, i*5+3, i*5+4, i*5+5)
		values = append(values, uuid.New().String(), item.AdsName, item.AdsPrice, time.Now(), newUUID)
	}
	query = query[:len(query)-1]
	_, err = db.Exec(query, values...)
	if err != nil {
		return nil
	}
	query = `UPDATE customers SET customer_status=$1, customer_pay_amount=$2, customer_price=$3 WHERE customer_id=$4;`

	_, err = db.Exec(query, "จอง", PayAmount, total_price, CustomerID)
	if err != nil {
		return nil
	}
	return nil
}
