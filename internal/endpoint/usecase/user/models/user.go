package models

import "github.com/nickylogan/guestbook/internal/endpoint/models"

type UserResponse struct {
	Data     []models.User
	NextPage bool
	PrevPage bool
	Start    int
	End      int
	NumPages int
}
