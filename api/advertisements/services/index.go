package AdsServices

import (
	AdsRepository "backend-app/api/advertisements/repository"
	AdsSchema "backend-app/api/advertisements/schema"
)

func GetAdsList() ([]AdsSchema.GetAds, error) {
	return AdsRepository.GetAdsList()
}
