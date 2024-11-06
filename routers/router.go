package routers

import(
	"github.com/gorilla/mux"
	"github.com/Douirat/groupie-tracker/controllers"
)

func GroupieRouter(router *mux.Router) {
	router.HandleFunc("/", controllers.GitAllGroups).Methods("GET")
	router.HandleFunc("/group/{id}", controllers.GetGroupById).Methods("GET")
	router.HandleFunc("/error", controllers.ErrorHandler).Methods("GET")
}