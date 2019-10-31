package functions

import (
	"github.com/nickylogan/guestbook/internal/endpoint/repository/visitor"
	. "github.com/nickylogan/guestbook/internal/endpoint/usecase/visitor"
)

type visitorUseCase struct {
	visitorRepo visitor.Repository
}

func NewVisitorUseCase(v visitor.Repository) UseCase {
	return &visitorUseCase{visitorRepo: v}
}
