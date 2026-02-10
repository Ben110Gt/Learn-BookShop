package service

import (
	"book/shop/internal/domain/book"
	"context"
	"errors"
)

type BookService struct {
	repo book.Repository
}

func NewBookService(repo book.Repository) book.Service {
	return &BookService{repo: repo}
}

func (s *BookService) CreateBook(ctx context.Context, req *book.CreateBookRequest) (*book.Book, error) {
	existingBook, _ := s.repo.GetBookByTitle(ctx, req.Title)
	if existingBook != nil {
		return nil, errors.New("book already exists")
	}
	newBook := &book.Book{
		Title:       req.Title,
		Author:      req.Author,
		Description: req.Description,
		Price:       req.Price,
	}
	return s.repo.CreateBook(ctx, newBook)

}

func (s *BookService) GetBookByID(ctx context.Context, id string) (*book.Book, error) {
	return s.repo.GetBookByID(ctx, id)
}

func (s *BookService) GetAllBooks(ctx context.Context) ([]*book.Book, error) {
	return s.repo.GetAllBooks(ctx)
}

func (s *BookService) UpdateBook(ctx context.Context, book *book.Book) (*book.Book, error) {
	return s.repo.UpdateBook(ctx, book)
}

func (s *BookService) DeleteBook(ctx context.Context, id string) error {
	return s.repo.DeleteBook(ctx, id)
}
