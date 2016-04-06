package main

import (
	"encoding/json"
	"net/http"

	"gitlab.com/buzz/user/model"
	"gitlab.com/buzz/user/reqres"
)

func handleCreateUser(svc UserService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Read the body into a string for json decoding
		var payload = &reqres.CreateUserRequest{}
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			respondWithError("unable to decode json request", err, w)
			return
		}

		// Do some validation
		if err := validateCreateUser(payload); err != nil {
			respondWithError("Validation error", err, w)
			return
		}

		// Create our new user struct
		newUser := &model.CreateUser{
			Email:     payload.Email,
			FirstName: payload.FirstName,
			LastName:  payload.LastName,
			Password:  payload.Password,
			Role:      payload.Role,
			Username:  payload.Username,
		}

		// save the app to our database
		user, err := svc.Create(newUser)
		if err != nil {
			respondWithError("unable to add app", err, w)
			return
		}

		// Generate our response
		resp := reqres.CreateUserResponse{User: user}

		// Marshal up the json response
		js, err := json.Marshal(resp)
		if err != nil {
			respondWithError("unable to marshal json response", err, w)
			return
		}

		// Return the response
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})
}

// Helper function to return a json error message
func respondWithError(msg string, err error, w http.ResponseWriter) {
	errMsg := reqres.ErrorResponse{Message: msg + ": " + err.Error()}

	js, err := json.Marshal(errMsg)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(js)
}
