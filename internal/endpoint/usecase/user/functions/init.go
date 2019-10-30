package functions

import (
	rUser "github.com/nickylogan/guestbook/internal/endpoint/repository/user"
	uUser "github.com/nickylogan/guestbook/internal/endpoint/usecase/user"
)

type userUseCase struct {
	userRepo rUser.Repository
}

func NewUserUseCase(userRepo rUser.Repository) uUser.UseCase {
	return &userUseCase{userRepo: userRepo}
}
