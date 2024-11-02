package Models

import(
	"decode/json"
	// "net/http"
)


// Create a date structure:
type Date struct {
	Id int
	Dates []string
}

// Create a location struct:
type Location struct {
	Id int `json:"id"`
	Locations []string `json:"locations"`
	Dates *Date `json:"dates"`
}

// Create the artist structure:
type Artists struct {
	Id  int `json:"id"`
	Image string `json:"image"`
	Name string `json:"name"`
	Members []string `json:"members"`
	CreationDate uint `json:"creationDate"`
	Locations  []*Location `json:"locations"`
}