package user

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserCode  string         `gorm:"size:10;uniqueIndex"`
	UserName  string         `json:"user_name"`
	Email     string         `gorm:"uniqueIndex" json:"email"`
	Password  string         `json:"-"`    // ไม่ส่งรหัสผ่านให้ client
	Role      string         `json:"role"` // admin / user
	CreatedAt time.Time      `gorm:"not null;default:now()" json:"created_at"`
	UpdatedAt time.Time      `gorm:"not null;default:now()" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
