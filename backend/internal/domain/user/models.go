package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	UserID    string         `gorm:"uniqueIndex" json:"user_id"`
	UserName  string         `json:"user_name"`
	Email     string         `gorm:"uniqueIndex" json:"email"`
	Password  string         `json:"-"`    // ไม่ส่งรหัสผ่านให้ client
	Role      string         `json:"role"` // admin / user
	CreatedAt time.Time      `gorm:"not null;default:now()" json:"created_at"`
	UpdatedAt time.Time      `gorm:"not null;default:now()" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
