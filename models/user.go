package models

type User struct {
	ID          string   `json:"_id"`
	Firstname   string   `json:"firstname"`
	Lastname    string   `json:"lastname"`
	Email       string   `json:"email"`
	Permissions []string `json:"permissions"`
}