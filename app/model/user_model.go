package model

type UserResponse struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Username  string `json:"username,omitempty"`
	Email     string `json:"email,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Token     string `json:"token,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}

type VerifyUserRequest struct {
	Token string `validate:"required, max=255"`
}

type RegisterUserRequest struct {
	Name     string `validate:"required, max=255" json:"name"`
	Username string `validate:"required, max=255" json:"username"`
	Password string `validate:"required, max=255" json:"password"`
	Email    string `validate:"required, max=255" json:"email"`
	Phone    string `validate:"required, max=16" json:"phone"`
}

type UpdateUserRequest struct {
	Name     string `validate:"required, max=255" json:"name"`
	Email    string `validate:"required, max=255" json:"email"`
	Phone    string `validate:"required, max=16" json:"phone"`
	Password string `validate:"required, max=255" json:"password"`
}

type LoginUserRequest struct {
	Username string `validate:"required, max=255" json:"username"`
	Password string `validate:"required, max=255" json:"password"`
}

type LogoutUserRequest struct {
	ID string `validate:"required, max=255" json:"id"`
}

type GetUserRequest struct {
	ID string `validate:"required, max=255" json:"id"`
}
