package main

import (
	"html/template"
	"log"

	"github.com/Douirat/groupie-tracker/controllers"
	"github.com/Douirat/groupie-tracker/routers"
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
)

func main() {
	var err error
	controllers.Tmpl, err = template.ParseGlob("views/*.html")
	if err != nil {
		log.Fatal(err)
	}
	var rtr = mux.NewRouter()
	routers.GroupieRouter(rtr)
	fmt.Println("Router listening at port :9090")
	http.ListenAndServe(":9090", rtr)
}