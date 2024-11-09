package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/Douirat/groupie-tracker/models"
)

var Tmpl *template.Template

type Errony struct {
	Status    string
	ErrorType string
}

func GitAllGroups(wr http.ResponseWriter, rq *http.Request) {
	// Retrieve artist data
	artists, err := models.GetAllArtists()
	if err != nil {
		http.Error(wr, "Error retrieving data from the server!", http.StatusInternalServerError)
		return
	}
	// Execute the template with the artists data directly
	if err := Tmpl.ExecuteTemplate(wr, "home.html", artists); err != nil {
		fmt.Println("Template execution error:", err)
		http.Redirect(wr, rq, "/error?status=404", http.StatusMovedPermanently)
	}
}

// Get group by id:
func GetGroupById(wr http.ResponseWriter, rq *http.Request) {
	var err error
	params := mux.Vars(rq)
	artistId, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(wr, "Error retrieving data from the server!", http.StatusInternalServerError)
		return
	}
	artist := models.GetArtisDetails(artistId-1)
	if artist == nil {
		http.Error(wr, "Error retrieving data from the server!", http.StatusNotFound)
		return
	}
	if err = Tmpl.ExecuteTemplate(wr, "details.html", artist); err != nil {
		
		fmt.Println("Template execution error:", err)
	}
}

// Controller to handle errors:
func ErrorHandler(wr http.ResponseWriter, rq *http.Request) {
	s := rq.URL.Query().Get("status")
	err := Errony{Status: s,ErrorType: "test"}
	if e := Tmpl.ExecuteTemplate(wr, "error.html", err); e != nil {
		http.Redirect(wr, rq, "/error?status=404", http.StatusPermanentRedirect)
	}
}
