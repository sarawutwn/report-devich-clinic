package AdsRepository

import (
	AdsSchema "backend-app/api/advertisements/schema"
	"backend-app/database"
)

func GetAdsList() ([]AdsSchema.GetAds, error) {
	db := database.DB
	rows, err := db.Query(`
		SELECT
			ads_id,
			ads_name
		FROM advertisements
		ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	var adsLists []AdsSchema.GetAds
	for rows.Next() {
		var adsList AdsSchema.GetAds
		err = rows.Scan(&adsList.AdsID, &adsList.AdsName)
		if err != nil {
			return nil, err
		}
		adsLists = append(adsLists, adsList)
	}
	return adsLists, nil
}
