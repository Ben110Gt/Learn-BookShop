package category

import (
	"context"
)

type Service interface {
	CreateCategory(ctx context.Context, category *Category) error
	DeleteCategory(ctx context.Context, id string) error
	Count(ctx context.Context) (int64, error)
	GetAllCategories(ctx context.Context) ([]*Category, error)
}
