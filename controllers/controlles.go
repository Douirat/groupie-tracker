package controllers

import (
	"html/template"
	"net/http"

	"github.com/Douirat/groupie-tracker/models"
)

var Tmpl *template.Template

func GitAllGroups(wr http.ResponseWriter, rq *http.Request) {
	artists, err := models.GetAllArtists()
	if err != nil {
		http.Error(wr, "Error difficulty retrieving dqtq from the server!", http.StatusInternalServerError)
	}
	wr.WriteHeader(http.StatusOK)
	err = Tmpl.ExecuteTemplate(wr, "index.html", artists)
	if err != nil {
		http.Error(wr, "Error difficulty retrieving dqtq from the server!", http.StatusInternalServerError)
	}
}
