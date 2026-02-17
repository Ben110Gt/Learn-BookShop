package service

import (
	"book/shop/internal/domain/category"
	"book/shop/internal/util"
	"context"
)

type CategoryService struct {
	repo category.Repository
}

func NewCategoryService(repo category.Repository) category.Service {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) CreateCategory(ctx context.Context, cat *category.Category) error {
	count, err := s.repo.Count(ctx)
	if err != nil {
		return err
	}

	newid := util.GenerateCode("C", count)

	newCategory := &category.Category{
		CategoryID: newid,
		Name:       cat.Name,
	}
	if err := s.repo.CreateCategory(ctx, newCategory); err != nil {
		return err
	}
	return nil
}

func (s *CategoryService) DeleteCategory(ctx context.Context, id string) error {
	return s.repo.DeleteCategory(ctx, id)
}

func (s *CategoryService) Count(ctx context.Context) (int64, error) {
	return s.repo.Count(ctx)
}

func (s *CategoryService) GetAllCategories(ctx context.Context) ([]*category.Category, error) {
	categories, err := s.repo.GetAllCategories(ctx)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
