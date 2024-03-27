package ReportServices

import (
	ReportRepositorye "backend-app/api/report/repository"
	Schema "backend-app/api/report/schema"
	GoCache "backend-app/cache/go-cache"
	"strconv"
	"time"

	"github.com/xuri/excelize/v2"

	"encoding/json"
)

func ReportDashbord(StoreID string) (*Schema.ReportDashbord, error) {
	var KEY = "REPORT_DASHBOARD"
	// REDIS, _ := RedisCache.RegisterRedisCache().Get(REDIS_KEY)
	a, found := GoCache.Cache.Get(KEY)

	if found {
		data := Schema.ReportDashbord{}
		json.Unmarshal([]uint8(a.([]uint8)), &data)
		return &data, nil
	} else {
		return ReportRepositorye.ReportDashbord(StoreID)
	}

}

func ReportDashbordOnDate(StoreID string, StartDate string, EndDate string) (*Schema.ReportDashbordOnDate, error) {
	var KEY = "REPORT_DASHBOARD_DATE"
	// REDIS, _ := RedisCache.RegisterRedisCache().Get(KEY)
	a, found := GoCache.Cache.Get(KEY)
	data := Schema.ReportDashbordOnDate{}

	if found {
		json.Unmarshal([]uint8(a.([]uint8)), &data)
		return &data, nil
	} else {
		return ReportRepositorye.ReportDashbordOnDate(StoreID, StartDate, EndDate)
	}
}

func GetQuotaByStoreID(StoreID string, StartDate string, EndDate string, pageSize int, Page int) ([]Schema.RewardQuota, int, error) {
	var KEY = "REPORT_TRANSECTION_REDEEM"
	// REDIS, _ := RedisCache.RegisterRedisCache().Get(KEY)
	a, found := GoCache.Cache.Get(KEY)
	data := []Schema.RewardQuota{}

	if found {
		json.Unmarshal([]uint8(a.([]uint8)), &data)
		return data, 0, nil
	} else {
		return ReportRepositorye.GetQuotaByStoreID(StoreID, StartDate, EndDate, pageSize, Page)
	}
}

func GetUsersByStoreID(StoreID string, StartDate string, EndDate string, pageSize int, Page int) ([]Schema.Customer, int, error) {
	var KEY = "REPORT_USERS"
	// REDIS, _ := RedisCache.RegisterRedisCache().Get(KEY)
	a, found := GoCache.Cache.Get(KEY)
	data := []Schema.Customer{}

	if found {
		json.Unmarshal([]uint8(a.([]uint8)), &data)
		return data, 0, nil
	} else {
		return ReportRepositorye.GetUsersByStoreID(StoreID, StartDate, EndDate, pageSize, Page)
	}
}

func GetTransectionByStoreID(StoreID string, StartDate string, EndDate string, pageSize int, Page int) ([]Schema.UserTransaction, int, error) {
	var KEY = "REPORT_USERS"
	// REDIS, _ := RedisCache.RegisterRedisCache().Get(KEY)
	a, found := GoCache.Cache.Get(KEY)
	data := []Schema.UserTransaction{}

	if found {
		json.Unmarshal([]uint8(a.([]uint8)), &data)
		return data, 0, nil
	} else {
		return ReportRepositorye.GetTransectionByStoreID(StoreID, StartDate, EndDate, pageSize, Page)
	}
}

func ExportUsersByStoreID(StoreID string, StartDate string, EndDate string) (string, error) {

	data, err := ReportRepositorye.ExportUsersByStoreID(StoreID, StartDate, EndDate)
	// fmt.Println(data)
	if err != nil {
		return "", err
	}
	f := excelize.NewFile()

	if err := f.Close(); err != nil {
		return "", err
	}

	startChar := 'A'
	endChar := 'W'
	f.SetColWidth("Sheet1", "A", "C", 30)
	f.SetColWidth("Sheet1", "D", "K", 15)

	headers := map[string]string{"A1": "บัญชีผู้ใช้งานไลน์", "B1": "รหัสผู้ใช้งาน", "C1": "ขื่อ - นามสกุล",
		"D1": "เพศ", "E1": "วันที่ลงทะเบียน", "F1": "เบอร์โทรศัพท์", "G1": "วันเกิด", "H1": "อายุ", "I1": "คะแนนที่ได้", "J1": "คะแนนคงเหลือ", "K1": "ยอดใช้จ่าย"}
	for key, value := range headers {
		f.SetCellValue("Sheet1", key, value)
	}
	style, err := f.NewStyle(
		&excelize.Style{
			Alignment: &excelize.Alignment{
				Vertical:   "center",
				Horizontal: "right",
				WrapText:   true,
			},
		},
	)
	if err != nil {
		return "", err
	}
	for idx, row := range data {
		if err != nil {
			return "", err
		}
		// fmt.Println(row, idx)
		for char := startChar; char <= endChar; char++ {
			cell := string(char) + strconv.Itoa(idx+2)
			if char == 'A' {
				f.SetCellValue("Sheet1", cell, string(row.CustomerId))
			} else if char == 'B' {
				f.SetCellValue("Sheet1", cell, "AAA_"+string(row.CustomerId))
			} else if char == 'C' {
				f.SetCellValue("Sheet1", cell, string(row.CustomerName))
			} else if char == 'D' {
				f.SetCellValue("Sheet1", cell, string(row.CustomerGender))
			} else if char == 'E' {
				date, _ := time.Parse(time.RFC3339, string(row.CreatedAt))
				f.SetCellValue("Sheet1", cell, date.Format("02/01/2006"))
			} else if char == 'F' {
				f.SetCellValue("Sheet1", cell, string(row.CustomerTelephone))
			} else if char == 'G' {
				date, _ := time.Parse(time.RFC3339, row.CustomerBirthOfDate)
				f.SetCellValue("Sheet1", cell, date.Format("02/01/2006"))
			} else if char == 'H' {
				f.SetCellValue("Sheet1", cell, (row.CustomerAge))
			} else if char == 'I' {
				f.SetCellValue("Sheet1", cell, row.CustomerPoint)
			} else if char == 'J' {
				f.SetCellValue("Sheet1", cell, row.CustomerPointAmount)
			} else if char == 'K' {
				f.SetCellStyle("Sheet1", cell, cell, style)
				f.SetCellValue("Sheet1", cell, strconv.FormatFloat(float64(row.CustomerTotalPrice), 'f', 2, 64))
				// f.SetCellValue("Sheet1", cell, row.CustomerTotalPrice)

			}
			//
		} //
	}
	fileName := "export-file/" + "Book1.xlsx"
	if err := f.SaveAs(fileName); err != nil {
		return "", err
	}
	return fileName, nil
}

