package CustomerServices

import (
	CustomerRepository "backend-app/api/customer/repository"
	CustomerSchema "backend-app/api/customer/schema"
)

func CreateCustomer(CustomerName string, AdsID string) error {
	return CustomerRepository.CreateCustomerADS(CustomerName, AdsID)
}

func UpdateStatusReading(CustomerID string) ([]CustomerSchema.GetCustomer, error) {
	return CustomerRepository.UpdateStatusReading(CustomerID)
}

func UpdateIsActive(CustomerID string, IsActive bool) ([]CustomerSchema.GetCustomer, error) {
	return CustomerRepository.UpdateStatusIsActive(CustomerID, IsActive)
}

func GetCustomer() ([]CustomerSchema.GetCustomer, error) {
	return CustomerRepository.GetCustomerAll()
}

func CreateInvoice(customerID string, Price int, InvoiceDate string, AdsID []string) ([]CustomerSchema.AdsInvoice, error) {
	result, err := CustomerRepository.GetAdsList(AdsID)
	if err != nil {
		return nil, err
	}
	err = CustomerRepository.CreateInvoice(customerID, InvoiceDate, result, Price)
	return result, err
}

func UpdateBranchData(AvgAmount string, DuoAmount int, BranchID string) error {
	return CustomerRepository.UpdateBranchData(AvgAmount, DuoAmount, BranchID)
}

func GetCustomerReportAds() ([]CustomerSchema.GetCustomerReport, error) {
	return CustomerRepository.GetCustomerAds()
}

func GetReportCustomerInvoice() ([]CustomerSchema.InvoiceItemData, error) {
	return CustomerRepository.GetReportInvoice()
}

func GetBranch() (*CustomerSchema.BranchData, error) {
	return CustomerRepository.GetBranch()
}
