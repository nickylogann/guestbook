package user

import (
	"context"

	"github.com/nickylogan/guestbook/internal/endpoint/usecase/user/models"
)

// UseCase represents use case related to the user
type UseCase interface {
	FetchAll(ctx context.Context, name string, page int) (res models.UserResponse, err error)
}
