package routers

import(
	"github.com/gorilla/mux"
	"github.com/Douirat/groupie-tracker/controllers"
)

func GroupieRouter(router *mux.Router) {
	router.HandleFunc("/groups", controllers.GitAllGroups).Methods("GET")
}