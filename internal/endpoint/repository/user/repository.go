package user

import (
	"context"

	"github.com/nickylogan/guestbook/internal/endpoint/models"
)

type Repository interface {
	Fetch(ctx context.Context, name string, count, offset int) (users []models.User, err error)
}