package handleFunc

import (
	"client/pkg/application/services"
	"net/http"
)

func GetAllClients(w http.ResponseWriter, r *http.Request) {
	response := services.GetAllClients()
	response.Send(w, "GetAllClients")
}
