package handleFunc

import (
	"client/pkg/application/services"
	"github.com/gorilla/mux"
	"net/http"
)

func GetById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	response := services.GetById(id)
	response.Send(w, "GetById")
}
