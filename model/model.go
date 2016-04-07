package model

// User struct describes a user's properties
type User struct {
	ID        string `bson:"_id" json:"id"`
	Email     string `bson:"email" json:"email"`
	FirstName string `bson:"first_name" json:"first_name"`
	LastName  string `bson:"last_name" json:"last_name"`
	Password  string `bson:"password" json:"password"`
	Role      string `bson:"role" json:"role"`
	Username  string `bson:"username" json:"username"`
}

// CreateUser is a struct that describes the properties for creating a new user
type CreateUser struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	Username  string `json:"username"`
}

// UpdateUser is a struct that describes the properties for updating a user
type UpdateUser struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	Username  string `json:"username"`
}

//JWTToken represts the JWTToken
type JWTToken string
