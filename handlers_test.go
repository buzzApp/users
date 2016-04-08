package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	"gitlab.com/buzz/user/model"
	"gitlab.com/buzz/user/reqres"
)

var (
	userID   string
	username = "testUser"
	password = "testPassword"
	token    model.JWTToken
)

func TestMain(m *testing.M) {
	result := m.Run()

	tearDown()

	os.Exit(result)
}

func TestCreateUserHTTPEndpoint(t *testing.T) {
	server := httptest.NewServer(handleCreateUser(userService{}))
	defer server.Close()

	usersURL := fmt.Sprintf("%s/users", server.URL)

	userJSON := `{"email": "test@test.com", "first_name": "testname", "last_name": "lasttest", "password": "` + password + `", "role": "student", "username": "` + username + `"}`

	jsonReader := strings.NewReader(userJSON)

	req, err := http.NewRequest("POST", usersURL, jsonReader)
	if err != nil {
		t.Error(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 201 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}

	getUserID(res.Body)
}

func TestGetUserByIDHTTPEndpoint(t *testing.T) {
	// Create a new mux router
	router := mux.NewRouter()

	router.Handle("/users/{id}", handleGetUserByID(userService{}))

	// Create a new server
	server := httptest.NewServer(router)

	// get url
	getByIDURL := fmt.Sprintf("%s/users/%s", server.URL, userID)

	resp, err := http.Get(getByIDURL)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("Expected a 200 status code response but got: %d", resp.StatusCode)
	}
}

func TestLoginUser(t *testing.T) {
	server := httptest.NewServer(handleLoginUser(userService{}))

	loginURL := fmt.Sprintf("%s/auth/authenticate", server.URL)

	loginJSON := `{"username": "` + username + `", "password": "` + password + `"}`

	req, _ := http.NewRequest("POST", loginURL, strings.NewReader(loginJSON))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("Expected a 200 status code response but got: %d", resp.StatusCode)
	}

	getToken(resp.Body)
}

func testRefreshToken(t *testing.T) {
	server := httptest.NewServer(handleLoginUser(userService{}))

	refreshTokenURL := fmt.Sprintf("%s/auth/refresh-token", server.URL)

	refreshTokenJSON := `{"token": "` + string(token) + `"}`

	req, _ := http.NewRequest("POST", refreshTokenURL, strings.NewReader(refreshTokenJSON))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("Expected a 200 status code response but got: %d", resp.StatusCode)
	}
}

func tearDown() {
	//Grab a copy of our session
	session, err := getSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	//Get our collection of applications
	db := session.DB("buzz-test-user")
	collection := db.C("users")

	//remove our applications from the collection
	err = collection.Remove(bson.M{"username": username})
	if err != nil {
		log.Fatal(err)
	}
}

func getUserID(respBody io.Reader) {
	var payload = &reqres.CreateUserResponse{}
	if err := json.NewDecoder(respBody).Decode(&payload); err != nil {
		log.Fatal("Error decoding json response")
	}

	userID = payload.User.ID
}

func getToken(respBody io.Reader) {
	var payload = &reqres.LoginResponse{}
	if err := json.NewDecoder(respBody).Decode(&payload); err != nil {
		log.Fatal("Error decoding json response")
	}

	token = payload.Token
}
