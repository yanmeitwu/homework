package models

type ShipmentPostOutput struct {
	ID     int         `json:"id"`
	Offers []OfferResp `json:"offers"`
}

type ShipmentGetOutput struct {
	Accepted bool        `json:"accepted"`
	Offers   []OfferResp `json:"offers"`
}
