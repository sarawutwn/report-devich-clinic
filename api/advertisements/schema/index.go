package AdsSchema

type GetAds struct {
	AdsID    string `json:"ads_id"`
	AdsName  string `json:"ads_name"`
	AdsPrice int    `json:"ads_price"`
}
