package models

type Vehicle struct {
	ID          string  `json:"_id"`
	Vin         string  `json:"vin"`
	NumberPlate string  `json:"numberPlate"`
	Brand       string  `json:"brand"`
	Model       string  `json:"model"`
	Mileage     float64 `json:"mileage"`
}
