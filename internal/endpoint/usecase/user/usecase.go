package user

import (
	"context"

	"github.com/nickylogan/guestbook/internal/endpoint/models"
)

// UseCase represents use case related to the user
type UseCase interface {
	FetchFilterName(ctx context.Context, name string, count, offset int) (users []models.User, err error)
	FetchAll(ctx context.Context, count, offset int) (users []models.User, err error)
}