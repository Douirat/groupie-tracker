package models

import (
	"encoding/json"
	"io"
	"net/http"
)

// Create a relation structure:
type Relation struct {
	DatesLocations map[string][]string `json:"datesLocations"`
}

// Create a date structure:
type Date struct {
Id int
Dates []string `json:"dates"`
}

// Create a location structure:
type Location struct {
	Id int
	Locations []string `json:"locations"`
}

// Create the artist structure:
type Artist struct {
	Id              int      `json:"id"`
	Image           string   `json:"image"`
	Name            string   `json:"name"`
	Members         []string `json:"members"`
	CreationDate    uint     `json:"creationDate"`
	Locations []string
	Dates []string
	DatesLocations  map[string][]string
	LocationsURL    string   `json:"locations"`
	ConcertDatesURL string   `json:"concertDates"`
	Relations       string   `json:"relations"`
}

// Declare a global variable to hold necessary data about the artists:
var artists []Artist

// Get all the data from the Artist API
func GetAllArtists() ([]Artist, error) {
	var data []byte
	var err error


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

// Get an artist based on id:
func GetArtisDetails() {

}