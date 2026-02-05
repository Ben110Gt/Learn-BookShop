package book

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID          string         `gorm:"primaryKey" json:"id"`
	BookID      string         `gorm:"uniqueIndex" json:"book_id"`
	Title       string         `json:"title"`
	Author      string         `json:"author"`
	Description string         `json:"description"`
	Price       float64        `json:"price"`
	Stock       int            `json:"stock"`
	CategoryID  string         `json:"category_id"` // FK อยู่ที่ Book
	Category    Category       `gorm:"foreignKey:CategoryID;references:ID" json:"category"`
	CoverImage  string         `json:"cover_image"`
	CreatedAt   time.Time      `gorm:"not null;default:now()" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"not null;default:now()" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

type Category struct {
	ID         string         `gorm:"primaryKey" json:"id"`
	CategoryID string         `gorm:"uniqueIndex" json:"category_id"`
	Name       string         `gorm:"uniqueIndex;not null" json:"name"`
	CreatedAt  time.Time      `gorm:"not null;default:now()" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"not null;default:now()" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}
