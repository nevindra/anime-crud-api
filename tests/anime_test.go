package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/nevindra/sample-go-crud/models"
)

// Create a Test Get All Animes
func TestGetAnimes(t *testing.T) {
	resp, err := http.Get("http://localhost:3000/api/animes")
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	// check if response is empty
	if len(body) == 0 {
		t.Error("No animes found")
	}

	fmt.Println(string(body))
}

// Create a Test Get Anime By ID
func TestGetAnime(t *testing.T) {
	resp, err := http.Get("http://localhost:3000/api/animes/1")
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	// check if response is empty
	if len(body) == 0 {
		t.Error("No anime found")
	}

	fmt.Println(string(body))
}

// Create a Test Create Anime with invalid input
func TestCreateAnimeInvalid(t *testing.T) {
	anime := models.Anime{
		Title:       "Naruto",
		Description: "Naruto is a Japanese manga series written and illustrated by Masashi Kishimoto. It tells the story of Naruto Uzumaki, an adolescent ninja who constantly searches for recognition and dreams to become the Hokage, the ninja in his village who is acknowledged as the leader and the strongest of all.",
	}

	// convert struct to json
	jsonData, err := json.Marshal(anime)
	if err != nil {
		t.Error(err)
	}

	// create request
	req, err := http.NewRequest("POST", "http://localhost:3000/api/animes", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Error(err)
	}

	// set header
	req.Header.Set("Content-Type", "application/json")

	// send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	// if status code is not 400
	if resp.StatusCode != 400 {
		t.Error("Status code is not 400")
	}

	// if status code is 400 then finish the test
	fmt.Println("Status code is 400")
}

// Create a Test Create Anime with valid input
func TestCreateAnimeValid(t *testing.T) {
	anime := models.Anime{
		Title:       "Naruto",
		Description: "Naruto is a Japanese manga series written and illustrated by Masashi Kishimoto. It tells the story of Naruto Uzumaki, an adolescent ninja who constantly searches for recognition and dreams to become the Hokage, the ninja in his village who is acknowledged as the leader and the strongest of all.",
		Episodes:    220,
	}

	// convert struct to json
	jsonData, err := json.Marshal(anime)
	if err != nil {
		t.Error(err)
	}

	// create request
	req, err := http.NewRequest("POST", "http://localhost:3000/api/animes", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Error(err)
	}

	// set header
	req.Header.Set("Content-Type", "application/json")

	// send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	// if status code is not 201
	if resp.StatusCode != 201 {
		t.Error("Status code is not 201")
	}

	// if status code is 201 then finish the test
	fmt.Println("Status code is 201")
}
