package user

type RegisterRequest struct {
	UserName string
	Email    string
	Password string
}

type LoginRequest struct {
	Email    string
	Password string
}

type UpdateRequest struct {
	UserName string
	Email    string
	Role     string
}

type LoginResponse struct {
	ID       uint   `json:"id"`
	UserID   string `json:"user_id"`
	Token    string `json:"token"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

type UserResponse struct {
	ID       uint   `json:"id"`
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

