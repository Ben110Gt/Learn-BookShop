package book

import (
	"context"
)

type Repository interface {
	CreateBook(ctx context.Context, book *Book) error
	GetBookByID(ctx context.Context, id string) (*Book, error)
	GetBookByTitle(ctx context.Context, title string) (*Book, error)
	GetAllBooks(ctx context.Context) ([]*Book, error)
	UpdateBook(ctx context.Context, book *Book) (*Book, error)
	DeleteBook(ctx context.Context, id string) error
	Count(ctx context.Context) (int64, error)
}