func ExportQuotaByStoreID(StoreID string, StartDate string, EndDate string) (string, error) {

	data, err := ReportRepositorye.ExportQuotaByStoreID(StoreID, StartDate, EndDate)
	// fmt.Println(data)
	if err != nil {
		return "", err
	}
	f := excelize.NewFile()

	if err := f.Close(); err != nil {
		return "", err
	}

	startChar := 'A'
	endChar := 'W'
	f.SetColWidth("Sheet1", "A", "B", 10)
	f.SetColWidth("Sheet1", "C", "D", 30)
	f.SetColWidth("Sheet1", "E", "K", 15)

	headers := map[string]string{"A1": "ลำดับ", "B1": "Code", "C1": "รหัสผู้ใช้งาน",
		"D1": "ชื่อ - นามสกุล", "E1": "เบอร์โทรศัพท์", "F1": "วันที่ Redeem"}
	for key, value := range headers {
		f.SetCellValue("Sheet1", key, value)
	}

	for idx, row := range data {
		if err != nil {
			return "", err
		}
		// fmt.Println(row, idx)
		for char := startChar; char <= endChar; char++ {
			cell := string(char) + strconv.Itoa(idx+2)
			if char == 'A' {
				f.SetCellValue("Sheet1", cell, strconv.Itoa(idx+1))
			} else if char == 'B' {
				f.SetCellValue("Sheet1", cell, string(row.QuataCode))
			} else if char == 'C' {
				f.SetCellValue("Sheet1", cell, "AAA_"+(*row.CustomerId))
			} else if char == 'D' {
				f.SetCellValue("Sheet1", cell, string(row.Customer.CustomerName))
			} else if char == 'E' {
				f.SetCellValue("Sheet1", cell, string(row.Customer.CustomerTelephone))
			} else if char == 'F' {
				date, _ := time.Parse(time.RFC3339, string(row.CreatedAt))
				f.SetCellValue("Sheet1", cell, date.Format("02/01/2006"))
			}
		} //
	}
	fileName := "export-file/" + "report redeem.xlsx"
	if err := f.SaveAs(fileName); err != nil {
		return "", err
	}
	return fileName, nil
}

func ExportTransectionByStoreID(StoreID string, StartDate string, EndDate string) (string, error) {

	data, err := ReportRepositorye.ExportTransectionByStoreID(StoreID, StartDate, EndDate)
	// fmt.Println(data)
	if err != nil {
		return "", err
	}
	f := excelize.NewFile()

	if err := f.Close(); err != nil {
		return "", err
	}

	startChar := 'A'
	endChar := 'W'
	f.SetColWidth("Sheet1", "A", "A", 10)
	f.SetColWidth("Sheet1", "B", "C", 30)
	f.SetColWidth("Sheet1", "D", "K", 15)

	headers := map[string]string{"A1": "ลำดับ", "B1": "รหัสผู้ใช้งาน", "C1": "ชื่อ - นามสกุล",
		"D1": "รายละเอียด", "E1": "จำนวน", "F1": "วันที่ทำรายการ"}
	for key, value := range headers {
		f.SetCellValue("Sheet1", key, value)
	}

	for idx, row := range data {
		if err != nil {
			return "", err
		}
		// fmt.Println(row, idx)
		for char := startChar; char <= endChar; char++ {
			cell := string(char) + strconv.Itoa(idx+2)
			if char == 'A' {
				f.SetCellValue("Sheet1", cell, strconv.Itoa(idx+1))
			} else if char == 'B' {
				f.SetCellValue("Sheet1", cell, "AAA_"+(row.CustomerId))
			} else if char == 'C' {
				f.SetCellValue("Sheet1", cell, string(row.CustomerName))
			} else if char == 'D' {
				f.SetCellValue("Sheet1", cell, string(row.TransactionDetail))
			} else if char == 'E' {
				f.SetCellValue("Sheet1", cell, (row.TransactionAmount))
			} else if char == 'F' {
				date, _ := time.Parse(time.RFC3339, string(row.CreatedAt))
				f.SetCellValue("Sheet1", cell, date.Format("02/01/2006"))
			}
		} //
	}
	fileName := "export-file/" + "report transection.xlsx"
	if err := f.SaveAs(fileName); err != nil {
		return "", err
	}
	return fileName, nil
}
