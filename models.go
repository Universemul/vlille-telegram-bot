package main

import "fmt"

type Position struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type Field struct {
	State         string `json:"etat"`
	BikeAvailable int    `json:"nbvelosdispo"`
	SpotAvailable int    `json:"nbplacesdispo"`
	UpdatedDate   string `json:"datemiseajour"`
	TypeMarket    string `json:"type"`
	Name          string `json:"nom"`
	Address       string `json:"adresse"`
}

type ApiRecord struct {
	Recordid        string   `json:"recordid"`
	RecordTimestamp string   `json:"record_timestamp"`
	Geometry        Position `json:"geometry"`
	Field           Field    `json:"fields"`
}

type ApiResult struct {
	Nhits  int         `json:"nhits"`
	Result []ApiRecord `json:"records"`
}

func (record ApiRecord) Display(UserLat float32, UserLng float32) string {
	lat := record.Geometry.Coordinates[1]
	lng := record.Geometry.Coordinates[0]
	distance := compute_distance(lat, lng, float64(UserLat), float64(UserLng), "K")
	return fmt.Sprintf(`Station %s street %s
		Number of bikes available: %d
		Number of spots available: %d
		CreditCard available: %s
		Distance (in km): %.2f
		Itinerary: https://www.google.com/maps?q=%f,%f`,
		record.Field.Name, record.Field.Address, record.Field.BikeAvailable, record.Field.SpotAvailable, record.Field.State, distance, lat, lng)
}
