package handleFunc

import (
	"client/pkg/application/services"
	"github.com/gorilla/mux"
	"net/http"
)

func GetByName(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	response := services.GetByName(name)
	response.Send(w, "GetByName")
}
