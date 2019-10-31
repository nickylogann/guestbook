package user

import (
	"context"

	"github.com/nickylogan/guestbook/internal/endpoint/models"
)

type Repository interface {
	Fetch(ctx context.Context, count, offset int) (users []models.User, err error)
	CountFetch(ctx context.Context, count, offset int) (res int, err error)
	FetchFilterName(ctx context.Context, name string, count, offset int) (users []models.User, err error)
	CountFetchFilterName(ctx context.Context, name string, count, offset int) (res int, err error)
}
