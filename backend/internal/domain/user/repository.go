package user

import "context"

type Repository interface {

	// CRUD
	CreateUser(ctx context.Context, u *User) error
	UpdateUser(ctx context.Context, u *User) error
	DeleteUser(ctx context.Context, id string) error

	GetUserByID(ctx context.Context, id string) (*User, error)
	GetAllUser(ctx context.Context) ([]*User, error)

	// Auth Support
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}
