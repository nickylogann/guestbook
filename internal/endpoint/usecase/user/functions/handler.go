package functions

import (
	"context"
	"math"

	"github.com/nickylogan/guestbook/internal/endpoint/models"
	uModels "github.com/nickylogan/guestbook/internal/endpoint/usecase/user/models"
)

// FetchAll fetches all users with a pagination support
// page uses 1-based indexing
func (u *userUseCase) FetchAll(ctx context.Context, name string, page int) (response uModels.UserResponse, err error) {
	// Init repo params
	page = page - 1
	if page < 0 {
		page = 0
	}
	perPage := 20
	offset := perPage * page

	var users []models.User
	var rowCount int

	if name == "" {
		users, err = u.userRepo.Fetch(ctx, perPage, offset)
		rowCount, err = u.userRepo.CountFetch(ctx, perPage, offset)
	} else {
		users, err = u.userRepo.FetchFilterName(ctx, name, perPage, offset)
		rowCount, err = u.userRepo.CountFetchFilterName(ctx, name, perPage, offset)
	}

	numPages := int(math.Ceil(float64(rowCount) / float64(perPage)))

	// Pagination logic
	maxPage := 18
	start := page - maxPage/2
	if start < 1 {
		start = 1
	}
	end := start + maxPage - 1
	if end > numPages {
		end = numPages
	}

	response = uModels.UserResponse{
		Data:     users,
		PrevPage: page > 1,
		NextPage: page < (numPages - 1),
		Start:    start,
		End:      end,
	}
	return
}
