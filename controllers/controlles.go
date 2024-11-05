package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/Douirat/groupie-tracker/models"
)

var Tmpl *template.Template

func GitAllGroups(wr http.ResponseWriter, rq *http.Request) {
	// Retrieve artist data
	artists, err := models.GetAllArtists()
	if err != nil {
		http.Error(wr, "Error retrieving data from the server!", http.StatusInternalServerError)
		return
	}

	// Execute the template with the artists data directly
	if err := Tmpl.ExecuteTemplate(wr, "home.html", artists); err != nil {
		http.Error(wr, "Error rendering template", http.StatusInternalServerError)
		fmt.Println("Template execution error:", err)
	}
}

// Get group by id:
func GetGroupById(wr http.ResponseWriter, rq *http.Request) {
}
