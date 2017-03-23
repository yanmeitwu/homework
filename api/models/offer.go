package models

type OfferPutInput struct {
	Status string `json:"status"`
}

type OfferResp struct {
	OfferID  int `json:"offerId"`
	DriverID int `json:"driverId"`
}
