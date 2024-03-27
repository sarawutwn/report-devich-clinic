package ReportRepositorye

import (
	Schema "backend-app/api/report/schema"
	GoCache "backend-app/cache/go-cache"

	"backend-app/database"
	"encoding/json"
	"time"
)

func ReportDashbord(StoreID string) (*Schema.ReportDashbord, error) {
	db := database.DB
	query := `SELECT count(*) as count, sum(customer_point_amount) as point_all FROM customers WHERE store_id=$1` //get all customer
	smtp, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer smtp.Close()
	reports := Schema.ReportDashbord{}
	err = smtp.QueryRow(StoreID).Scan(&reports.CountUser, &reports.PointAll)
	if err != nil {
		return nil, err
	}

	query = `SELECT count(*) as count_today FROM customers WHERE store_id=$1 AND created_at=$2` //get all customer
	smtp, err = db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer smtp.Close()
	err = smtp.QueryRow(StoreID, time.Now().Format("2006-01-02 00:00:00")).Scan(&reports.CountToday)
	if err != nil {
		return nil, err
	}

	query = `SELECT count(*) as count,to_date(cast(created_at as TEXT),'YYYY-MM')  as date FROM customers 
	WHERE store_id=$1 and created_at>=$2
	GROUP BY date ORDER BY date desc`
	time6 := time.Now().AddDate(0, -6, 0)
	rows, err := db.Query(query, StoreID, time6)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	CountNewUsers := []Schema.CountNewUsers{}
	for rows.Next() {
		data := Schema.CountNewUsers{}
		err := rows.Scan(&data.CountNewUsers, &data.Date)
		if err != nil {
			return nil, err
		}
		CountNewUsers = append(CountNewUsers, data)
	}
	query = `SELECT count(*) as count,to_date(cast(r.redeem_date as TEXT),'YYYY-MM')  as date 
	FROM reward_quota r
	JOIN customers c ON c.customer_id=r.customer_id
	WHERE c.store_id=$1 and r.redeem_date>=$2
	GROUP BY date ORDER BY date desc`
	rows, err = db.Query(query, StoreID, time6)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	CountRewards := []Schema.CountNewUsers{}
	for rows.Next() {
		data := Schema.CountNewUsers{}
		err := rows.Scan(&data.CountNewUsers, &data.Date)
		if err != nil {
			return nil, err
		}
		CountRewards = append(CountRewards, data)
	}

	reports.CountNewUsers = CountNewUsers
	reports.CountReward = CountRewards

	query = `SELECT sum(transaction_amount) as use_point FROM transaction WHERE transaction_type='REDEEM' and store_id=$1` //get all customer
	smtp, err = db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer smtp.Close()
	err = smtp.QueryRow(StoreID).Scan(&reports.UsePoint)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(&reports)
	GoCache.Cache.Set("REPORT_DASHBOARD", data, 20*time.Minute)
	// RedisCache.RegisterRedisCache().Set("REPORT_DASHBOARD", string(data), 20*time.Second)
	return &reports, err
}

