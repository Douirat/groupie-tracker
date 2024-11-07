package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Create a relation structure:
type Relation struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// Create a date structure:
type Date struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

// Create a location structure:
type Location struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
}

// Create the artist structure:
type Artist struct {
	Id              int      `json:"id"`
	Image           string   `json:"image"`
	Name            string   `json:"name"`
	Members         []string `json:"members"`
	CreationDate    uint     `json:"creationDate"`
	Locations       []string
	Dates           *Date
	DatesLocations  map[string][]string
	LocationsURL    string `json:"locations"`
	ConcertDatesURL string `json:"concertDates"`
	Relations       string `json:"relations"`
}


// A general dynamic function to retrieve data from the API:
func FetchData(url string, inst any) error {
	Response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer Response.Body.Close()
	if Response.StatusCode != http.StatusOK {
		return errors.New(Response.Status)
	}
	err = json.NewDecoder(Response.Body).Decode(&inst)
	if err != nil {
		return err
	}
	return nil
}

// Get all the data from the Artist API
func GetAllArtists() ([]Artist, error) {
	var err error
	var artists []Artist
	url := "https://groupietrackers.herokuapp.com/api/artists"
	err = FetchData(url, &artists)	
	if err != nil {
		return nil, err
	}
	return artists, nil
}

// Get UserDates:
func GetDates(url string) {
	fmt.Println(url)
	date := new(Date)
	err := FetchData(url, date)
	if err != nil {
		return
	}
	fmt.Println(date.Dates)
}


// Get an artist based on id:
func GetArtisDetails(id int) *Artist {
	artists, err := GetAllArtists()
	if err != nil {
		return nil
	}
	GetDates(artists[id].ConcertDatesURL)
	return &artists[id]
}

// 
