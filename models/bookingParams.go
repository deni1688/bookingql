package models

type BookingParams struct {
	Limit string `json:"limit"`
	Page  string `json:"page"`
	Sort
}
