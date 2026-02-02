package user

import (
	"context"
)

type Service interface {

	// Auth
	Register(ctx context.Context, req RegisterRequest) (*User, error)
	Login(ctx context.Context, req LoginRequest) (LoginResponse, error)

	// CRUD
	GetUserByID(ctx context.Context, id string) (*User, error)
	GetAllUser(ctx context.Context) ([]*User, error)
	UpdateUser(ctx context.Context, id string, req UpdateRequest) error
	DeleteUser(ctx context.Context, id string) error
}
