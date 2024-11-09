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
	Locations       *Location
	Dates           *Date
	DatesLocations  *Relation
	LocationsURL    string `json:"locations"`
	ConcertDatesURL string `json:"concertDates"`
	RelationsURL    string `json:"relations"`
}

// Create a global variable to the artists:
var (
	artists []Artist
	err     error
)

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
	url := "https://groupietrackers.herokuapp.com/api/artists"
	err = FetchData(url, &artists)
	if err != nil {
		return nil, err
	}
	return artists, nil
}

// Get UserDates:
func GetDates(artist *Artist) {
	fmt.Println(artist.ConcertDatesURL)
	artist.Dates = new(Date)
	err := FetchData(artist.ConcertDatesURL, artist.Dates)
	if err != nil {
		return
	}
	fmt.Println(artist.Dates.Dates)
}


// Get an artist based on id:
func GetArtisDetails(id int) *Artist {
	artists, err = GetAllArtists()
	artists[id].Dates = new(Date)
	err = FetchData(artists[id].ConcertDatesURL, artists[id].Dates)
	artists[id].Locations = new(Location)
	err = FetchData(artists[id].LocationsURL, artists[id].Locations)
	artists[id].DatesLocations = new(Relation)
	err = FetchData(artists[id].RelationsURL, artists[id].DatesLocations)
	if err != nil {
		return nil
	}
	return &artists[id]
}
