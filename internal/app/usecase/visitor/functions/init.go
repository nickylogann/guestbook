package functions

import (
	"github.com/nickylogan/guestbook/internal/app/repository/visitor"
	. "github.com/nickylogan/guestbook/internal/app/usecase/visitor"
)

type visitorUseCase struct {
	visitorRepo visitor.Repository
}

func NewVisitorUseCase(v visitor.Repository) UseCase {
	return &visitorUseCase{visitorRepo: v}
}
