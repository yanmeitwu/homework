package models

type DriverPostInput struct {
	Capacity int `json:"capacity"`
}

type DriverPostOutput struct {
	ID int `json:"id"`
}

type DriverGetOutput struct {
	OfferID    int `json:"offerId"`
	ShipmentID int `json:"shipmentId"`
}
