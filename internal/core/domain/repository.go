package domain

import "context"

type Repository[T AggregateRoot] interface {
	Save(ctx context.Context, entity T) error
	FindByID(ctx context.Context, id string) (T, error)
	Delete(ctx context.Context, id string) error
	Exists(ctx context.Context, id string) (bool, error)
}

type ReadOnlyRepository[T any] interface {
	FindByID(ctx context.Context, id string) (T, error)
	FindAll(ctx context.Context, filter Filter) ([]T, error)
	Count(ctx context.Context, filter Filter) (int64, error)
	Exists(ctx context.Context, id string) (bool, error)
}

type Filter struct {
	Limit  int                    `json:"limit"`
	Offset int                    `json:"offset"`
	Sort   string                 `json:"sort"`
	Order  string                 `json:"order"`
	Where  map[string]interface{} `json:"where"`
}

func NewFilter() Filter {
	return Filter{
		Limit:  10,
		Offset: 0,
		Sort:   "created_at",
		Order:  "desc",
		Where:  make(map[string]interface{}),
	}
}
