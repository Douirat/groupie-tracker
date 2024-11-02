package Router

import(
	"github.com/gorilla/mux"
	"github.com/Douirat/groupie-tracker/Controllers"
)

func GroupieRouter(router *mux.Router) {
	router.HandleFunc("/groups", Controllers.GitAllGroups)
}