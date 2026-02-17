package category

import (
	"context"
)

type Repository interface {
	CreateCategory(ctx context.Context, category *Category) error
	DeleteCategory(ctx context.Context, id string) error
	GetAllCategories(ctx context.Context) ([]*Category, error)
	Count(ctx context.Context) (int64, error)
}
