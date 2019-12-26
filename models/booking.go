package models

type Booking struct {
	ID                string  `json:"_id"`
	LocationStartID   string  `json:"locationStartId"`
	VehicleID         string  `json:"vehicleId"`
	UserID            string  `json:"userId"`
	StartDate         string  `json:"startDate"`
	EndDate           string  `json:"endDate"`
	Type              string  `json:"type"`
	State             string  `json:"state"`
	Created           string  `json:"created"`
	EstimatedDistance float64 `json:"estimatedDistance"`
	StartMileage      float64 `json:"startMileage"`
	EndMileage        float64 `json:"endMileage"`
	Cost              float64 `json:"cost"`
}
