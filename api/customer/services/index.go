package CustomerServices

import (
	CustomerRepository "backend-app/api/customer/repository"
	CustomerSchema "backend-app/api/customer/schema"
)

func CreateCustomer(CustomerName string, AdsID string) error {
	return CustomerRepository.CreateCustomerADS(CustomerName, AdsID)
}

func UpdateStatusReading(CustomerID string) error {
	return CustomerRepository.UpdateStatusReading(CustomerID)
}

func UpdateIsActive(CustomerID string, IsActive bool) error {
	return CustomerRepository.UpdateStatusIsActive(CustomerID, IsActive)
}

func GetCustomer() ([]CustomerSchema.GetCustomer, error) {
	return CustomerRepository.GetCustomerAll()
}
