package handleFunc

import (
	"client/pkg/application/services"
	"client/pkg/domain"
	"encoding/json"
	"log"
	"net/http"
)

func AddClient(w http.ResponseWriter, r *http.Request) {
	var request domain.Client

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte("Can't Decode into json")); err != nil {
			log.Println("handleFunc/AddClient/Decode/Write")
		}
	}
	response := services.AddClient(request)
	response.Send(w, "AddClient")
}
