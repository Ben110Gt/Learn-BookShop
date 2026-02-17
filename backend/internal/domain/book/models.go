package book

import (
	"book/shop/internal/domain/category"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	ID          uuid.UUID         `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	BookID      string            `gorm:"uniqueIndex" json:"book_id"`
	Title       string            `json:"title"`
	Author      string            `json:"author"`
	Description string            `json:"description"`
	Price       float64           `json:"price"`
	Stock       int               `json:"stock"`
	CategoryID  string            `json:"category_id"`
	Category    category.Category `gorm:"foreignKey:CategoryID;references:CategoryID" json:"category"`
	CoverImage  string            `json:"cover_image"`
	CreatedAt   time.Time         `gorm:"not null;default:now()" json:"created_at"`
	UpdatedAt   time.Time         `gorm:"not null;default:now()" json:"updated_at"`
	DeletedAt   gorm.DeletedAt    `gorm:"index" json:"-"`
}

// type Category struct {
// 	ID         uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
// 	CategoryID string         `gorm:"uniqueIndex" json:"category_id"`
// 	Name       string         `gorm:"uniqueIndex;not null" json:"name"`
// 	CreatedAt  time.Time      `gorm:"not null;default:now()" json:"created_at"`
// 	UpdatedAt  time.Time      `gorm:"not null;default:now()" json:"updated_at"`
// 	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
// }
