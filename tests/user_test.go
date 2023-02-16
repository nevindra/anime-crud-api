package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Create mock user with email and password
func mockUser() User {
	return User{
		Email:    "nevindra@gmail.com",
		Password: "123456",
	}
}

// Create a mock user with invalid email and password
func mockUserInvalid() User {
	return User{
		Email:    "nevindra",
		Password: "123",
	}
}

// Test Get All Users
func TestGetUsers(t *testing.T) {
	resp, err := http.Get("http://localhost:3000/api/users")
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	// check if response is empty
	if len(body) == 0 {
		t.Error("No users found")
	}

	fmt.Println(string(body))
}

// Test Get User By ID
func TestGetUser(t *testing.T) {
	resp, err := http.Get("http://localhost:3000/api/users/737RJsKyvKmagvs3tPV3n")
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	// check if user is found
	if resp.StatusCode == 404 {
		t.Error("User not found")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(body))
}

// Test Create User with empty body
func TestCreateUserEmptyBody(t *testing.T) {
	resp, err := http.Post("http://localhost:3000/api/users", "application/json", nil)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	// check if user is created
	if resp.StatusCode == 201 {
		t.Error("User created")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(body))
}

// Create a Negative Test Case for Create User with Invalid Email and Password
func TestCreateUserInvalid(t *testing.T) {
	user := mockUserInvalid()

	// convert user to json
	json, err := json.Marshal(user)
	if err != nil {
		t.Error(err)
	}

	resp, err := http.Post("http://localhost:3000/api/users", "application/json", bytes.NewBuffer(json))

	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	// check if user is created
	if resp.StatusCode == 201 {
		t.Error("User created")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	// check if `message` field is "Invalid email format"
	if string(body) != `{"message":"Invalid email format"}` {
		t.Error("Invalid message")
	}

	fmt.Println(string(body))
}

// Create a Positive Test Case for Create User
func TestCreateUser(t *testing.T) {
	user := mockUser()

	// convert user to json
	json, err := json.Marshal(user)
	if err != nil {
		t.Error(err)
	}

	resp, err := http.Post("http://localhost:3000/api/users", "application/json", bytes.NewBuffer(json))

	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	// check if user is created
	if resp.StatusCode != 201 {
		t.Error("User not created")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(body))
}

// Create a case where email is already taken
func TestCreateUserEmailTaken(t *testing.T) {
	user := mockUser()

	// convert user to json
	json, err := json.Marshal(user)
	if err != nil {
		t.Error(err)
	}

	resp, err := http.Post("http://localhost:3000/api/users", "application/json", bytes.NewBuffer(json))

	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	// check if user is created
	if resp.StatusCode == 201 {
		t.Error("User created")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	// check if `message` field is "Email already taken"
	if string(body) != `{"message":"Email already taken"}` {
		t.Error("Invalid message")
	}

	fmt.Println(string(body))
}
