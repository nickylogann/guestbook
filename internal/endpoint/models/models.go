package models

import "time"

// User is an abstraction of a user type
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	MSISDN    string    `json:"msisdn"`
	Email     string    `json:"email"`
	BirthDate time.Time `json:"birth_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Age       int       `json:"age"`
}

// Visitor is an abstraction of the number of visitors
type Visitor struct {
	Count int
}
