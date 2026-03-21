package models

type User struct {
	ID    uint   `json:"id" gorm:"primaryKey" example:"1"`
	Name  string `json:"name" example:"John Doe"`
	Email string `json:"email" example:"john@example.com"`
}
