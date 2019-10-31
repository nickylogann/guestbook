package functions

import (
	"context"

	"github.com/nickylogan/guestbook/internal/endpoint/models"
)

// FetchAll fetches all users with a pagination support
func (u *userUseCase) FetchAll(ctx context.Context, name string, count, offset int) (users []models.User, err error) {
	if count <= 0 {
		count = 10
	}
	if name == "" {
		return u.userRepo.Fetch(ctx, count, offset)
	} else {
		return u.userRepo.FetchFilterName(ctx, name, count, offset)
	}
}
