package visitor

import "context"

type UseCase interface {
	Visit(ctx context.Context) (res int, err error)
}
