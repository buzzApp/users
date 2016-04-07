package main

import "gitlab.com/buzz/user/reqres"

func validateCreateUser(user *reqres.CreateUserRequest) error {
	return nil
}

func validateGetUserByID(id string) error {
	return nil
}

func validateGetUserByUsername(username string) error {
	return nil
}

func validateLoginUser(payload *reqres.LoginRequest) error {
	return nil
}
