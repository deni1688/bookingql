package models

type Location struct {
	ID           string `json:"_id"`
	Name         string `json:"name"`
	StreetName   string `json:"streetName"`
	StreetNumber string `json:"streetNumber"`
	Postcode     string `json:"postcode"`
	City         string `json:"city"`
	Timezone     string `json:"timezone"`
}
