package model

// User - user data structure
type User struct {
	ID       int    `json:"id" db:"id"`
	Email    string `json:"email" db:"email"`
	Password string `json:"-" db:"password"`
	Name     string `json:"name" db:"name"`
}
