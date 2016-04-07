package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"gitlab.com/buzz/user/model"
	"gitlab.com/buzz/user/reqres"
)

func handleCreateUser(svc UserService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Read the body into a string for json decoding
		var payload = &reqres.CreateUserRequest{}
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			respondWithError("unable to decode json request", err, w, http.StatusInternalServerError)
			return
		}

		// Do some validation
		if err := validateCreateUser(payload); err != nil {
			respondWithError("Validation error", err, w, http.StatusBadRequest)
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
			respondWithError("unable to add user", err, w, http.StatusInternalServerError)
			return
		}

		// Generate our response
		resp := reqres.CreateUserResponse{User: user}

		// Marshal up the json response
		js, err := json.Marshal(resp)
		if err != nil {
			respondWithError("unable to marshal json response", err, w, http.StatusInternalServerError)
			return
		}

		// Return the response
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})
}

func handleGetUserByID(svc UserService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the user ID from the url
		id := mux.Vars(r)["id"]

		// Do some validation
		if err := validateGetUserByID(id); err != nil {
			respondWithError("Validation error", err, w, http.StatusBadRequest)
			return
		}

		// get the user from our database
		user, err := svc.GetByID(id)
		if err != nil {
			respondWithError("unable to get user", err, w, http.StatusInternalServerError)
			return
		}

		// Generate our response
		resp := reqres.GetUserResponse{User: user}

		// Marshal up the json response
		js, err := json.Marshal(resp)
		if err != nil {
			respondWithError("unable to marshal json response", err, w, http.StatusInternalServerError)
			return
		}

		// Return the response
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})
}

func handleGetUserByUsername(svc UserService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the user ID from the url
		username := mux.Vars(r)["username"]

		// Do some validation
		if err := validateGetUserByUsername(username); err != nil {
			respondWithError("Validation error", err, w, http.StatusBadRequest)
			return
		}

		// get the user from our database
		user, err := svc.GetByUsername(username)
		if err != nil {
			respondWithError("unable to get user", err, w, http.StatusInternalServerError)
			return
		}

		// Generate our response
		resp := reqres.GetUserResponse{User: user}

		// Marshal up the json response
		js, err := json.Marshal(resp)
		if err != nil {
			respondWithError("unable to marshal json response", err, w, http.StatusInternalServerError)
			return
		}

		// Return the response
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})
}

func handleLoginUser(svc UserService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Read the body into a string for json decoding
		var payload = &reqres.LoginRequest{}
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			respondWithError("unable to decode json request", err, w, http.StatusInternalServerError)
			return
		}

		// Do some validation
		if err := validateLoginUser(payload); err != nil {
			respondWithError("Validation error", err, w, http.StatusBadRequest)
			return
		}

		// save the app to our database
		token, err := svc.Login(payload.Username, payload.Password, r.Referer())
		if err != nil {
			respondWithError("unable to add user", err, w, http.StatusInternalServerError)
			return
		}

		// Generate our response
		resp := reqres.LoginResponse{Token: token}

		// Marshal up the json response
		js, err := json.Marshal(resp)
		if err != nil {
			respondWithError("unable to marshal json response", err, w, http.StatusInternalServerError)
			return
		}

		// Return the response
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})
}

// Helper function to return a json error message
func respondWithError(msg string, err error, w http.ResponseWriter, status int) {
	errMsg := reqres.ErrorResponse{Message: msg + ": " + err.Error()}

	js, err := json.Marshal(errMsg)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(status)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
}
