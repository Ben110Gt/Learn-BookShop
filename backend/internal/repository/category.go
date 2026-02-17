package repository

import (
	"book/shop/internal/domain/category"
	"context"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) CreateCategory(ctx context.Context, category *category.Category) error {
	return r.db.WithContext(ctx).Create(category).Error
}

func (r *CategoryRepository) DeleteCategory(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&category.Category{}, "id = ?", id).Error
}

func (r *CategoryRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&category.Category{}).Count(&count).Error
	return count, err
}

func (r *CategoryRepository) GetAllCategories(ctx context.Context) ([]*category.Category, error) {

	var categories []*category.Category
	if err := r.db.WithContext(ctx).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
