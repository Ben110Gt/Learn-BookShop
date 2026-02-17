package service

import (
	"book/shop/internal/domain/book"
	"book/shop/internal/util"
	"context"
	"errors"
	"fmt"
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

	count, err := s.repo.Count(ctx)
	if err != nil {
		return nil, err
	}

	newid := util.GenerateCode("B", count)

	newBook := &book.Book{
		BookID:      newid,
		Title:       req.Title,
		Author:      req.Author,
		CategoryID:  req.CategoryID,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
	}
	fmt.Printf("DEBUG: CategoryID from Req: %s\n", req.CategoryID)
	if err := s.repo.CreateBook(ctx, newBook); err != nil {
		return nil, err
	}
	fmt.Printf("DEBUG: Book after Create: %+v\n", newBook)

	// if err := s.repo.CreateBook(ctx, newBook); err != nil {
	// 	return nil, err
	// }
	return newBook, nil
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
