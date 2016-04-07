package reqres

import "gitlab.com/buzz/user/model"

// CreateUserRequest desribes the request for creating a new user
type CreateUserRequest struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	Username  string `json:"username"`
}

// CreateUserResponse desribes the response for creating a new user
type CreateUserResponse struct {
	User *model.User `json:"user"`
}

// GetUserResponse describes the response of getting an user by is or username
type GetUserResponse struct {
	User *model.User
}

// LoginRequest describes the request for a user to login
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse describes the response for a user to login
type LoginResponse struct {
	Token model.JWTToken `json:"token"`
}

// RefreshTokenRequest describes the request for refreshing a token
type RefreshTokenRequest struct {
	Token string `json:"token"`
}

// RefreshTokenResponse describes the response for refreshing a token
type RefreshTokenResponse struct {
	Token model.JWTToken `json:"token"`
}

/*****************************/
/* GENERIC RESPONSES */
/*****************************/

// ErrorResponse describes a response for when there is an error
type ErrorResponse struct {
	Message string `json:"message"`
}

// MessageResponse describes a message JSON response
type MessageResponse struct {
	Message string `json:"message"`
}
