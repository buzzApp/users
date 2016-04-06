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
