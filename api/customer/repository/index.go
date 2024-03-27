package CustomerRepository

import (
	CustomerSchema "backend-app/api/customer/schema"
	"backend-app/database"
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

func UpdateStatusReading(customerID string) error {
	db := database.DB
	query := `UPDATE customers SET customer_status=$1 WHERE customer_id=$2;`
	_, err := db.Exec(query, "อ่านแล้ว", customerID)
	if err != nil {
		return err
	}
	return nil
}

func UpdateStatusIsActive(customerID string, isActive bool) error {
	db := database.DB
	query := `UPDATE customers SET customer_is_active=$1 WHERE customer_id=$2;`
	_, err := db.Exec(query, isActive, customerID)
	if err != nil {
		return err
	}
	return nil
}

func GetCustomerAll() ([]CustomerSchema.GetCustomer, error) {
	db := database.DB
	rows, err := db.Query(`
		SELECT
		    customer_id,
			customer_name,
			customer_status,
			customer_is_active,
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
		err = rows.Scan(&customer.CustomerID, &customer.CustomerName, &customer.CustomerStatus, &customer.CustomerIsActive, &customer.CreatedAt, &customer.AdsID)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}
	return customers, nil
}
