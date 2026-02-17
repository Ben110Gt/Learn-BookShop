package service

import (
	"book/shop/internal/domain/user"
	"book/shop/internal/util"
	"context"
	"errors"
)

type UserService struct {
	repo user.Repository
}

func NewUserService(repo user.Repository) user.Service {
	return &UserService{repo: repo}
}

// Register a new user
func (s *UserService) Register(ctx context.Context, req user.RegisterRequest) (*user.User, error) {
	existing, _ := s.repo.GetUserByEmail(ctx, req.Email)
	if existing != nil {
		return nil, errors.New("email already exists")
	}
	count, err := s.repo.Count(ctx)
	if err != nil {
		return nil, err
	}

	hash, _ := util.HashPassword(req.Password)
	code := util.GenerateCode("U", count)

	u := &user.User{
		UserCode: code,
		UserName: req.UserName,
		Email:    req.Email,
		Password: hash,
		Role:     "customer", // Default role
	}

	if err := s.repo.Create(ctx, u); err != nil {
		return nil, err
	}
	return u, nil
}

// Login an existing user
func (s *UserService) Login(ctx context.Context, req user.LoginRequest) (user.LoginResponse, error) {
	u, err := s.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return user.LoginResponse{}, errors.New("invalid email or password")
	}
	err = util.CheckPassword(req.Password, u.Password)
	if err != nil {
		return user.LoginResponse{}, errors.New("invalid email or password")
	}
	token, err := util.GenerateToken(u.UserCode, u.UserName, u.Role)
	if err != nil {
		return user.LoginResponse{}, errors.New("failed to generate token")
	}
	return user.LoginResponse{

		Token:    token,
		Username: u.UserName,
		Role:     u.Role,
	}, nil
}

// GetUserProfile retrieves user profile by ID
func (s *UserService) GetUserByID(ctx context.Context, id string) (*user.User, error) {
	return s.repo.GetByID(ctx, id)
}

// GetAllUser retrieves all users
func (s *UserService) GetAllUser(ctx context.Context) ([]*user.User, error) {
	return s.repo.GetAll(ctx)
}

// UpdateUserProfile updates user profile
func (s *UserService) UpdateUser(ctx context.Context, id string, req user.UpdateRequest) error {
	u, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	u.UserName = req.UserName
	u.Email = req.Email
	u.Role = req.Role
	return s.repo.Update(ctx, u)
}

// DeleteUserProfile deletes user profile
func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
