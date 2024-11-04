package models

import (
	"encoding/json"
	"io"
	"net/http"
)

// Create the artist structure:
type Artist struct {
	Id              int      `json:"id"`
	Image           string   `json:"image"`
	Name            string   `json:"name"`
	Members         []string `json:"members"`
	CreationDate    uint     `json:"creationDate"`
	LocationsURL    string   `json:"locations"`
	ConcertDatesURL string   `json:"concertDates"`
	Relations       string   `json:"relations"`
}

// Get all the data from the Artist API
func GetAllArtists() ([]Artist, error) {
	var data []byte
	var err error
	var artists []Artist
	var response *http.Response
	url := "https://groupietrackers.herokuapp.com/api/artists"
	response, err = http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	data, err = io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &artists)
	if err != nil {
		return nil, err
	}

	return artists, nil
}
