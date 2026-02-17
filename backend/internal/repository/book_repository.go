package repository

import (
	"book/shop/internal/domain/book"
	"context"

	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) book.Repository {
	return &BookRepository{db: db}
}

func (r *BookRepository) CreateBook(ctx context.Context, book *book.Book) error {
	if err := r.db.WithContext(ctx).Create(book).Error; err != nil {
		return err
	}
	if err := r.db.WithContext(ctx).Preload("Category").First(book, "id = ?", book.ID).Error; err != nil {
		return err
	}
	return nil

}

func (r *BookRepository) GetBookByID(ctx context.Context, id string) (*book.Book, error) {
	var b book.Book
	if err := r.db.WithContext(ctx).First(&b, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &b, nil
}

func (r *BookRepository) GetBookByTitle(ctx context.Context, title string) (*book.Book, error) {
	var b book.Book
	if err := r.db.WithContext(ctx).First(&b, "title = ?", title).Error; err != nil {
		return nil, err
	}
	return &b, nil
}

func (r *BookRepository) GetAllBooks(ctx context.Context) ([]*book.Book, error) {
	var books []*book.Book
	if err := r.db.WithContext(ctx).Find(&books).Error; err != nil {
		return nil, err
	}
	if err := r.db.WithContext(ctx).Preload("Category").Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func (r *BookRepository) UpdateBook(ctx context.Context, book *book.Book) (*book.Book, error) {
	if err := r.db.WithContext(ctx).Save(book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func (r *BookRepository) DeleteBook(ctx context.Context, id string) error {
	if err := r.db.WithContext(ctx).Delete(&book.Book{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (r *BookRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&book.Book{}).Count(&count).Error
	return count, err
}
