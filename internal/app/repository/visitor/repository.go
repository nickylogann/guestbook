package visitor

import "context"

type Repository interface {
	Increment(ctx context.Context) (res int, err error)
	Get(ctx context.Context) (res int, err error)
}
