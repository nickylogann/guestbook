package functions

import (
	"context"

	"github.com/nickylogan/guestbook/internal/endpoint/models"
)

// FetchFilterName fetches all users with a similar name
func (u *userUseCase) FetchFilterName(ctx context.Context, name string, count, offset int) (users []models.User, err error) {
	if count <= 0 {
		count = 10
	}
	return u.userRepo.Fetch(ctx, name, count, offset)
}

// FetchAll fetches all users with a pagination support
func (u *userUseCase) FetchAll(ctx context.Context, count, offset int) (users []models.User, err error) {
	if count <= 0 {
		count = 10
	}
	return u.userRepo.Fetch(ctx, "", count, offset)
}
