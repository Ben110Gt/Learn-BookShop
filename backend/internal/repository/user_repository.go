package repository

import (
	"context"
	"log"

	"book/shop/internal/domain/user"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.Repository {
	// return &UserRepository{db: db}
	if db == nil {
		// panic("DB connection is nil")
		log.Fatal("DB connection is nil")
	}
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *user.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (*user.User, error) {
	var user user.User
	err := r.db.Where("user_id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(ctx context.Context, user *user.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepository) Delete(ctx context.Context, id string) error {
	return r.db.Where("user_id = ?", id).Delete(&user.User{}).Error
}

func (r *UserRepository) GetAll(ctx context.Context) ([]*user.User, error) {
	var users []*user.User
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*user.User, error) {
	var user user.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&user.User{}).Count(&count).Error
	return count, err
}