func ReportDashbordOnDate(StoreID string, StartDate string, EndDate string) (*Schema.ReportDashbordOnDate, error) {
	db := database.DB
	reports := Schema.ReportDashbordOnDate{}

	query := `SELECT customer_gender,count(*) count FROM customers WHERE store_id=$1 and created_at >= $2 and created_at <= $3 GROUP BY customer_gender`
	rows, err := db.Query(query, StoreID, StartDate, EndDate)
	if err != nil {
		return nil, err
	}
	genders := []Schema.Gender{}
	for rows.Next() {
		gender := Schema.Gender{}
		err = rows.Scan(&gender.GenderName, &gender.Count)
		if err != nil {
			return nil, err
		}
		genders = append(genders, gender)
	}
	defer rows.Close()

	reports.Gender = genders

	query = `SELECT to_date(cast(customer_birth_of_date as TEXT),'YYYY-MM-DD')  as birth_year,count(*) count FROM customers WHERE store_id=$1 and created_at >= $2 and created_at <= $3 GROUP BY birth_year`
	rows, err = db.Query(query, StoreID, StartDate, EndDate) //EXTRACT(YEAR FROM customer_birth_of_date::date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	birthdates := []Schema.DateCount{}
	for rows.Next() {
		birthdate := Schema.DateCount{}
		err = rows.Scan(&birthdate.Date, &birthdate.Count)
		if err != nil {
			return nil, err
		}
		birthdates = append(birthdates, birthdate)
	}

	bt := [6]Schema.DateCount{
		{Date: "1-20", Count: 0},
		{Date: "21-30", Count: 0},
		{Date: "31-40", Count: 0},
		{Date: "41-50", Count: 0},
		{Date: "51-60", Count: 0},
		{Date: "60+", Count: 0},
	}
	for _, value := range birthdates {
		date, _ := time.Parse(time.RFC3339, value.Date)
		age := calculateAge(date)
		switch {
		case age < 21:
			bt[0].Count += 1
		case age > 20 && age < 31:
			bt[1].Count += 1
		case age > 30 && age < 41:
			bt[2].Count += 1
		case age > 40 && age < 51:
			bt[3].Count += 1
		case age > 51 && age < 60:
			bt[4].Count += 1
		default:
			bt[5].Count += 1
		}
	}

	reports.BirthDate = bt[:]

	query = `SELECT customer_id,count(*) count FROM transaction WHERE store_id=$1 and created_at >= $2 and created_at <= $3 GROUP BY customer_id`
	rows, err = db.Query(query, StoreID, StartDate, EndDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	frequencys := []Schema.DataCount{}
	for rows.Next() {
		frequency := Schema.DataCount{}
		err = rows.Scan(&frequency.Name, &frequency.Count)
		if err != nil {
			return nil, err
		}
		frequencys = append(frequencys, frequency)
	}
	// reports.Frequency = frequencys
	ft := [3]Schema.DataCount{
		{Name: "1-2", Count: 0},
		{Name: "3-4", Count: 0},
		{Name: "5+", Count: 0},
	}
	for _, value := range frequencys {
		switch {
		case value.Count > 0 && value.Count < 3:
			ft[0].Count += 1
		case value.Count > 2 && value.Count < 5:
			ft[1].Count += 1
		default:
			ft[2].Count += 1
		}
	}

	reports.Frequency = ft[:]

	query = `SELECT to_date(cast(created_at as TEXT),'YYYY-MM-DD')  as date,count(*) count FROM customers WHERE store_id=$1 and created_at >= $2 and created_at <= $3 GROUP BY date`
	rows, err = db.Query(query, StoreID, StartDate, EndDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	newUsers := []Schema.DateCount{}
	for rows.Next() {
		newUser := Schema.DateCount{}
		err = rows.Scan(&newUser.Date, &newUser.Count)
		if err != nil {
			return nil, err
		}
		newUsers = append(newUsers, newUser)
	}
	reports.NewUsers = newUsers

	query = `SELECT transaction_type,to_date(cast(created_at as TEXT),'YYYY-MM-DD')  as date,count(*) count FROM transaction WHERE store_id=$1 and created_at >= $2 and created_at <= $3 GROUP BY date,transaction_type`
	rows, err = db.Query(query, StoreID, StartDate, EndDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	transactions := []Schema.Transaction{}
	for rows.Next() {
		transaction := Schema.Transaction{}
		err = rows.Scan(&transaction.TransactionType, &transaction.Date, &transaction.Count)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	reports.Transaction = transactions

	data, err := json.Marshal(&reports)
	GoCache.Cache.Set("REPORT_GRAPH", data, 20*time.Minute)

	return &reports, err
}

func GetQuotaByStoreID(StoreID string, StartDate string, EndDate string, pageSize int, Page int) ([]Schema.RewardQuota, int, error) {
	db := database.DB
	offset := (Page - 1) * pageSize
	rows, err := db.Query(`SELECT rq.quota_id, rq.quota_code, rq.redeem_date, rq.created_at, rq.customer_id, rq.reward_id, c.customer_name, c.customer_telephone
	FROM reward_quota rq LEFT JOIN customers c ON rq.customer_id = c.customer_id
	where c.store_id=$1 and rq.redeem_date >= $2 and rq.redeem_date <= $3 ORDER BY created_at asc LIMIT $4 OFFSET $5;`, StoreID, StartDate, EndDate, pageSize, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	quotas := []Schema.RewardQuota{}
	for rows.Next() {
		quota := Schema.RewardQuota{}
		err := rows.Scan(
			&quota.QuotaId, &quota.QuataCode, &quota.RedeemDate, &quota.CreatedAt, &quota.CustomerId,
			&quota.RewardId, &quota.Customer.CustomerName, &quota.Customer.CustomerTelephone,
		)
		if err != nil {
			return nil, 0, err
		}
		quotas = append(quotas, quota)
	}
	query := "SELECT COUNT(*) FROM reward_quota where store_id=$1 and redeem_date >= $2 and redeem_date <= $3"
	var count int = 0

	err = db.QueryRow(query, StoreID, StartDate, EndDate).Scan(&count)
	if err != nil {
		return nil, 0, err
	}
	return quotas, count, nil
}

func GetUsersByStoreID(StoreID string, StartDate string, EndDate string, pageSize int, Page int) ([]Schema.Customer, int, error) {
	db := database.DB
	offset := (Page - 1) * pageSize
	rows, err := db.Query(`SELECT customer_id, customer_name, customer_telephone, customer_gender, customer_birth_of_date, customer_token, customer_total_price, customer_point_amount, created_at
	FROM customers	
	where store_id=$1 and created_at >= $2 and created_at <= $3 ORDER BY created_at asc LIMIT $4 OFFSET $5;`, StoreID, StartDate, EndDate, pageSize, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	users := []Schema.Customer{}
	for rows.Next() {
		user := Schema.Customer{}
		err := rows.Scan(
			&user.CustomerId, &user.CustomerName, &user.CustomerTelephone, &user.CustomerGender, &user.CustomerBirthOfDate,
			&user.CustomerToken, &user.CustomerTotalPrice, &user.CustomerPointAmount, &user.CreatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		birth_date, err := time.Parse(time.RFC3339, user.CustomerBirthOfDate)
		if err != nil {
			return nil, 0, err
		}
		user.CustomerAge = calculateAge(birth_date)
		users = append(users, user)
	}
	query := "SELECT COUNT(*) FROM customers where store_id=$1 and created_at >= $2 and created_at <= $3"
	var count int = 0

	err = db.QueryRow(query, StoreID, StartDate, EndDate).Scan(&count)
	if err != nil {
		return nil, 0, err
	}
	return users, count, nil
}

func GetTransectionByStoreID(StoreID string, StartDate string, EndDate string, pageSize int, Page int) ([]Schema.UserTransaction, int, error) {
	db := database.DB
	offset := (Page - 1) * pageSize
	rows, err := db.Query(`SELECT c.customer_id, c.customer_name, t.transaction_type, t.transaction_detail, t.transaction_amount, t.created_at, t.store_id
	FROM transaction t JOIN customers c ON c.customer_id=t.customer_id
	where c.store_id=$1 and t.created_at >= $2 and t.created_at <= $3 ORDER BY t.created_at asc LIMIT $4 OFFSET $5;`, StoreID, StartDate, EndDate, pageSize, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	users := []Schema.UserTransaction{}
	for rows.Next() {
		user := Schema.UserTransaction{}
		err := rows.Scan(
			&user.CustomerId, &user.CustomerName, &user.TransactionType, &user.TransactionDetail, &user.TransactionAmount,
			&user.CreatedAt, &user.StoreId,
		)
		if err != nil {
			return nil, 0, err
		}
		users = append(users, user)
	}
	query := `SELECT COUNT(*) count FROM transaction t 
	JOIN customers c ON c.customer_id=t.customer_id where c.store_id=$1 and t.created_at >= $2 and t.created_at <= $3`
	var count int = 0

	err = db.QueryRow(query, StoreID, StartDate, EndDate).Scan(&count)
	if err != nil {
		return nil, 0, err
	}
	return users, count, nil
}

func calculateAge(dateOfBirth time.Time) int {
	now := time.Now()
	age := now.Year() - dateOfBirth.Year()

	// Check if the birthday has occurred this year
	if now.YearDay() < dateOfBirth.YearDay() {
		age--
	}

	return age
}

func ExportUsersByStoreID(StoreID string, StartDate string, EndDate string) ([]Schema.Customer, error) {
	db := database.DB
	rows, err := db.Query(`SELECT c.customer_id, c.customer_name, c.customer_telephone, c.customer_gender, c.customer_birth_of_date, c.customer_token, c.customer_total_price, c.customer_point_amount, c.created_at,sum(t.transaction_amount) as customer_point
	FROM customers c LEFT JOIN transaction t ON t.customer_id=c.customer_id
	where c.store_id=$1 and t.transaction_type='RECEIVE' and c.created_at >= $2 and c.created_at <= $3 
	GROUP BY c.customer_id ORDER BY c.created_at asc`, StoreID, StartDate, EndDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []Schema.Customer{}
	for rows.Next() {
		user := Schema.Customer{}
		err := rows.Scan(
			&user.CustomerId, &user.CustomerName, &user.CustomerTelephone, &user.CustomerGender, &user.CustomerBirthOfDate,
			&user.CustomerToken, &user.CustomerTotalPrice, &user.CustomerPointAmount, &user.CreatedAt, &user.CustomerPoint,
		)
		if err != nil {
			return nil, err
		}
		birth_date, err := time.Parse(time.RFC3339, user.CustomerBirthOfDate)
		if err != nil {
			return nil, err
		}
		user.CustomerAge = calculateAge(birth_date)
		users = append(users, user)
	}

	return users, nil
}

func ExportQuotaByStoreID(StoreID string, StartDate string, EndDate string) ([]Schema.RewardQuota, error) {
	db := database.DB
	rows, err := db.Query(`SELECT rq.quota_id, rq.quota_code, rq.redeem_date, rq.created_at, rq.customer_id, rq.reward_id, c.customer_name, c.customer_telephone
	FROM reward_quota rq LEFT JOIN customers c ON rq.customer_id = c.customer_id
	where c.store_id=$1 and rq.redeem_date >= $2 and rq.redeem_date <= $3 ORDER BY created_at asc;`, StoreID, StartDate, EndDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	quotas := []Schema.RewardQuota{}
	for rows.Next() {
		quota := Schema.RewardQuota{}
		err := rows.Scan(
			&quota.QuotaId, &quota.QuataCode, &quota.RedeemDate, &quota.CreatedAt, &quota.CustomerId,
			&quota.RewardId, &quota.Customer.CustomerName, &quota.Customer.CustomerTelephone,
		)
		if err != nil {
			return nil, err
		}
		quotas = append(quotas, quota)
	}

	return quotas, nil
}

func ExportTransectionByStoreID(StoreID string, StartDate string, EndDate string) ([]Schema.UserTransaction, error) {
	db := database.DB
	rows, err := db.Query(`SELECT c.customer_id, c.customer_name, t.transaction_type, t.transaction_detail, t.transaction_amount, t.created_at, t.store_id
	FROM transaction t JOIN customers c ON c.customer_id=t.customer_id
	where c.store_id=$1 and t.created_at >= $2 and t.created_at <= $3 ORDER BY t.created_at asc;`, StoreID, StartDate, EndDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []Schema.UserTransaction{}
	for rows.Next() {
		user := Schema.UserTransaction{}
		err := rows.Scan(
			&user.CustomerId, &user.CustomerName, &user.TransactionType, &user.TransactionDetail, &user.TransactionAmount,
			&user.CreatedAt, &user.StoreId,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
