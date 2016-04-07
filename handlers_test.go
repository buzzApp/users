package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateUserHTTPEndpoint(t *testing.T) {
	server := httptest.NewServer(handleCreateUser(userService{}))
	defer server.Close()

	usersURL := fmt.Sprintf("%s/users", server.URL)

	userJSON := `{"email": "test@test.com", "first_name": "testname", "last_name": "lasttest", "password": "testPass", "role": "student", "username": "testuser"}`

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
}
