package domain

import "time"

// User - Estructura de un usuario
type User struct {
	UserID    int64     `gorm:"primary_key" json:"user_id"`
	Username  string    `gorm:"size:255" json:"username"`
	Password  string    `gorm:"size:255" json:"password"`
	FirstName string    `gorm:"size:255" json:"firstname"`
	LastName  string    `gorm:"size:255" json:"lastname"`
	Email     string    `gorm:"size:255" json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
