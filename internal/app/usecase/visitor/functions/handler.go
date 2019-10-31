package functions

import "context"

func (v visitorUseCase) Visit(ctx context.Context) (res int, err error) {
	return v.visitorRepo.Increment(ctx)
}
