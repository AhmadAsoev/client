package handleFunc

import (
	"client/pkg/application/services"
	"client/pkg/domain"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func UpdateByID(w http.ResponseWriter, r *http.Request) {
	var request domain.Client
	id := mux.Vars(r)["id"]

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte("Can't Decode into json")); err != nil {
			log.Println("handleFunc/UpdateByID/Decode/Write")
			return
		}
		return
	}

	response := services.UpdateByID(id, request)
	response.Send(w, "UpdateByID")
}
