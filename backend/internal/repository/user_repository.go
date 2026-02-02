package repository

import (
	"context"

	"book/shop/internal/domain/user"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) user.Repository {
	return &UserRepository{db: db}
}

// Query CreateUsrer
func (r *UserRepository) CreateUser(ctx context.Context, u *user.User) error {
	query := ` INSERT INTO users (user_name, email, password, role) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at`
	return r.db.QueryRow(ctx, query, u.UserName, u.Email, u.Password, u.Role).Scan(&u.ID, &u.CreatedAt, &u.UpdatedAt)
}
func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*user.User, error) {
	query := `SELECT id,user_name,email,password,role,created_at,updated_at FROM users WHERE email=$1 AND deleted_at IS NULL`
	u := &user.User{}
	err := r.db.QueryRow(ctx, query, email).Scan(&u.ID, &u.UserName, &u.Email, &u.Password, &u.Role, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// Query GetUserByID
func (r *UserRepository) GetUserByID(ctx context.Context, id string) (*user.User, error) {
	query := `SELECT id,user_name,email,role,created_at,updated_at FROM users WHERE id=$1 AND deleted_at IS NULL`
	u := &user.User{}
	err := r.db.QueryRow(ctx, query, id).Scan(&u.ID, &u.UserName, &u.Email, &u.Role, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// Query GetAllUser
func (r *UserRepository) GetAllUser(ctx context.Context) ([]*user.User, error) {
	query := `SELECT id,user_name,email,role,created_at,updated_at FROM users WHERE deleted_at IS NULL ORDER BY id DESC`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*user.User
	for rows.Next() {
		u := &user.User{}
		err := rows.Scan(&u.ID, &u.UserName, &u.Email, &u.Role, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

// Query UpdateUser
func (r *UserRepository) UpdateUser(ctx context.Context, u *user.User) error {
	query := `UPDATE users SET user_id = $1, user_name = $2, role = $3, updated_at = NOW() WHERE id = $4`

	_, err := r.db.Exec(ctx, query, u.UserID, u.UserName, u.Role, u.ID)

	return err

}

// Query DeleteUser
func (r *UserRepository) DeleteUser(ctx context.Context, id string) error {
	query := `	UPDATE users SET deleted_at = NOW() WHERE id=$1 AND deleted_at IS NULL`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
