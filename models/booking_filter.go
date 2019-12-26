package models

type BookingFilter struct {
	Limit string `json:"limit"`
	Page  string `json:"page"`
	Sort
}
