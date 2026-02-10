package book

import (
	"context"
)

type Service interface {
	CreateBook(ctx context.Context, req *CreateBookRequest) (*Book, error)
	GetBookByID(ctx context.Context, id string) (*Book, error)
	GetAllBooks(ctx context.Context) ([]*Book, error)
	UpdateBook(ctx context.Context, book *Book) (*Book, error)
	DeleteBook(ctx context.Context, id string) error
}
