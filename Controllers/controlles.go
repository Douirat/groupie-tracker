package Controllers

import (
	"fmt"
	"net/http"

	"github.com/Douirat/groupie-tracker/Models"
)

func GitAllGroups(wr http.ResponseWriter, rq *http.Request){
	art := Models.Artists{}
	fmt.Println(art)
}