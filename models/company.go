package models

type Company struct {
	ID         string   `json:"_id"`
	Name       string   `json:"name"`
	Features   []string `json:"features"`
	Language   string   `json:"language"`
	AdminGroup string   `json:"adminGroup"`
}
