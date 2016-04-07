package main

import (
	"errors"
	"regexp"

	"gitlab.com/buzz/user/reqres"
)

func validateCreateUser(user *reqres.CreateUserRequest) error {
	if user.Email == "" || !isValidEmail(user.Email) {
		return errors.New("Invalid email address or email address not provided")
	}

	if user.FirstName == "" {
		return errors.New("Please provide a first name")
	}

	if user.LastName == "" {
		return errors.New("Please provide a last name")
	}

	if user.Username == "" {
		return errors.New("Please provide an username")
	}

	if user.Password == "" {
		return errors.New("Please provide a password")
	}

	if user.Role == "" {
		return errors.New("Please provide a role")
	}

	return nil
}

func validateGetUserByID(id string) error {
	if id == "" {
		return errors.New("Please provide an id")
	}

	return nil
}

func validateLoginUser(payload *reqres.LoginRequest) error {
	if payload.Username == "" {
		return errors.New("Please provide an username")
	}

	if payload.Password == "" {
		return errors.New("Please provide a password")
	}

	return nil
}

func validateRefreshToken(payload *reqres.RefreshTokenRequest) error {
	if payload.Token == "" {
		return errors.New("Please provide a token")
	}

	return nil
}

func isValidEmail(email string) bool {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return Re.MatchString(email)
}
