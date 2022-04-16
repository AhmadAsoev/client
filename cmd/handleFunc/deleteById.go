package handleFunc

import (
	"client/pkg/application/services"
	"github.com/gorilla/mux"
	"net/http"
)

func DeleteById(w http.ResponseWriter, r *http.Request) {
	urlId := mux.Vars(r)["id"]
	response := services.DeleteById(urlId)
	response.Send(w, "DeleteById")
}
