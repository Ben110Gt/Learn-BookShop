package user

import "context"

type Repository interface {

	// CRUD
	Create(ctx context.Context, u *User) error
	Update(ctx context.Context, u *User) error
	Delete(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (*User, error)
	GetAll(ctx context.Context) ([]*User, error)

	// Auth Support
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}
